package book

import (
	appBook "github.com/maclav3/cleanarch-hegaxon-demo/pkg/app/book"
	"github.com/maclav3/cleanarch-hegaxon-demo/pkg/domain/book"
)

type MemoryRepository struct {
	books map[book.ID]*book.Book
}

// NewMemoryRepository returns a new memory repository for books.
func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		books: map[book.ID]*book.Book{},
	}
}

func (r *MemoryRepository) ByID(id book.ID) (*book.Book, error) {
	b, ok := r.books[id]
	if !ok {
		return nil, book.ErrNotFound
	}
	return b, nil
}

func (r *MemoryRepository) Save(b *book.Book) error {
	r.books[b.ID()] = b
	return nil
}

// AllBooks finds all books given the criteria in q.
// Note that it imports both the app and domain layer, which is OK by the rules of Clean Architecture.
// The domain package defines the aggregate, and the app layer the query, which is connected merely to the use case.
func (r *MemoryRepository) AllBooks(q appBook.ListQuery) ([]*book.Book, error) {
	books := []*book.Book{}

	for _, b := range r.books {
		if q.Loaned == nil || *q.Loaned && b.Loaned() || !*q.Loaned && !b.Loaned() {
			books = append(books, b)
		}
	}

	return books, nil
}
