package book

import (
	"github.com/maclav3/cleanarch-hegaxon-demo/internal/log"
	"github.com/maclav3/cleanarch-hegaxon-demo/pkg/domain/book"
	"github.com/maclav3/cleanarch-hegaxon-demo/pkg/domain/reader"
	"github.com/pkg/errors"
)

type LoanBookCommandHandler interface {
	Handle(cmd Loan) error
}

type loanBookCommandHandler struct {
	bookRepo   book.Repository
	readerRepo readerRepo
}

type readerRepo interface {
	// reader.Repository interface provides the basic reader repository methods.
	ByID(reader.ID) (*reader.Reader, error)
}

func NewLoanBookCommandHandler(logger log.Logger, bookRepo book.Repository, readerRepo readerRepo) LoanBookCommandHandler {
	// we panic if any dependency is nil
	if logger == nil {
		panic("logger is nil")
	}
	// because it is not a recoverable state
	if bookRepo == nil {
		panic("bookRepo is nil")
	}

	// and should be fixed in compile time
	if readerRepo == nil {
		panic("readerRepo is nil")
	}

	return &loanBookCommandHandlerLogger{
		logger: logger,
		wrapped: &loanBookCommandHandler{
			bookRepo:   bookRepo,
			readerRepo: readerRepo,
		},
	}
}

type Loan struct {
	BookID   book.ID
	ReaderID reader.ID
}

func (cmd Loan) validate() error {
	// we perform some simple data validation on the application layer.
	// however, it is the responsibility of the  domain layer
	// to should prohibit any actions that would violate the domain rules.
	if cmd.BookID.Empty() {
		return errors.New("book id is empty")
	}

	if cmd.ReaderID.Empty() {
		return errors.New("reader id is empty")
	}

	return nil
}

// Loan retrieves a book and a reader and makes the loan.
// Note that the app layer mostly orchestrates retrieving the aggregates and calling the domain methods;
// It is the domain layer that says which actions are allowed and which aren't.
func (h *loanBookCommandHandler) Handle(cmd Loan) error {
	if err := cmd.validate(); err != nil {
		return errors.Wrap(err, "invalid command")
	}

	b, err := h.bookRepo.ByID(cmd.BookID)
	if err != nil {
		return errors.Wrap(err, "could not find book by ID")
	}

	r, err := h.readerRepo.ByID(cmd.ReaderID)
	if err != nil {
		return errors.Wrap(err, "could not find reader by ID")
	}

	err = b.Loan(r)
	if err != nil {
		return errors.Wrap(err, "could not loan book for reader")
	}

	err = h.bookRepo.Save(b)
	if err != nil {
		return errors.Wrap(err, "could not save book")
	}

	return nil
}
