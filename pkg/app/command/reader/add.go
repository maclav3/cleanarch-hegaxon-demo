package reader

import (
	"github.com/maclav3/cleanarch-hegaxon-demo/internal/log"
	"github.com/maclav3/cleanarch-hegaxon-demo/pkg/domain/reader"
	"github.com/pkg/errors"
)

var (
	// ErrReaderAlreadyExists occurs when trying to add an already existing reader to the inventory.
	ErrReaderAlreadyExists = errors.New("trying to add a reader that already exists")
)

type AddReaderCommandHandler interface {
	Handle(cmd Add) error
}

type addReaderCommandHandler struct {
	repo reader.Repository
}

func NewAddReaderCommandHandler(logger log.Logger, repo reader.Repository) AddReaderCommandHandler {
	// we panic if any dependency is nil
	if logger == nil {
		panic("logger is nil")
	}
	// because it is not a recoverable state
	// and should be fixed in compile time
	if repo == nil {
		panic("readerRepo is nil")
	}

	return &addReaderCommandHandlerLogger{
		logger: logger,
		wrapped: &addReaderCommandHandler{
			repo: repo,
		},
	}
}

type Add struct {
	ID   reader.ID
	Name string
}

func (cmd Add) validate() error {
	// we perform some simple data validation on the application layer.
	// however, it is the responsibility of the  domain layer
	// to should prohibit any actions that would violate the domain rules.
	if cmd.ID.Empty() {
		return errors.New("reader id is empty")
	}
	return nil
}

func (h *addReaderCommandHandler) Handle(cmd Add) error {
	if err := cmd.validate(); err != nil {
		return errors.Wrap(err, "invalid command")
	}

	_, err := h.repo.ByID(cmd.ID)
	if err == nil {
		return errors.Wrap(ErrReaderAlreadyExists, cmd.ID.String())
	}

	newReader, err := reader.NewReader(cmd.ID, cmd.Name)
	if err != nil {
		return errors.Wrap(err, "could not create a new reader")
	}

	err = h.repo.Save(newReader)
	if err != nil {
		return errors.Wrap(err, "could not save the new reader to repository")
	}

	return nil
}
