package reader

import (
	"github.com/maclav3/cleanarch-hegaxon-demo/internal/log"
	"github.com/maclav3/cleanarch-hegaxon-demo/pkg/domain/reader"
	"github.com/pkg/errors"
)

type DeactivateReaderCommandHandler interface {
	Handle(cmd Deactivate) error
}

type deactivateReaderCommandHandler struct {
	repo reader.Repository
}

func NewDeactivateReaderCommandHandler(logger log.Logger, repo reader.Repository) DeactivateReaderCommandHandler {
	// we panic if any dependency is nil
	if logger == nil {
		panic("logger is nil")
	}
	// because it is not a recoverable state
	// and should be fixed in compile time
	if repo == nil {
		panic("readerRepo is nil")
	}

	return &deactivateReaderCommandHandlerLogger{
		logger: logger,
		wrapped: &deactivateReaderCommandHandler{
			repo: repo,
		},
	}
}

type Deactivate struct {
	ID reader.ID
}

func (cmd Deactivate) validate() error {
	if cmd.ID.Empty() {
		return errors.New("reader id is empty")
	}
	return nil
}

func (h *deactivateReaderCommandHandler) Handle(cmd Deactivate) error {
	err := cmd.validate()
	if err != nil {
		return errors.Wrap(err, "invalid command")
	}

	rdr, err := h.repo.ByID(cmd.ID)
	if err != nil {
		return errors.Wrap(err, "could not retrieve reader by ID")
	}

	err = rdr.Deactivate()
	if err != nil {
		return errors.Wrap(err, "could not deactivate reader")
	}

	err = h.repo.Save(rdr)
	if err != nil {
		return errors.Wrap(err, "could not persist reader")
	}

	return nil
}
