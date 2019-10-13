package reader

import "github.com/maclav3/cleanarch-hegaxon-demo/pkg/domain/reader"

type memoryRepository struct {
	readers map[reader.ID]*reader.Reader
}

func NewMemoryRepository() reader.Repository {
	return &memoryRepository{
		readers: map[reader.ID]*reader.Reader{},
	}
}

func (m *memoryRepository) ByID(id reader.ID) (*reader.Reader, error) {
	r, ok := m.readers[id]
	if !ok {
		return nil, reader.ErrNotFound
	}
	return r, nil
}

func (m *memoryRepository) Save(r *reader.Reader) error {
	m.readers[r.ID()] = r
	return nil
}
