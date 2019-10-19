package book

import (
	"github.com/maclav3/cleanarch-hegaxon-demo/internal/log"
	"github.com/maclav3/cleanarch-hegaxon-demo/pkg/domain/book"
	"github.com/pkg/errors"
)

var (
	// ErrBookAlreadyExists occurs when trying to add an already existing book to the inventory.
	ErrBookAlreadyExists = errors.New("trying to add a book that already exists")
)

type AddBookCommandHandler interface {
	Handle(cmd Add) error
}

type addBookCommandHandler struct {
	repo book.Repository
}

func NewAddBookCommandHandler(logger log.Logger, repo book.Repository) AddBookCommandHandler {
	// we panic if any dependency is nil
	if logger == nil {
		panic("logger is nil")
	}
	// because it is not a recoverable state
	// and should be fixed in compile time
	if repo == nil {
		panic("repo is nil")
	}

	return &addBookCommandHandlerLogger{
		logger: logger,
		wrapped: &addBookCommandHandler{
			repo: repo,
		},
	}
}

type Add struct {
	ID     book.ID
	Author string
	Title  string
}

func (cmd Add) validate() error {
	// we perform some simple data validation on the application layer.
	// however, it is the responsibility of the  domain layer
	// to should prohibit any actions that would violate the domain rules.
	if cmd.ID.Empty() {
		return errors.New("book id is empty")
	}

	return nil
}

// Add adds a new book to the inventory.
// Notice that the repository exposes ByID and Save, but this use case makes sure
// that an error is returned if we try to add a book that already exists.
func (i *addBookCommandHandler) Handle(cmd Add) error {
	if err := cmd.validate(); err != nil {
		return errors.Wrap(err, "invalid command")
	}

	_, err := i.repo.ByID(cmd.ID)
	if err == nil {
		return errors.Wrap(ErrBookAlreadyExists, cmd.ID.String())
	}
	if err != book.ErrNotFound {
		return errors.Wrap(err, "error checking if book already exists")
	}

	// book not found, carry on
	newBook, err := book.NewBook(cmd.ID, cmd.Author, cmd.Title)
	if err != nil {
		return errors.Wrap(err, "could not create a new book")
	}

	err = i.repo.Save(newBook)
	if err != nil {
		return errors.Wrap(err, "could not save a new book")
	}

	return nil
}
