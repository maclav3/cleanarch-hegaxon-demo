package book

import (
	"github.com/maclav3/cleanarch-hegaxon-demo/internal/log"
	"github.com/maclav3/cleanarch-hegaxon-demo/pkg/domain/book"
)

type listBooksQueryHandlerLogger struct {
	logger  log.Logger
	wrapped ListBooksQueryHandler
}

func (l listBooksQueryHandlerLogger) Query(q ListQuery) (books []*book.Book, err error) {
	logger := l.logger.
		WithField(log.QueryHandlerKey, "book.List").
		WithJSON(log.QueryKey, q)
	defer func() {
		if err != nil {
			logger.WithError(err).Error("could not list books")
			return
		}

		// don't log successful queries, that would be too much clutter
	}()

	return l.wrapped.Query(q)
}
