package reader

import (
	"github.com/maclav3/cleanarch-hegaxon-demo/pkg/domain/reader"
	"github.com/pkg/errors"
)

type Deactivate struct {
	ID reader.ID
}

func (cmd Deactivate) validate() error {
	if cmd.ID.Empty() {
		return errors.New("reader id is empty")
	}
	return nil
}

func (r *registry) Deactivate(cmd Deactivate) error {
	err := cmd.validate()
	if err != nil {
		return errors.Wrap(err, "invalid command")
	}

	rdr, err := r.readerRepo.ByID(cmd.ID)
	if err != nil {
		return errors.Wrap(err, "could not retrieve reader by ID")
	}

	err = rdr.Deactivate()
	if err != nil {
		return errors.Wrap(err, "could not deactivate reader")
	}

	err = r.readerRepo.Save(rdr)
	if err != nil {
		return errors.Wrap(err, "could not persist reader")
	}

	return nil
}
