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

func (d inventoryLoggingDecorator) List(q ListQuery) (books []*book.Book, err error) {
	logger := d.logger.
		WithField(log.QueryHandlerKey, inventoryHandlerName).
		WithField(log.MethodKey, "List").
		WithJSON(log.QueryKey, q)
	defer func() {
		if err != nil {
			logger.WithError(err).Error("could not list books")
			return
		}

		logger.Info("list books successful")
	}()

	return d.wrapped.List(q)
}

func (d inventoryLoggingDecorator) Add(cmd Add) (err error) {
	logger := d.logger.
		WithField(log.CommandHandlerKey, inventoryHandlerName).
		WithField(log.MethodKey, "Add").
		WithJSON(log.CommandKey, cmd)
	defer func() {
		if err != nil {
			logger.WithError(err).Error("could not add book")
			return
		}

		logger.Info("add book successful")
	}()

	return d.wrapped.Add(cmd)
}

func (d inventoryLoggingDecorator) Loan(cmd Loan) (err error) {
	logger := d.logger.
		WithField(log.CommandHandlerKey, inventoryHandlerName).
		WithField(log.MethodKey, "Loan").
		WithJSON(log.CommandKey, cmd)
	defer func() {
		if err != nil {
			logger.WithError(err).Error("could not add book")
			return
		}

		logger.Info("add book successful")
	}()

	return d.wrapped.Loan(cmd)
}
