package reader_test

import (
	"testing"

	"github.com/maclav3/cleanarch-hegaxon-demo/pkg/domain/reader"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewReader(t *testing.T) {
	testCases := []struct {
		testName  string
		id        reader.ID
		name      string
		expectErr bool
	}{
		{
			testName:  "valid",
			id:        reader.ID("1234"),
			name:      "K. I. Verkhovtseva",
			expectErr: false,
		},
		{
			testName:  "invalid_empty_id",
			id:        "",
			name:      "K. I. Verkhovtseva",
			expectErr: true,
		},
		{
			testName:  "invalid_empty_name",
			id:        reader.ID("1234"),
			name:      "",
			expectErr: true,
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.testName, func(t *testing.T) {
			t.Parallel()
			r, err := reader.NewReader(tc.id, tc.name)
			if tc.expectErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			assert.Equal(t, tc.id, r.ID())
			assert.Equal(t, tc.name, r.Name())
			assert.False(t, r.Active(), "expected a new reader not to be active")
		})
	}
}
