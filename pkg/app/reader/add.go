package reader

import (
	"github.com/maclav3/cleanarch-hegaxon-demo/pkg/domain/reader"
	"github.com/pkg/errors"
)

var (
	// ErrReaderAlreadyExists occurs when trying to add an already existing reader to the inventory.
	ErrReaderAlreadyExists = errors.New("trying to add a reader that already exists")
)

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

func (r *registry) Add(cmd Add) error {
	if err := cmd.validate(); err != nil {
		return errors.Wrap(err, "invalid command")
	}

	_, err := r.readerRepo.ByID(cmd.ID)
	if err == nil {
		return errors.Wrap(ErrReaderAlreadyExists, cmd.ID.String())
	}

	newReader, err := reader.NewReader(cmd.ID, cmd.Name)
	if err != nil {
		return errors.Wrap(err, "could not create a new reader")
	}

	err = r.readerRepo.Save(newReader)
	if err != nil {
		return errors.Wrap(err, "could not save the new reader to repository")
	}

	return nil
}