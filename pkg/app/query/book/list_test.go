package book_test

import (
	"testing"

	"github.com/maclav3/cleanarch-hegaxon-demo/internal/log"
	"github.com/maclav3/cleanarch-hegaxon-demo/pkg/adapters/book"
	bookQuery "github.com/maclav3/cleanarch-hegaxon-demo/pkg/app/query/book"
	bookDomain "github.com/maclav3/cleanarch-hegaxon-demo/pkg/domain/book"
	"github.com/maclav3/cleanarch-hegaxon-demo/pkg/domain/book/test"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_ListBooks(t *testing.T) {
	bookRepo := book.NewMemoryRepository()
	queryHandler := bookQuery.NewListBooksQueryHandler(log.NewNoopLogger(), bookRepo)

	unloanedBook := test.NewBook(t)
	require.NoError(t, bookRepo.Save(unloanedBook))
	loanedBook := test.NewLoanedBook(t)
	require.NoError(t, bookRepo.Save(loanedBook))

	testCases := []struct {
		name            string
		query           bookQuery.ListQuery
		expectedBookIDs []bookDomain.ID
	}{
		{
			name: "no_filters",
			query: bookQuery.ListQuery{
				Loaned: nil,
			},
			expectedBookIDs: []bookDomain.ID{
				unloanedBook.ID(),
				loanedBook.ID(),
			},
		},
		{
			name: "only_loaned",
			query: bookQuery.ListQuery{
				Loaned: ptrBool(true),
			},
			expectedBookIDs: []bookDomain.ID{
				loanedBook.ID(),
			},
		},
		{
			name: "only_unloaned",
			query: bookQuery.ListQuery{
				Loaned: ptrBool(false),
			},
			expectedBookIDs: []bookDomain.ID{
				unloanedBook.ID(),
			},
		},
	}

	for i := range testCases {
		t.Run(testCases[i].name, func(t *testing.T) {
			t.Parallel()
			tc := testCases[i]

			books, err := queryHandler.Query(tc.query)
			require.NoError(t, err)

			bookIDs := make([]bookDomain.ID, len(books))
			for j := range books {
				bookIDs[j] = books[j].ID()
			}

			assert.EqualValues(t, tc.expectedBookIDs, bookIDs)
		})
	}
}

func ptrBool(b bool) *bool {
	return &b
}
