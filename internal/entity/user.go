package entity

import (
	"github.com/google/uuid"
	"github.com/savioafs/web-scraper-collector-go/internal/common"
	"time"
)

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func NewUser(name, email string) (*User, error) {
	u := &User{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
	}

	if err := u.Validate(); err != nil {
		return nil, err
	}

	return u, nil
}

func (u *User) Validate() error {
	if u.Name == "" {
		return common.ErrNameIsRequired
	}

	if u.Email == "" {
		return common.ErrEmailIsRequired
	}

	return nil
}
