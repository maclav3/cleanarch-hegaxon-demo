package book_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/maclav3/cleanarch-hegaxon-demo/pkg/domain/book"
)

func TestNewBook(t *testing.T) {
	testCases := []struct {
		name      string
		id        book.ID
		author    string
		title     string
		expectErr bool
	}{
		{
			name:      "valid",
			id:        book.ID("1234"),
			author:    "A. Smith",
			title:     "The Wealth of Nations",
			expectErr: false,
		},
		{
			name:      "invalid_empty_id",
			id:        "",
			author:    "A. Smith",
			title:     "The Wealth of Nations",
			expectErr: true,
		},
		{
			name:      "invalid_empty_author",
			id:        book.ID("1234"),
			author:    "",
			title:     "The Wealth of Nations",
			expectErr: true,
		},
		{
			name:      "invalid_empty_title",
			id:        book.ID("1234"),
			author:    "A. Smith",
			title:     "",
			expectErr: true,
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			b, err := book.NewBook(tc.id, tc.author, tc.title)
			if tc.expectErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			assert.Equal(t, tc.id, b.ID())
			assert.Equal(t, tc.author, b.Author())
			assert.Equal(t, tc.title, b.Title())
			assert.False(t, b.Loaned(), "expected a new book not to be loaned")
		})
	}
}
