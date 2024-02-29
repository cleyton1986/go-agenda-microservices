package use_case

import (
	"log"

	"github.com/cleyton1986/go-agenda-microservices/cmd/internal/contacts/models"
)

type createContactUseCase struct {
	// instance DB

}

func NewCreateContactUseCase() *createContactUseCase {
	return &createContactUseCase{}

}

func (u *createContactUseCase) Execute(name, surname, email string) error {
	contact, err := models.NewContact(name, surname, email)

	if err != nil {
		return err
	}

	log.Println(contact)

	return nil
}