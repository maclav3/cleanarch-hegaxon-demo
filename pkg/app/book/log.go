package book

import (
	"github.com/maclav3/cleanarch-hegaxon-demo/internal/log"
	"github.com/maclav3/cleanarch-hegaxon-demo/pkg/domain/book"
)

const inventoryHandlerName = "book.Inventory"

type inventoryLoggingDecorator struct {
	logger  log.Logger
	wrapped Inventory
}

func (d inventoryLoggingDecorator) ListBooks(q ListBooksQuery) (books []*book.Book, err error) {
	logger := d.logger.
		WithField(log.QueryHandlerKey, inventoryHandlerName).
		WithField(log.QueryKey, q)
	defer func() {
		if err != nil {
			logger.WithError(err).Error("could not list books")
			return
		}

		logger.Info("list books successful")
	}()

	return d.wrapped.ListBooks(q)
}

func (d inventoryLoggingDecorator) AddBook(cmd AddBook) (err error) {
	logger := d.logger.
		WithField(log.CommandHandlerKey, inventoryHandlerName).
		WithField(log.CommandKey, cmd)
	defer func() {
		if err != nil {
			logger.WithError(err).Error("could not add book")
			return
		}

		logger.Info("add book successful")
	}()

	return d.wrapped.AddBook(cmd)
}

func (d inventoryLoggingDecorator) LoanBook(cmd LoanBook) (err error) {
	logger := d.logger.
		WithField(log.CommandHandlerKey, inventoryHandlerName).
		WithField(log.CommandKey, cmd)
	defer func() {
		if err != nil {
			logger.WithError(err).Error("could not add book")
			return
		}

		logger.Info("add book successful")
	}()

	return d.wrapped.LoanBook(cmd)
}
