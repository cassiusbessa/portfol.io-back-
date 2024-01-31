package usecases__test

import (
	"testing"

	"github.com/Grupo-38-Orange-Juice/orange-portfolio-back/domain/entities"
	usecases "github.com/Grupo-38-Orange-Juice/orange-portfolio-back/domain/use-cases"
	data_mocks "github.com/Grupo-38-Orange-Juice/orange-portfolio-back/domain/use-cases/tests/mocks"
)

func TestCreateProjectUseCase(t *testing.T) {
	testcases := []struct {
		name        string
		projectRepo usecases.ProjectRepository
		userRepo    usecases.UserRepository
		wantErr     bool
	}{
		{
			name:        "should return nil when project doesn't exist",
			projectRepo: data_mocks.NewMockProjectRepository(false),
			userRepo:    data_mocks.NewMockUserRepository(true),
			wantErr:     false,
		},
		{
			name:        "should return error when there is already a project with the same name from the same user",
			projectRepo: data_mocks.NewMockProjectRepository(true),
			userRepo:    data_mocks.NewMockUserRepository(true),
			wantErr:     true,
		},
		{
			name:        "should return error when user doesn't exist",
			projectRepo: data_mocks.NewMockProjectRepository(false),
			userRepo:    data_mocks.NewMockUserRepository(false),
			wantErr:     true,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			projectUseCase := usecases.NewProjectUseCase(tc.projectRepo, tc.userRepo)
			project := &entities.Project{}
			err := projectUseCase.CreateProject(project, "any_id", []string{"any_id"})
			if tc.wantErr && err == nil {
				t.Errorf("expected %v but got nil", tc.wantErr)
			}
			if !tc.wantErr && err != nil {
				t.Errorf("expected nil but got %v", err)
			}
		})
	}
}
