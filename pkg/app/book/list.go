package book

import (
	"github.com/maclav3/cleanarch-hegaxon-demo/pkg/domain/book"
	"github.com/pkg/errors"
)

// listBooksRepository defines the methods that we expect of a repository
// that is able to support a ListBooks query.
type listBooksRepository interface {
	AllBooks(q ListBooksQuery) ([]*book.Book, error)
}

type ListBooksQuery struct {
	// Loaned, if not-nil, this will filter books by their loaned status.
	Loaned *bool
}

func (i *inventory) ListBooks(q ListBooksQuery) ([]*book.Book, error) {
	books, err := i.bookRepo.AllBooks(q)
	if err != nil {
		return nil, errors.Wrap(err, "could not list books")
	}

	return books, nil
}
