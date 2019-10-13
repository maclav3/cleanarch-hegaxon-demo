package service

import (
	"context"
	"sync"
	"time"

	"github.com/pkg/errors"

	"github.com/maclav3/cleanarch-hegaxon-demo/internal/log"
	adaptersBook "github.com/maclav3/cleanarch-hegaxon-demo/pkg/adapters/book"
	"github.com/maclav3/cleanarch-hegaxon-demo/pkg/adapters/reader"
	"github.com/maclav3/cleanarch-hegaxon-demo/pkg/app/book"
)

type Service struct {
	Logger log.Logger

	BookInventory book.Inventory

	onStartup  []callback
	onShutdown []callback

	// callbackTimeout limits the time spent on startup/shutdown callbacks. Defaults to 30s per each callback.
	callbackTimeout time.Duration
	closeCh         chan struct{}

	runMutex      sync.Mutex
	running       bool
	shutdownMutex sync.Mutex
	shuttingDown  bool
}

// NewService returns a new Service, that has the application command & query handlers initialized.
// Run() needs to be called to start any servers providing in/out to the service.
// Shutdown() then needs to be called to execute a graceful shutdown, closing the components in an ordered way.
func NewService(ctx context.Context) *Service {
	service := &Service{
		onStartup:  []callback{},
		onShutdown: []callback{},

		closeCh: make(chan struct{}),

		runMutex:      sync.Mutex{},
		running:       false,
		shutdownMutex: sync.Mutex{},
		shuttingDown:  false,
	}

	service.Logger = log.NewLogger("cleanarch-hexagon-demo")

	// initialize the repositories.
	// the simple in-memory repositories are enough for this simple demonstration app.
	// in production, relational or NoSQL databases might be used to satisfy the dependencies of the app layer.
	bookRepository := adaptersBook.NewMemoryRepository()
	readerRepository := reader.NewMemoryRepository()

	service.BookInventory = book.NewBookInventory(service.Logger, bookRepository, readerRepository)

	go func() {
		// shutdown on ctx close
		<-ctx.Done()
		service.Logger.Info("Service context expired, shutting down service")
		err := service.Shutdown()
		if err != nil {
			panic(err)
		}
	}()

	return service
}

// Run executes all the registered startup functions (spinning up servers etc.)
// Run exits with error if startup was unsuccessful, or returns after service is up and running.
func (s *Service) Run() error {
	s.runMutex.Lock()
	if s.running {
		return nil
	}
	s.runMutex.Unlock()
	s.running = true

	errChan := make(chan error)
	for _, callback := range s.onStartup {
		go func() {
			errChan <- callback()
		}()

		select {
		case err := <-errChan:
			if err != nil {
				return errors.Wrap(err, "error executing onStartup callback")
			}
		case <-time.After(s.callbackTimeout):
			return errors.New("timeouted while exectuting onStartup callback")
		}
	}

	return nil
}

func (s *Service) Running() <-chan struct{} {
	return s.closeCh
}

func (s *Service) Shutdown() error {
	s.shutdownMutex.Lock()
	if s.shuttingDown {
		return nil
	}
	s.shutdownMutex.Unlock()
	s.shuttingDown = true

	errChan := make(chan error)
	for _, callback := range s.onShutdown {
		go func() {
			errChan <- callback()
		}()

		select {
		case err := <-errChan:
			if err != nil {
				return errors.Wrap(err, "error executing onShutdown callback")
			}
		case <-time.After(s.callbackTimeout):
			return errors.New("timeouted while exectuting onShutdown callback")
		}
	}

	close(s.closeCh)
	return nil
}

type callback func() error
