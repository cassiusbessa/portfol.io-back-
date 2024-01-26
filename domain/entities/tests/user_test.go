package entities_test

import (
	"testing"

	"github.com/Grupo-38-Orange-Juice/orange-portfolio-back/domain/entities"
)

func TestNewUser(t *testing.T) {
	user := entities.User{
		FullName: "Valid Name",
		Email:    "valid@email.com",
		Password: "Valid@Password1",
	}
	testCases := []struct {
		description string
		user        entities.User
		wantErr     bool
	}{
		{
			description: "should return nil when user is valid",
			user:        user,
			wantErr:     false,
		},
		{
			description: "should return error when user full name is empty",
			user: entities.User{
				FullName: "",
				Email:    user.Email,
				Password: user.Password,
			},
			wantErr: true,
		},
		{
			description: "should return error when user full name is too short",
			user: entities.User{
				FullName: "a",
				Email:    user.Email,
				Password: user.Password,
			},
			wantErr: true,
		},
		{
			description: "should return error when user email is empty",
			user: entities.User{
				FullName: user.FullName,
				Email:    "",
				Password: user.Password,
			},
			wantErr: true,
		},
		{
			description: "should return error when user email is invalid",
			user: entities.User{
				FullName: user.FullName,
				Email:    "invalidemail@email.com",
				Password: user.Password,
			},
			wantErr: true,
		},
		{
			description: "should return error when user password is empty",
			user: entities.User{
				FullName: user.FullName,
				Email:    user.Email,
				Password: "",
			},
			wantErr: true,
		},
		{
			description: "should return error when user password is too short",
			user: entities.User{
				FullName: user.FullName,
				Email:    user.Email,
				Password: "a",
			},
			wantErr: true,
		},
		{
			description: "should return error when user password is invalid",
			user: entities.User{
				FullName: user.FullName,
				Email:    user.Email,
				Password: "invalidpassword",
			},
			wantErr: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			_, err := entities.NewUser(testCase.user)
			if err != nil && !testCase.wantErr {
				t.Errorf("NewUser() error = %v, wantErr %v", err, testCase.wantErr)
			}
			if err == nil && testCase.wantErr {
				t.Errorf("NewUser() error = %v, wantErr %v", err, testCase.wantErr)
			}
		})
	}
}
