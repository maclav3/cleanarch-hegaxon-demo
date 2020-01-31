package app

import (
	bookCommand "github.com/maclav3/cleanarch-hegaxon-demo/pkg/app/command/book"
	readerCommand "github.com/maclav3/cleanarch-hegaxon-demo/pkg/app/command/reader"
	bookQuery "github.com/maclav3/cleanarch-hegaxon-demo/pkg/app/query/book"
	readerQuery "github.com/maclav3/cleanarch-hegaxon-demo/pkg/app/query/reader"
)

type Application struct {
	Queries  *Queries
	Commands *Commands
}

type Queries struct {
	ListBooks   bookQuery.ListBooksQueryHandler
	ListReaders readerQuery.ListReadersQueryHandler
}

type Commands struct {
	AddBook  bookCommand.AddBookCommandHandler
	LoanBook bookCommand.LoanBookCommandHandler

	AddReader        readerCommand.AddReaderCommandHandler
	ActivateReader   readerCommand.ActivateReaderCommandHandler
	DeactivateReader readerCommand.DeactivateReaderCommandHandler
}
