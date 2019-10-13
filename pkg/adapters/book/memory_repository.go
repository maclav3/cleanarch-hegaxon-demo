package book

import "github.com/maclav3/cleanarch-hegaxon-demo/pkg/domain/book"

type memoryRepository struct {
	books map[book.ID]*book.Book
}

func NewMemoryRepository() book.Repository {
	return &memoryRepository{
		books: map[book.ID]*book.Book{},
	}
}

func (r *memoryRepository) ByID(id book.ID) (*book.Book, error) {
	b, ok := r.books[id]
	if !ok {
		return nil, book.ErrNotFound
	}
	return b, nil
}

func (r *memoryRepository) Save(b *book.Book) error {
	r.books[b.ID()] = b
	return nil
}
