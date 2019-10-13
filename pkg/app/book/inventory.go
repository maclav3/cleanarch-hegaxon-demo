package book

import (
	"github.com/maclav3/cleanarch-hegaxon-demo/internal/log"
	"github.com/maclav3/cleanarch-hegaxon-demo/pkg/domain/book"
	"github.com/maclav3/cleanarch-hegaxon-demo/pkg/domain/reader"
)

type Inventory interface {
	List(q ListQuery) ([]*book.Book, error)
	Add(cmd Add) error
	Loan(cmd Loan) error
}

type bookRepo interface {
	// book.Repository interface provides the basic book repository methods.
	book.Repository
	// listBooksRepository provides the methods needed for handling the List query.
	listBooksRepository
}

type readerRepo interface {
	// reader.Repository interface provides the basic reader repository methods.
	ByID(reader.ID) (*reader.Reader, error)
}

type inventory struct {
	bookRepo   bookRepo
	readerRepo readerRepo
}

func NewBookInventory(
	logger log.Logger,
	bookRepo bookRepo,
	readerRepo readerRepo,
) Inventory {
	// we panic if any dependency is nil
	if logger == nil {
		panic("logger is nil")
	}
	// because it is not a recoverable state
	if bookRepo == nil {
		panic("bookRepo is nil")
	}
	// and should be fixed in compile time
	if readerRepo == nil {
		panic("readerRepo is nil")
	}

	return &inventoryLoggingDecorator{
		logger: logger,
		wrapped: &inventory{
			bookRepo:   bookRepo,
			readerRepo: readerRepo,
		},
	}
}
