package test

import (
	"testing"

	"github.com/maclav3/cleanarch-hegaxon-demo/pkg/domain/reader"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/require"
)

func NewActiveReader(t *testing.T) *reader.Reader {
	readerID := reader.ID(uuid.Must(uuid.NewV4()).String())
	r, err := reader.NewReader(readerID, "Adam Miauczynski")
	require.NoError(t, err, "expected no error creating a new reader aggregate")

	err = r.Activate()
	require.NoError(t, err, "expected no error activating the reader")

	return r
}
