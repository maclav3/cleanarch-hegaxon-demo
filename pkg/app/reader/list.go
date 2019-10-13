package reader

import "github.com/maclav3/cleanarch-hegaxon-demo/pkg/domain/reader"

type ListQuery struct {
}

func (r *registry) List(q ListQuery) ([]*reader.Reader, error) {
	return nil, nil
}
