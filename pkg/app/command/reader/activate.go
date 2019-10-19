package reader

import (
	"github.com/maclav3/cleanarch-hegaxon-demo/internal/log"
	"github.com/maclav3/cleanarch-hegaxon-demo/pkg/domain/reader"

	"github.com/pkg/errors"
)

type ActivateReaderCommandHandler interface {
	Handle(cmd Activate) error
}

type activateReaderCommandHandler struct {
	repo reader.Repository
}

func NewActivateReaderCommandHandler(logger log.Logger, repo reader.Repository) ActivateReaderCommandHandler {
	// we panic if any dependency is nil
	if logger == nil {
		panic("logger is nil")
	}
	// because it is not a recoverable state
	// and should be fixed in compile time
	if repo == nil {
		panic("readerRepo is nil")
	}

	return &activateReaderCommandHandlerLogger{
		logger: logger,
		wrapped: &activateReaderCommandHandler{
			repo: repo,
		},
	}
}

type Activate struct {
	ID reader.ID
}

func (cmd Activate) validate() error {
	// we perform some simple data validation on the application layer.
	// however, it is the responsibility of the  domain layer
	// to should prohibit any actions that would violate the domain rules.
	if cmd.ID.Empty() {
		return errors.New("reader id is empty")
	}
	return nil
}

func (h *activateReaderCommandHandler) Handle(cmd Activate) error {
	err := cmd.validate()
	if err != nil {
		return errors.Wrap(err, "invalid command")
	}

	rdr, err := h.repo.ByID(cmd.ID)
	if err != nil {
		return errors.Wrap(err, "could not retrieve reader by ID")
	}

	err = rdr.Activate()
	if err != nil {
		return errors.Wrap(err, "could not activate reader")
	}

	err = h.repo.Save(rdr)
	if err != nil {
		return errors.Wrap(err, "could not persist reader")
	}

	return nil
}
