package book

import (
	"github.com/maclav3/cleanarch-hegaxon-demo/pkg/domain/book"
	"github.com/maclav3/cleanarch-hegaxon-demo/pkg/domain/reader"
	"github.com/pkg/errors"
)

type LoanBook struct {
	BookID   book.ID
	ReaderID reader.ID
}

func (cmd LoanBook) validate() error {
	// we could validate some data on the application layer.
	// the domain layer should prohibit any actions that would violate the domain rules.
	return nil
}

// LoanBook retrieves a book and a reader and makes the loan.
// Note that the app layer mostly orchestrates retrieving the aggregates and calling the domain methods;
// It is the domain layer that says which actions are allowed and which aren't.
func (i *inventory) LoanBook(cmd LoanBook) error {
	if err := cmd.validate(); err != nil {
		return errors.Wrap(err, "invalid command")
	}

	b, err := i.bookRepo.ByID(cmd.BookID)
	if err != nil {
		return errors.Wrap(err, "could not find book by ID")
	}

	r, err := i.readerRepo.ByID(cmd.ReaderID)
	if err != nil {
		return errors.Wrap(err, "could not find reader by ID")
	}

	err = b.Loan(r)
	if err != nil {
		return errors.Wrap(err, "could not loan book for reader")
	}

	err = i.bookRepo.Save(b)
	if err != nil {
		return errors.Wrap(err, "could not save book")
	}

	return nil
}
