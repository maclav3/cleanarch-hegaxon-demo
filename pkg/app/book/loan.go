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

func (b *inventory) LoanBook(cmd LoanBook) error {
	if err := cmd.validate(); err != nil {
		return errors.Wrap(err, "invalid command")
	}
	return nil
}
