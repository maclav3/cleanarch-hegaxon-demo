package book

import (
	"github.com/maclav3/cleanarch-hegaxon-demo/pkg/domain/book"
	"github.com/pkg/errors"
)

// listBooksRepository defines the methods that we expect of a repository
// that is able to support a List query.
type listBooksRepository interface {
	AllBooks(q ListQuery) ([]*book.Book, error)
}

type ListQuery struct {
	// Loaned, if not-nil, this will filter books by their loaned status.
	Loaned *bool
}

func (i *inventory) List(q ListQuery) ([]*book.Book, error) {
	books, err := i.bookRepo.AllBooks(q)
	if err != nil {
		return nil, errors.Wrap(err, "could not list books")
	}

	return books, nil
}
