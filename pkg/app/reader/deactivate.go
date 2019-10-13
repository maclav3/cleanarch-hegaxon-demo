package reader

import "github.com/maclav3/cleanarch-hegaxon-demo/pkg/domain/reader"

type Deactivate struct {
	ID reader.ID
}

func (r *registry) Deactivate(cmd Deactivate) error {
	return nil
}
