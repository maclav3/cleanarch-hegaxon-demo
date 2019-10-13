package reader

import "github.com/pkg/errors"

var (
	// ErrNotFound occurs when a reader could not be found.
	ErrNotFound = errors.New("reader not found")
)

// Repository defines the minimal interface that all implementations of a Reader repository must implement.
type Repository interface {
	// ByID retrieves a reader by ID or returns ErrNotFound.
	ByID(ID) (*Reader, error)
	// Save stores the current state of the Reader in the repository.
	Save(*Reader) error
}
