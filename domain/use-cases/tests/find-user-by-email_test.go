package usecases__test

import (
	"testing"

	"github.com/Grupo-38-Orange-Juice/orange-portfolio-back/domain/entities"
	usecases "github.com/Grupo-38-Orange-Juice/orange-portfolio-back/domain/use-cases"
	data_mocks "github.com/Grupo-38-Orange-Juice/orange-portfolio-back/domain/use-cases/tests/mocks"
)

func TestFindUserByEmailUseCase(t *testing.T) {
	testcases := []struct {
		name     string
		userRepo usecases.UserRepository
		wantErr  bool
		expected *entities.User
	}{
		{
			name:     "should return error when email doesn't exist",
			userRepo: data_mocks.NewMockUserRepository(false),
			wantErr:  true,
			expected: nil,
		},
		{
			name:     "should return a user when repository returns a user",
			userRepo: data_mocks.NewMockUserRepository(true),
			wantErr:  false,
			expected: &entities.User{
				ID: "validID",
			},
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			userUseCase := usecases.NewUserUseCase(tc.userRepo)
			_, err := userUseCase.FindUserByEmail("any_email@email.com")
			if tc.wantErr && err == nil {
				t.Errorf("expected %v but got nil", tc.wantErr)
			}
			if !tc.wantErr && err != nil {
				t.Errorf("expected nil but got %v", err)
			}
		})
	}
}
