package reader

import (
	"github.com/maclav3/cleanarch-hegaxon-demo/pkg/domain/reader"
	"github.com/pkg/errors"
)

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

func (r *registry) Activate(cmd Activate) error {
	err := cmd.validate()
	if err != nil {
		return errors.Wrap(err, "invalid command")
	}

	rdr, err := r.readerRepo.ByID(cmd.ID)
	if err != nil {
		return errors.Wrap(err, "could not retrieve reader by ID")
	}

	err = rdr.Activate()
	if err != nil {
		return errors.Wrap(err, "could not activate reader")
	}

	err = r.readerRepo.Save(rdr)
	if err != nil {
		return errors.Wrap(err, "could not persist reader")
	}

	return nil
}
