package book

import (
	"github.com/pkg/errors"
)

var (
	// ErrNotFound occurs when a book could not be found.
	ErrNotFound = errors.New("book not found")
)

// Repository defines the minimal interface that all implementations of a Book repository must implement.
type Repository interface {
	// ByID retrieves a book by ID or returns ErrNotFound.
	ByID(ID) (*Book, error)
	// Save stores the current state of the Book in the repository.
	Save(*Book) error
}
