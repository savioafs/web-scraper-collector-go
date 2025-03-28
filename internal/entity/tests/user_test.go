package tests

import (
	"github.com/savioafs/web-scraper-collector-go/internal/common"
	"github.com/savioafs/web-scraper-collector-go/internal/entity"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestNewUser(t *testing.T) {
	type userInput struct {
		name     string
		email    string
		password string
	}

	tests := []struct {
		name      string
		user      userInput
		expectErr error
	}{
		{
			name: "Create user with success",
			user: userInput{
				name:     "user one",
				email:    "userone@gmail.com",
				password: "password",
			},
			expectErr: nil,
		},
		{
			name: "Create user with error name is required",
			user: userInput{
				email:    "userone@gmail.com",
				password: "password",
			},
			expectErr: common.ErrNameIsRequired,
		},
		{
			name: "Create user with error email is required",
			user: userInput{
				name:     "user one",
				password: "password",
			},
			expectErr: common.ErrEmailIsRequired,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			as := assert.New(t)

			user, err := entity.NewUser(tt.user.name, tt.user.email, tt.user.password)
			if err != nil {
				as.Equal(err, tt.expectErr)
				return
			}
			as.NotNil(user)
			as.Equal(user.Email, tt.user.email)
			as.Equal(user.Name, tt.user.name)
			as.NotEqual(user.Password, tt.user.password)

			err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(tt.user.password))
			as.NoError(err)

		})
	}
}
