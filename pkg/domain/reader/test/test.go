package test

import (
	"testing"

	"github.com/gofrs/uuid"
	"github.com/maclav3/cleanarch-hegaxon-demo/pkg/domain/reader"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// RepositoryTests defines the minimal set of tests that every implementation of Repository must pass.
// Make sure that the repository passed into this test is empty.
func RepositoryTests(t *testing.T, repo reader.Repository) {
	t.Run("found_not_found", func(t *testing.T) {
		r := NewActiveReader(t)
		_, err := repo.ByID(r.ID())
		require.Equal(t, reader.ErrNotFound, errors.Cause(err), "expected that reader is not found")

		err = repo.Save(r)
		require.NoError(t, err, "expected no error saving the reader")

		readerFromRepo, err := repo.ByID(r.ID())
		require.NoError(t, err, "expected no error retrieving reader from repository")

		assert.Equal(t, r.ID(), readerFromRepo.ID())
	})

	t.Run("save_updates", func(t *testing.T) {
		r := NewActiveReader(t)

		err := repo.Save(r)
		require.NoError(t, err, "expected no error saving the reader")

		readerFromRepo, err := repo.ByID(r.ID())
		require.NoError(t, err, "expected no error retrieving reader from repository")
		require.True(t, readerFromRepo.Active(), "expected the new reader to be active")

		err = readerFromRepo.Deactivate()
		require.NoError(t, err, "expected no error deactivating the reader")

		err = repo.Save(readerFromRepo)
		require.NoError(t, err, "expected no error saving the reader")

		readerFromRepo, err = repo.ByID(readerFromRepo.ID())
		require.NoError(t, err, "expected no error retrieving reader from repository")
		require.False(t, readerFromRepo.Active(), "expected that the reader is inactive now")
	})
}

func NewActiveReader(t *testing.T) *reader.Reader {
	readerID := reader.ID(uuid.Must(uuid.NewV4()).String())
	r, err := reader.NewReader(readerID, "Adam Miauczynski")
	require.NoError(t, err, "expected no error creating a new reader aggregate")

	err = r.Activate()
	require.NoError(t, err, "expected no error activating the reader")

	return r
}
