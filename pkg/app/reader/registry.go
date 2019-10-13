package reader

import (
	"github.com/maclav3/cleanarch-hegaxon-demo/internal/log"
	"github.com/maclav3/cleanarch-hegaxon-demo/pkg/domain/reader"
)

type readerRepo interface {
	// reader.Repository interface provides the basic reader repository methods.
	reader.Repository
	// All returns all readers from the repository.
	All() ([]*reader.Reader, error)
}

type Registry interface {
	Add(cmd Add) error
	Activate(cmd Activate) error
	Deactivate(cmd Deactivate) error
	List(q ListQuery) ([]*reader.Reader, error)
}

type registry struct {
	readerRepo readerRepo
}

func NewRegistry(logger log.Logger, readerRepo readerRepo) Registry {
	// we panic if any dependency is nil
	if logger == nil {
		panic("logger is nil")
	}
	// because it is not a recoverable state
	// and should be fixed in compile time
	if readerRepo == nil {
		panic("readerRepo is nil")
	}

	return &registryLoggingDecorator{
		logger: logger,
		wrapped: &registry{
			readerRepo: readerRepo,
		},
	}
}
