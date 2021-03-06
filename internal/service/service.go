package service

import (
	"context"
	"sync"
	"time"

	"github.com/maclav3/cleanarch-hegaxon-demo/internal/log"
	adaptersBook "github.com/maclav3/cleanarch-hegaxon-demo/pkg/adapters/book"
	adaptersReader "github.com/maclav3/cleanarch-hegaxon-demo/pkg/adapters/reader"
	"github.com/maclav3/cleanarch-hegaxon-demo/pkg/app"
	bookCommand "github.com/maclav3/cleanarch-hegaxon-demo/pkg/app/command/book"
	readerCommand "github.com/maclav3/cleanarch-hegaxon-demo/pkg/app/command/reader"
	bookQuery "github.com/maclav3/cleanarch-hegaxon-demo/pkg/app/query/book"
	readerQuery "github.com/maclav3/cleanarch-hegaxon-demo/pkg/app/query/reader"
	"github.com/maclav3/cleanarch-hegaxon-demo/pkg/port/cli"

	"github.com/pkg/errors"
)

type Service struct {
	Logger log.Logger

	App *app.Application

	onStartup  []startupCallback
	onShutdown []shutdownCallback

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
		onStartup:  []startupCallback{},
		onShutdown: []shutdownCallback{},

		callbackTimeout: 10 * time.Second,
		closeCh:         make(chan struct{}),

		runMutex:      sync.Mutex{},
		running:       false,
		shutdownMutex: sync.Mutex{},
		shuttingDown:  false,
	}

	service.Logger = log.NewLogger("cleanarch-hexagon-demo")

	// Initialize the repositories.
	// The simple in-memory repositories are enough for this simple demonstration app.
	// In production, relational or NoSQL databases might be used to satisfy the dependencies of the app layer.
	bookRepository := adaptersBook.NewMemoryRepository()
	readerRepository := adaptersReader.NewMemoryRepository()

	// Here, we construct and inject all dependencies manually. However, in a larger project, this becomes increasingly hard.
	// In this case, some kind of automated DI approach is preferred, for example github.com/google/wire.
	// Also, it allows us to initialize each application handler per request, which aids e.g. with tracing and transactions.
	service.App = &app.Application{
		Queries: &app.Queries{
			ListBooks:   bookQuery.NewListBooksQueryHandler(service.Logger, bookRepository),
			ListReaders: readerQuery.NewListReadersQueryHandler(service.Logger, readerRepository),
		},
		Commands: &app.Commands{
			AddBook:          bookCommand.NewAddBookCommandHandler(service.Logger, bookRepository),
			LoanBook:         bookCommand.NewLoanBookCommandHandler(service.Logger, bookRepository, readerRepository),
			AddReader:        readerCommand.NewAddReaderCommandHandler(service.Logger, readerRepository),
			ActivateReader:   readerCommand.NewActivateReaderCommandHandler(service.Logger, readerRepository),
			DeactivateReader: readerCommand.NewDeactivateReaderCommandHandler(service.Logger, readerRepository),
		},
	}

	// Initialize the ports.
	// There is a single simple CLI router that is running while the application lives
	// It captures text from a socket with zeroMQ, and runs cobra commands with the text.
	cliRouter := cli.NewRouter(service.Logger, service.App, "localhost:5555")
	service.onStartupShutdown(
		cliRouter.Run,
		cliRouter.Close,
	)

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
func (s *Service) Run(ctx context.Context) error {
	s.runMutex.Lock()
	defer s.runMutex.Unlock()
	if s.running {
		return nil
	}
	s.running = true

	errChan := make(chan error)
	for _, callback := range s.onStartup {
		go func() {
			errChan <- callback(ctx)
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
	defer s.shutdownMutex.Unlock()
	if s.shuttingDown {
		return nil
	}
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

func (s *Service) onStartupShutdown(startup startupCallback, shutdown shutdownCallback) {
	s.onStartup = append(s.onStartup, startup)
	s.onShutdown = append(s.onShutdown, shutdown)
}

type startupCallback func(ctx context.Context) error
type shutdownCallback func() error
