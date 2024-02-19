package contact

import (
	"architecture_go/services/contact/internal/domain/contact/phoneNumber"
	"fmt"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"architecture_go/services/contact/internal/domain/contact/firstname"
	"architecture_go/services/contact/internal/domain/contact/middlename"
	"architecture_go/services/contact/internal/domain/contact/surname"
)

var (
	ErrPhoneNumberRequired = errors.New("phone number is required")
)

type Contact struct {
	id          uuid.UUID
	phoneNumber phoneNumber.PhoneNumber
	firstname   firstname.Firstname
	surname     surname.Surname
	middlename  middlename.Middlename
}

func NewWithID(
	id uuid.UUID,
	phoneNumber phoneNumber.PhoneNumber,
	firstname firstname.Firstname,
	surname surname.Surname,
	middlename middlename.Middlename,
) (*Contact, error) {

	if phoneNumber.IsEmpty() {
		return nil, ErrPhoneNumberRequired
	}

	if id == uuid.Nil {
		id = uuid.New()
	}

	return &Contact{
		id:          id,
		phoneNumber: phoneNumber,
		firstname:   firstname,
		surname:     surname,
		middlename:  middlename,
	}, nil
}

func New(
	phoneNumber phoneNumber.PhoneNumber,
	firstname firstname.Firstname,
	surname surname.Surname,
	middlename middlename.Middlename,
) (*Contact, error) {

	if phoneNumber.IsEmpty() {
		return nil, ErrPhoneNumberRequired
	}

	return &Contact{
		id:          uuid.New(),
		phoneNumber: phoneNumber,
		firstname:   firstname,
		surname:     surname,
		middlename:  middlename,
	}, nil
}

func (c Contact) ID() uuid.UUID {
	return c.id
}

func (c Contact) PhoneNumber() phoneNumber.PhoneNumber {
	return c.phoneNumber
}

func (c Contact) Name() firstname.Firstname {
	return c.firstname
}

func (c Contact) Surname() surname.Surname {
	return c.surname
}

func (c Contact) Patronymic() middlename.Middlename {
	return c.middlename
}

func (c Contact) FullName() string {
	return fmt.Sprintf("%s %s %s", c.surname, c.firstname, c.middlename)
}

func (c Contact) Equal(contact Contact) bool {
	return c.id == contact.id
}
