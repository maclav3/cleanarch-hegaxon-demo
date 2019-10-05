package reader

import "github.com/pkg/errors"

// Reader represents a Reader person that may loan and return books.
type Reader struct {
	id ID
	name string

	active bool
}

func NewReader(id ID, name string) (*Reader, error) {
	if id.Empty() {
		return nil, errors.New("id is empty")
	}
	if name == "" {
		return nil, errors.New("name is empty")
	}

	return &Reader{
		id:     id,
		name:   name,
		// new Readers are inactive by default
		active: false,
	}, nil
}

func (r Reader) ID() ID {
	return r.id
}

func (r Reader) Name() string {
	return r.name
}

// Active tells if an Reader account is active.
// Inactive Readers cannot loan any books.
func (r Reader) Active() bool {
return r.active
}

// Activate makes a Reader account active,
// which means that it can loan books.
func (r *Reader) Activate() error {
	// enforce domain rules on which accounts may be activated here.
	// return non-nil error if there is some problem with that.
	r.active = true
	return nil
}

// Deactivate makes a Reader account inactive,
// which means that it can no longer loan books.
func (r *Reader) Deactivate() error {
	// enforce domain rules on which accounts may be deactivated here.
	// return non-nil error if there is some problem with that.
	r.active = false
	return nil
}
