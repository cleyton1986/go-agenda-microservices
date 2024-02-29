package models

import (
	"fmt"
	"time"

	errpkg "github.com/cleyton1986/go-agenda-microservices/src/cmd/pkg/error"
)

type Contact struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewContact(name, surname, email string) (*Contact, error) {
	contact := &Contact{
		Name:      name,
		Surname:   surname,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := contact.IsValid()

	if err != nil {
		return nil, err
	}

	return contact, nil
}

func (c *Contact) IsValid() error {
	ec := errpkg.NewErrorCollection()

	if len(c.Name) < 5 {
		ec.Add(fmt.Errorf("nome não pode ter menos de 5 caracteres. O nome tem apenas %d caracteres", len(c.Name)))
	}

	if ec.HasErrors() {
		return ec.Throw()
	}

	return nil
}
