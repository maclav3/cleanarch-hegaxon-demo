package book_test

import (
	"testing"

	"github.com/maclav3/cleanarch-hegaxon-demo/internal/log"
	"github.com/maclav3/cleanarch-hegaxon-demo/pkg/adapters/book"
	"github.com/maclav3/cleanarch-hegaxon-demo/pkg/adapters/reader"
	appBook "github.com/maclav3/cleanarch-hegaxon-demo/pkg/app/book"
	domainBook "github.com/maclav3/cleanarch-hegaxon-demo/pkg/domain/book"
	"github.com/maclav3/cleanarch-hegaxon-demo/pkg/domain/book/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMemoryRepository_ListBooks(t *testing.T) {
	bookRepo := book.NewMemoryRepository()
	readerRepo := reader.NewMemoryRepository()
	inventory := appBook.NewBookInventory(log.NewNoopLogger(), bookRepo, readerRepo)

	unloanedBook := test.NewBook(t)
	require.NoError(t, bookRepo.Save(unloanedBook))
	loanedBook := test.NewLoanedBook(t)
	require.NoError(t, bookRepo.Save(loanedBook))

	testCases := []struct {
		name            string
		query           appBook.ListQuery
		expectedBookIDs []domainBook.ID
	}{
		{
			name: "no_filters",
			query: appBook.ListQuery{
				Loaned: nil,
			},
			expectedBookIDs: []domainBook.ID{
				unloanedBook.ID(),
				loanedBook.ID(),
			},
		},
		{
			name: "only_loaned",
			query: appBook.ListQuery{
				Loaned: ptrBool(true),
			},
			expectedBookIDs: []domainBook.ID{
				loanedBook.ID(),
			},
		},
		{
			name: "only_unloaned",
			query: appBook.ListQuery{
				Loaned: ptrBool(false),
			},
			expectedBookIDs: []domainBook.ID{
				unloanedBook.ID(),
			},
		},
	}

	for i := range testCases {
		t.Run(testCases[i].name, func(t *testing.T) {
			t.Parallel()
			tc := testCases[i]

			books, err := inventory.List(tc.query)
			require.NoError(t, err)

			bookIDs := make([]domainBook.ID, len(books))
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
