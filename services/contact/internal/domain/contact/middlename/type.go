package middlename

import "github.com/pkg/errors"

var (
	MaxLength      = 100
	ErrWrongLength = errors.Errorf("patronymic must be less than or equal to %d characters", MaxLength)
)

type Middlename string

func (p Middlename) String() string {
	return string(p)
}

func New(middlename string) (*Middlename, error) {
	if len([]rune(middlename)) > MaxLength {
		return nil, ErrWrongLength
	}
	p := Middlename(middlename)
	return &p, nil
}
