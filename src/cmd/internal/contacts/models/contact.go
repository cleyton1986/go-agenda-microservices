package models

import (
	"fmt"
	"time"
)

type contact struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewContact(name, surname, email string) (*contact, error) {
	contact := &contact{
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

func (c *contact) IsValid() error {


	if len(c.Name) < 5 {
		return fmt.Errorf("nome não pode ter menos de 5 caracteres. O nome tem apenas %d caracteres", len(c.Name))
	}



	return nil
}
