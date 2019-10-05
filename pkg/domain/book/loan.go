package book

import (
	// it is OK to import another domain package in the domain layer
	"github.com/maclav3/cleanarch-hegaxon-demo/pkg/domain/reader"
	"time"
)

// Loan represents the act of a Reader loaning a Book for a specified time.
// Loan is a Value Object, which means that it is immutable once created.
// There is no public construtor for Loan, because these objects may be created
// only through a call of Book.Loan().
type Loan struct {
	from time.Time
	to time.Time
	loanedBy reader.ID
}

