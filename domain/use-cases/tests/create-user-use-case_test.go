package usecases__test

import (
	"testing"

	"github.com/Grupo-38-Orange-Juice/orange-portfolio-back/domain/entities"
	usecases "github.com/Grupo-38-Orange-Juice/orange-portfolio-back/domain/use-cases"
	data_mocks "github.com/Grupo-38-Orange-Juice/orange-portfolio-back/domain/use-cases/tests/mocks"
)

func TestCreateUserUseCase(t *testing.T) {
	testcases := []struct {
		name     string
		userRepo usecases.UserRepository
		crypto   usecases.Crypto
		wantErr  bool
	}{
		{
			name:     "should return nil when email doesn't exist",
			userRepo: data_mocks.NewMockUserRepository(false),
			crypto:   data_mocks.NewMockCrypto(true),
			wantErr:  false,
		},
		{
			name:     "should return error when email already exists",
			userRepo: data_mocks.NewMockUserRepository(true),
			crypto:   data_mocks.NewMockCrypto(true),
			wantErr:  true,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			userUseCase := usecases.NewUserUseCase(tc.userRepo, tc.crypto)
			user := &entities.User{}
			err := userUseCase.CreateUser(user)
			if tc.wantErr && err == nil {
				t.Errorf("expected error but got nil")
			}
			if !tc.wantErr && err != nil {
				t.Errorf("expected nil but got error")
			}
		})
	}
}
