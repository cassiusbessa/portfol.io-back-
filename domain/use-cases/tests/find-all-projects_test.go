package usecases__test

import (
	"testing"

	usecases "github.com/Grupo-38-Orange-Juice/orange-portfolio-back/domain/use-cases"
	data_mocks "github.com/Grupo-38-Orange-Juice/orange-portfolio-back/domain/use-cases/tests/mocks"
)

func TestFindAllProjectsUseCase(t *testing.T) {
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
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			projectUseCase := usecases.NewProjectUseCase(tc.projectRepo, tc.userRepo)
			projects, err := projectUseCase.FindAllProjects()
			if tc.wantErr && err == nil {
				t.Errorf("expected %v but got nil", tc.wantErr)
			}
			if !tc.wantErr && err != nil {
				t.Errorf("expected nil but got %v", err)
			}
			if len(projects) == 0 && tc.wantErr {
				println(projects)
				t.Errorf("expected a slice of projects but got an error")
			}
		})
	}
}
