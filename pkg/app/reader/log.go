package reader

import (
	"github.com/maclav3/cleanarch-hegaxon-demo/internal/log"
	"github.com/maclav3/cleanarch-hegaxon-demo/pkg/domain/reader"
)

const registryHandlerName = "reader.Registry"

type registryLoggingDecorator struct {
	logger  log.Logger
	wrapped Registry
}

func (d *registryLoggingDecorator) Add(cmd Add) (err error) {
	logger := d.logger.
		WithField(log.CommandHandlerKey, registryHandlerName).
		WithField(log.MethodKey, "Add").
		WithJSON(log.CommandKey, cmd)
	defer func() {
		if err != nil {
			logger.WithError(err).Error("could not add reader")
			return
		}

		logger.Info("add reader successful")
	}()

	return d.wrapped.Add(cmd)
}

func (d *registryLoggingDecorator) Activate(cmd Activate) (err error) {
	logger := d.logger.
		WithField(log.CommandHandlerKey, registryHandlerName).
		WithField(log.MethodKey, "Activate").
		WithJSON(log.CommandKey, cmd)
	defer func() {
		if err != nil {
			logger.WithError(err).Error("could not activate reader")
			return
		}

		logger.Info("activate reader successful")
	}()

	return d.wrapped.Activate(cmd)
}

func (d *registryLoggingDecorator) Deactivate(cmd Deactivate) (err error) {
	logger := d.logger.
		WithField(log.CommandHandlerKey, registryHandlerName).
		WithField(log.MethodKey, "Deactivate").
		WithJSON(log.CommandKey, cmd)
	defer func() {
		if err != nil {
			logger.WithError(err).Error("could not deactivate reader")
			return
		}

		logger.Info("deactivate reader successful")
	}()

	return d.wrapped.Deactivate(cmd)
}

func (d *registryLoggingDecorator) List(q ListQuery) (readers []*reader.Reader, err error) {
	logger := d.logger.
		WithField(log.QueryHandlerKey, registryHandlerName).
		WithField(log.MethodKey, "List").
		WithJSON(log.QueryKey, q)
	defer func() {
		if err != nil {
			logger.WithError(err).Error("could not list readers")
			return
		}

		logger.Info("list readers successful")
	}()

	return d.wrapped.List(q)
}
