package reader

import (
	"github.com/maclav3/cleanarch-hegaxon-demo/internal/log"
	"github.com/maclav3/cleanarch-hegaxon-demo/pkg/domain/reader"
)

type ListReadersQueryHandler interface {
	Query(q ListQuery) ([]*reader.Reader, error)
}

type listReadersQueryHandler struct {
	repo listReadersRepository
}

type listReadersRepository interface {
	ListReaders(q ListQuery) ([]*reader.Reader, error)
}

type ListQuery struct {
}

func NewListReadersQueryHandler(logger log.Logger, repo listReadersRepository) ListReadersQueryHandler {
	// we panic if any dependency is nil
	if logger == nil {
		panic("logger is nil")
	}
	// because it is not a recoverable state
	// and should be fixed in compile time
	if repo == nil {
		panic("readerRepo is nil")
	}

	return &listReadersQueryHandlerLogger{
		logger: logger,
		wrapped: &listReadersQueryHandler{
			repo: repo,
		},
	}
}

func (h *listReadersQueryHandler) Query(q ListQuery) ([]*reader.Reader, error) {
	return nil, nil
}
