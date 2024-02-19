package firstname

import "github.com/pkg/errors"

var (
	MaxLength      = 50
	ErrWrongLength = errors.Errorf("name must be less than or equal to %d characters", MaxLength)
)

type Firstname string

func (n Firstname) String() string {
	return string(n)
}

func New(firstname string) (*Firstname, error) {
	if len([]rune(firstname)) > MaxLength {
		return nil, ErrWrongLength
	}
	n := Firstname(firstname)
	return &n, nil
}
