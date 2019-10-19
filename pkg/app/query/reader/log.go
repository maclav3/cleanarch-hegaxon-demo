package reader

import (
	"github.com/maclav3/cleanarch-hegaxon-demo/internal/log"
	"github.com/maclav3/cleanarch-hegaxon-demo/pkg/domain/reader"
)

type listReadersQueryHandlerLogger struct {
	logger  log.Logger
	wrapped ListReadersQueryHandler
}

func (l *listReadersQueryHandlerLogger) Query(q ListQuery) (readers []*reader.Reader, err error) {
	logger := l.logger.
		WithField(log.QueryHandlerKey, "reader.ListQuery").
		WithJSON(log.QueryKey, q)
	defer func() {
		if err != nil {
			logger.WithError(err).Error("could not list readers")
			return
		}

		// don't log successful queries, that would be too much clutter
	}()

	return l.wrapped.Query(q)
}
