package book

import (
	"github.com/maclav3/cleanarch-hegaxon-demo/internal/log"
)

// Note: granted, that's a lot of boilerplate.
// However:
// 1) These command handlers are quite simple, but as they grow, the separation of concerns will be appreciated
// 2) In practice, they might be go-generated with gowrap for example, which reduces the workload on the developer.
type addBookCommandHandlerLogger struct {
	logger  log.Logger
	wrapped AddBookCommandHandler
}

func (l addBookCommandHandlerLogger) Handle(cmd Add) (err error) {
	logger := l.logger.
		WithField(log.CommandHandlerKey, "book.Add").
		WithJSON(log.CommandKey, cmd)
	defer func() {
		if err != nil {
			logger.WithError(err).Error("could not add book")
			return
		}

		logger.Info("add book successful")
	}()

	return l.wrapped.Handle(cmd)
}

type loanBookCommandHandlerLogger struct {
	logger  log.Logger
	wrapped LoanBookCommandHandler
}

func (l loanBookCommandHandlerLogger) Handle(cmd Loan) (err error) {
	logger := l.logger.
		WithField(log.CommandHandlerKey, "book.Loan").
		WithJSON(log.CommandKey, cmd)
	defer func() {
		if err != nil {
			logger.WithError(err).Error("Could not loan book")
			return
		}

		logger.Info("loan book successful")
	}()

	return l.wrapped.Handle(cmd)
}
