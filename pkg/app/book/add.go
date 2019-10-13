package book

import (
	"github.com/maclav3/cleanarch-hegaxon-demo/pkg/domain/book"
	"github.com/pkg/errors"
)

type AddBook struct {
	ID     book.ID
	Author string
	Title  string
}

func (cmd AddBook) validate() error {
	// we could validate some data on the application layer.
	// the domain layer should prohibit any actions that would violate the domain rules.
	return nil
}

func (b *inventory) AddBook(cmd AddBook) error {
	if err := cmd.validate(); err != nil {
		return errors.Wrap(err, "invalid command")
	}
	return nil
}
