package entity

import (
	"errors"
	"github.com/google/uuid"
)

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var (
	ErrEmailIsRequired = errors.New("email is required")
)

func NewUser(name, email string) (*User, error) {
	u := &User{
		ID:    uuid.New().String(),
		Name:  name,
		Email: email,
	}

	if err := u.Validate(); err != nil {
		return nil, err
	}

	return u, nil
}

func (u *User) Validate() error {
	if u.Name == "" {
		return ErrNameIsRequired
	}

	if u.Email == "" {
		return ErrEmailIsRequired
	}

	return nil
}
