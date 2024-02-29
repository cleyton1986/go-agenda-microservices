package entities

import "time"

type Contact struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewContact(name, surname, email string) *Contact {
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

	return contact
}


func (c*Contact) IsValid() error {


	if (len(c.Name) < 5){
		return fmt.Error("O nome não pode ter menos de 5 caracter. nome só têm %d", len(c.Name))
	}

	return nil
}


