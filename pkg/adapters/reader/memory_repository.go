package reader

import (
	query "github.com/maclav3/cleanarch-hegaxon-demo/pkg/app/query/reader"
	"github.com/maclav3/cleanarch-hegaxon-demo/pkg/domain/reader"
)

type MemoryRepository struct {
	readers map[reader.ID]*reader.Reader
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		readers: map[reader.ID]*reader.Reader{},
	}
}

func (m *MemoryRepository) ByID(id reader.ID) (*reader.Reader, error) {
	r, ok := m.readers[id]
	if !ok {
		return nil, reader.ErrNotFound
	}
	return r, nil
}

func (m *MemoryRepository) Save(r *reader.Reader) error {
	m.readers[r.ID()] = r
	return nil
}

func (m *MemoryRepository) ListReaders(q query.ListQuery) ([]*reader.Reader, error) {
	all := []*reader.Reader{}
	for _, r := range m.readers {
		all = append(all, r)
	}
	return all, nil
}
