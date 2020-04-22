package test

import (
	"testing"

	"github.com/maclav3/cleanarch-hegaxon-demo/pkg/domain/reader/test"

	"github.com/gofrs/uuid"
	"github.com/maclav3/cleanarch-hegaxon-demo/pkg/domain/book"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// RepositoryTests defines the minimal set of tests that every implementation of Repository must pass.
// Make sure that the repository passed into this test is empty.
func RepositoryTests(t *testing.T, repo book.Repository) {
	t.Run("found_not_found", func(t *testing.T) {
		b := NewBook(t)
		_, err := repo.ByID(b.ID())
		require.Equal(t, book.ErrNotFound, errors.Cause(err), "expected that book is not found")

		err = repo.Save(b)
		require.NoError(t, err, "expected no error saving the book")

		bookFromRepo, err := repo.ByID(b.ID())
		require.NoError(t, err, "expected no error retrieving book from repository")

		assert.Equal(t, b.ID(), bookFromRepo.ID())
	})

	t.Run("save_updates", func(t *testing.T) {
		// this test checks that the changes to the state of the book are persisted.
		// it is a cursory test, no sense in checking that every single property is saved.
		b := NewBook(t)
		r := test.NewActiveReader(t)

		err := repo.Save(b)
		require.NoError(t, err, "expected no error saving the book")

		bookFromRepo, err := repo.ByID(b.ID())
		require.NoError(t, err, "expected no error retrieving book from repository")
		require.False(t, bookFromRepo.Loaned(), "did not expect a new book to be loaned")

		err = bookFromRepo.Loan(r, book.DefaultLoanPeriod)
		require.NoError(t, err, "expected no error loaning the book")

		err = repo.Save(bookFromRepo)
		require.NoError(t, err, "expected no error saving the book")

		bookFromRepo, err = repo.ByID(bookFromRepo.ID())
		require.NoError(t, err, "expected no error retrieving book from repository")
		require.True(t, bookFromRepo.Loaned(), "expected that the book is loaned now")
	})
}

func NewBook(t *testing.T) *book.Book {
	bookID := book.ID(uuid.Must(uuid.NewV4()).String())
	b, err := book.NewBook(bookID, "F. M. Dostoyevsky", "The Brothers Karamazov")
	require.NoError(t, err, "expected no error creating a new book aggregate")

	return b
}

func NewLoanedBook(t *testing.T) *book.Book {
	bookID := book.ID(uuid.Must(uuid.NewV4()).String())
	b, err := book.NewBook(bookID, "A. Mickiewicz", "Pan Tadeusz")
	require.NoError(t, err, "expected no error creating a new book aggregate")

	r := test.NewActiveReader(t)
	err = b.Loan(r, book.DefaultLoanPeriod)
	require.NoError(t, err)
	return b
}
