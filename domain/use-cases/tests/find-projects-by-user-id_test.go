package usecases__test

import (
	"testing"

	usecases "github.com/Grupo-38-Orange-Juice/orange-portfolio-back/domain/use-cases"
	data_mocks "github.com/Grupo-38-Orange-Juice/orange-portfolio-back/domain/use-cases/tests/mocks"
)

func TestFindProjectsByUserIdUseCase(t *testing.T) {
	testcases := []struct {
		name        string
		projectRepo usecases.ProjectRepository
		userRepo    usecases.UserRepository
		wantErr     bool
	}{
		{
			name:        "should return a slice of projects",
			projectRepo: data_mocks.NewMockProjectRepository(true),
			userRepo:    data_mocks.NewMockUserRepository(true),
			wantErr:     false,
		},
		{
			name:        "should return an empty slice of projects if there are no projects",
			projectRepo: data_mocks.NewMockProjectRepository(false),
			userRepo:    data_mocks.NewMockUserRepository(true),
			wantErr:     false,
		},
		{
			name:        "should return an error if user doesn't exist",
			projectRepo: data_mocks.NewMockProjectRepository(true),
			userRepo:    data_mocks.NewMockUserRepository(false),
			wantErr:     true,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			projectUseCase := usecases.NewProjectUseCase(tc.projectRepo, tc.userRepo)
			_, err := projectUseCase.FindProjectsByUserId("any_id")
			if tc.wantErr && err == nil {
				t.Errorf("expected %v but got nil", tc.wantErr)
			}
			if !tc.wantErr && err != nil {
				t.Errorf("expected nil but got %v", err)
			}
		})
	}
}
