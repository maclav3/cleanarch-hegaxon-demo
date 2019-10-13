package reader_test

import (
	"testing"

	"github.com/maclav3/cleanarch-hegaxon-demo/pkg/adapters/reader"
	"github.com/maclav3/cleanarch-hegaxon-demo/pkg/domain/reader/test"
)

func TestMemoryRepository(t *testing.T) {
	repo := reader.NewMemoryRepository()
	test.RepositoryTests(t, repo)
}
