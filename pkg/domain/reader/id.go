package reader

// ID is an unique identifier for a Reader.
type ID string

func (i ID) String() string {
	return string(i)
}

func (i ID) Empty() bool {
	return string(i) == ""
}
