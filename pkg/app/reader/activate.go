package reader

import (
	"github.com/maclav3/cleanarch-hegaxon-demo/pkg/domain/reader"
	"github.com/pkg/errors"
)

type Activate struct {
	ID reader.ID
}

func (r *registry) Activate(cmd Activate) error {
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
