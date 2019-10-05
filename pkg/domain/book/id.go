package book

// ID is an unique identifier for a Book.
// A separate ID type helps enforce that proper objects are supplied to functions
// and, for example, a Reader ID is not passed by mistake.
// This is optional, and it is also OK to use plain strings for domain entity IDs.
type ID string

func (i ID) String() string {
	return string(i)
}

func (i ID) Empty() bool {
	return string(i) == ""
}
