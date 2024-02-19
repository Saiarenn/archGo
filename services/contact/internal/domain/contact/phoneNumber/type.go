package phoneNumber

import (
	"strings"
)

type PhoneNumber struct {
	value string
}

func (p PhoneNumber) String() string {
	return p.value
}

func New(phone string) *PhoneNumber {
	return &PhoneNumber{
		value: getNumbers(phone),
	}
}

func (p PhoneNumber) Equal(phoneNumber PhoneNumber) bool {
	return p.value == phoneNumber.value
}

func (p PhoneNumber) IsEmpty() bool {
	return len(strings.TrimSpace(p.value)) == 0
}

func getNumbers(input string) string {
	var number string

	for _, t := range input {
		if t >= 48 && t <= 57 {
			number += string(t)
		}
	}

	return number
}
