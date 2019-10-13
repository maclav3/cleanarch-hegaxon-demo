package book

import (
	"github.com/maclav3/cleanarch-hegaxon-demo/internal/log"
	"github.com/maclav3/cleanarch-hegaxon-demo/pkg/domain/book"
	"github.com/maclav3/cleanarch-hegaxon-demo/pkg/domain/reader"
)

type Inventory interface {
	ListBooks(q ListBooksQuery) ([]*book.Book, error)
	AddBook(cmd AddBook) error
	LoanBook(cmd LoanBook) error
}

type inventory struct {
	bookRepo   book.Repository
	readerRepo reader.Repository
}

func NewBookInventory(
	logger log.Logger,
	bookRepo book.Repository,
	readerRepo reader.Repository,
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
