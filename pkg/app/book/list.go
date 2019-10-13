package book

import "github.com/maclav3/cleanarch-hegaxon-demo/pkg/domain/book"

type ListBooksQuery struct {
	// Loaned, if not-nil, this will filter books by their loaned status.
	Loaned *bool
}

func (b *inventory) ListBooks(q ListBooksQuery) ([]*book.Book, error) {
	panic("implement me")
}
