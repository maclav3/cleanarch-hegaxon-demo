package reader

import "github.com/maclav3/cleanarch-hegaxon-demo/internal/log"

// Note: granted, that's a lot of boilerplate.
// However:
// 1) These command handlers are quite simple, but as they grow, the separation of concerns will be appreciated
// 2) In practice, they might be go-generated with gowrap for example, which reduces the workload on the developer.
type addReaderCommandHandlerLogger struct {
	logger  log.Logger
	wrapped AddReaderCommandHandler
}

func (l *addReaderCommandHandlerLogger) Handle(cmd Add) (err error) {
	logger := l.logger.
		WithField(log.CommandHandlerKey, "reader.Add").
		WithJSON(log.CommandKey, cmd)
	defer func() {
		if err != nil {
			logger.WithError(err).Error("Could not add reader")
			return
		}

		logger.Info("Adding reader successful")
	}()

	return l.wrapped.Handle(cmd)
}

type deactivateReaderCommandHandlerLogger struct {
	logger  log.Logger
	wrapped DeactivateReaderCommandHandler
}

func (l *deactivateReaderCommandHandlerLogger) Handle(cmd Deactivate) (err error) {
	logger := l.logger.
		WithField(log.CommandHandlerKey, "reader.Deactivate").
		WithJSON(log.CommandKey, cmd)
	defer func() {
		if err != nil {
			logger.WithError(err).Error("Could not deactivate reader")
			return
		}

		logger.Info("Deactivating reader successful")
	}()

	return l.wrapped.Handle(cmd)
}

type activateReaderCommandHandlerLogger struct {
	logger  log.Logger
	wrapped ActivateReaderCommandHandler
}

func (l *activateReaderCommandHandlerLogger) Handle(cmd Activate) (err error) {
	logger := l.logger.
		WithField(log.CommandHandlerKey, "reader.Activate").
		WithJSON(log.CommandKey, cmd)
	defer func() {
		if err != nil {
			logger.WithError(err).Error("could not activate reader")
			return
		}

		logger.Info("activate reader successful")
	}()

	return l.wrapped.Handle(cmd)
}
