package book_test

import (
	"testing"

	adaptersBook "github.com/maclav3/cleanarch-hegaxon-demo/pkg/adapters/book"
	"github.com/maclav3/cleanarch-hegaxon-demo/pkg/domain/book/test"
)

func TestMemoryRepository(t *testing.T) {
	repo := adaptersBook.NewMemoryRepository()
	test.RepositoryTests(t, repo)
}
