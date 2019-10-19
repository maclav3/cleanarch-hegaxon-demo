package book

import (
	"github.com/maclav3/cleanarch-hegaxon-demo/internal/log"
	"github.com/maclav3/cleanarch-hegaxon-demo/pkg/domain/book"
	"github.com/pkg/errors"
)

type ListBooksQueryHandler interface {
	Query(cmd ListQuery) ([]*book.Book, error)
}

type listBooksQueryHandler struct {
	repo listBooksRepository
}

func NewListBooksQueryHandler(logger log.Logger, repo listBooksRepository) ListBooksQueryHandler {
	// we panic if any dependency is nil
	if logger == nil {
		panic("logger is nil")
	}
	// because it is not a recoverable state
	// and should be fixed in compile time
	if repo == nil {
		panic("readerRepo is nil")
	}

	return &listBooksQueryHandlerLogger{
		logger: logger,
		wrapped: &listBooksQueryHandler{
			repo: repo,
		},
	}
}

// listBooksRepository defines the methods that we expect of a repository
// that is able to support a List query.
type listBooksRepository interface {
	AllBooks(q ListQuery) ([]*book.Book, error)
}

type ListQuery struct {
	// Loaned, if not-nil, this will filter books by their loaned status.
	Loaned *bool
}

func (h *listBooksQueryHandler) Query(q ListQuery) ([]*book.Book, error) {
	books, err := h.repo.AllBooks(q)
	if err != nil {
		return nil, errors.Wrap(err, "could not list books")
	}

	return books, nil
}
