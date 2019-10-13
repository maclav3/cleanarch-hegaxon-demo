package book

import (
	"github.com/maclav3/cleanarch-hegaxon-demo/pkg/domain/book"
	"github.com/pkg/errors"
)

var (
	// ErrBookAlreadyExists occurs when trying to add an already existing book to the inventory.
	ErrBookAlreadyExists = errors.New("trying to add a book that already exists")
)

type Add struct {
	ID     book.ID
	Author string
	Title  string
}

func (cmd Add) validate() error {
	// we could validate some data on the application layer.
	// the domain layer should prohibit any actions that would violate the domain rules.
	return nil
}

// Add adds a new book to the inventory.
// Notice that the repository exposes ByID and Save, but this use case makes sure
// that an error is returned if we try to add a book that already exists.
func (i *inventory) Add(cmd Add) error {
	if err := cmd.validate(); err != nil {
		return errors.Wrap(err, "invalid command")
	}

	id := book.ID(cmd.ID)

	_, err := i.bookRepo.ByID(id)
	if err == nil {
		return errors.Wrap(ErrBookAlreadyExists, cmd.ID.String())
	}
	if err != book.ErrNotFound {
		return errors.Wrap(err, "error checking if book already exists")
	}

	// book not found, carry on
	newBook, err := book.NewBook(id, cmd.Author, cmd.Title)
	if err != nil {
		return errors.Wrap(err, "could not create a new book")
	}

	err = i.bookRepo.Save(newBook)
	if err != nil {
		return errors.Wrap(err, "could not save a new book")
	}

	return nil
}
