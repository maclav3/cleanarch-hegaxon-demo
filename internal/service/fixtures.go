package service

import (
	"fmt"

	"github.com/maclav3/cleanarch-hegaxon-demo/pkg/app/query/book"

	book2 "github.com/maclav3/cleanarch-hegaxon-demo/pkg/app/command/book"

	readerCommand "github.com/maclav3/cleanarch-hegaxon-demo/pkg/app/command/reader"
)

// AddFixtures adds some simple fixtures to make the demonstration app filled with some sample data.
func (s *Service) AddFixtures() {
	assertNoError(s.AddReaderCommandHandler.Handle(readerCommand.Add{
		ID:   "1",
		Name: "Reader1",
	}))
	assertNoError(s.ActivateReaderCommandHandler.Handle(readerCommand.Activate{
		ID: "1",
	}))
	assertNoError(s.AddReaderCommandHandler.Handle(readerCommand.Add{
		ID:   "2",
		Name: "Reader2",
	}))
	assertNoError(s.ActivateReaderCommandHandler.Handle(readerCommand.Activate{
		ID: "2",
	}))

	assertNoError(s.AddBookCommandHandler.Handle(book2.Add{
		ID:     "3",
		Author: "Author3",
		Title:  "Title3",
	}))
	assertNoError(s.AddBookCommandHandler.Handle(book2.Add{
		ID:     "4",
		Author: "Author4",
		Title:  "Title4",
	}))
	assertNoError(s.AddBookCommandHandler.Handle(book2.Add{
		ID:     "5",
		Author: "Author5",
		Title:  "Title5",
	}))

	assertNoError(s.LoanBookCommandHandler.Handle(book2.Loan{
		BookID:   "3",
		ReaderID: "1",
	}))
	assertNoError(s.LoanBookCommandHandler.Handle(book2.Loan{
		BookID:   "4",
		ReaderID: "2",
	}))

	loanedList, err := s.ListBooksQueryHandler.Query(book.ListQuery{
		Loaned: func(b bool) *bool { return &b }(true),
	})
	if err != nil {
		panic(err)
	}

	unloanedList, err := s.ListBooksQueryHandler.Query(book.ListQuery{
		Loaned: func(b bool) *bool { return &b }(false),
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("LOANED BOOKS:\n%+v\n\n", loanedList)
	fmt.Printf("UNLOANED BOOKS:\n%+v\n\n", unloanedList)
}

func assertNoError(err error) {
	if err != nil {
		panic(err)
	}
}
