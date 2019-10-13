package book_test

import (
	"testing"

	"github.com/maclav3/cleanarch-hegaxon-demo/pkg/adapters/book"
	"github.com/maclav3/cleanarch-hegaxon-demo/pkg/domain/book/test"
)

func TestMemoryRepository(t *testing.T) {
	repo := book.NewMemoryRepository()
	test.RepositoryTests(t, repo)
}
