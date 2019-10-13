package book

import (
	"time"

	"github.com/maclav3/cleanarch-hegaxon-demo/pkg/domain/reader"
	"github.com/pkg/errors"
)

var (
	// ErrAlreadyLoaned occurs when somebody tries to loan a book that is already loaned.
	ErrAlreadyLoaned = errors.New("book is already loaned")
	// ErrInactiveReader occurs when an inactive reader tries to loan a book.
	ErrInactiveReader = errors.New("inactive readers cannot loan books")
)

// Book represents a book asset that may be loaned by a Reader.
// note: the fields are encapsulated. They may be accessed by readers,
// and changed only via dedicated methods that conform to the domain invariants.
type Book struct {
	id ID

	author string
	title  string

	loan *Loan
}

// NewBook returns a new Book aggregate.
// The constructor is the only way for outside packages to create a Book,
// which means that the `domain/book` package has exclusive control over
// the contents of a Book struct, which is needed to be sure that
// the business invariants are enforced at all times.
func NewBook(id ID, author string, title string) (*Book, error) {
	if id.Empty() {
		return nil, errors.New("book id is empty")
	}
	if author == "" {
		// in the future, we could require a proper Author object with more details.
		// for now, let's just keep it a plain string
		return nil, errors.New("author is empty")
	}
	if title == "" {
		return nil, errors.New("title is empty")
	}

	// A new book is by default in a non-loaned state.
	// Book.Loan() must be called in order to mark the book as loaned.
	return &Book{
		id:     id,
		author: author,
		title:  title,
		loan:   nil,
	}, nil
}

func (b *Book) ID() ID {
	return b.id
}

func (b *Book) Author() string {
	return b.author
}

func (b *Book) Title() string {
	return b.title
}

func (b Book) Loaned() bool {
	return b.loan != nil
}

func (b *Book) Loan(forReader *reader.Reader) error {
	// domain rules go here
	// a book cannot be loaned when it is already loaned to somebody
	if b.Loaned() {
		return ErrAlreadyLoaned
	}

	if !forReader.Active() {
		return ErrInactiveReader
	}

	from := time.Now()
	to := from.Add(7 * 24 * time.Hour)
	loan := &Loan{
		from: from,
		// TODO: have a domain service with injected loan period
		to: to,
		// In DDD, we don't store one aggregate within another.
		// we store only a reference by ID.
		loanedBy: forReader.ID(),
	}

	b.loan = loan
	return nil
}
