package entities_test

import (
	"testing"

	"github.com/Grupo-38-Orange-Juice/orange-portfolio-back/domain/entities"
)

func TestNewProject(t *testing.T) {
	project := entities.Project{
		Name:        "Valid Name",
		Description: "Valid Description",
	}
	testCases := []struct {
		description string
		project     entities.Project
		wantErr     bool
	}{
		{
			description: "should return nil when project is valid",
			project:     project,
			wantErr:     false,
		},
		{
			description: "should return error when project name is empty",
			project: entities.Project{
				Name:        "",
				Description: project.Description,
			},
			wantErr: true,
		},
		{
			description: "should return error when project name is too short",
			project: entities.Project{
				Name:        "a",
				Description: project.Description,
			},
			wantErr: true,
		},
		{
			description: "should return error when project description is empty",
			project: entities.Project{
				Name:        project.Name,
				Description: "",
			},
			wantErr: true,
		},
		{
			description: "should return error when project description is too short",
			project: entities.Project{
				Name:        project.Name,
				Description: "a",
			},
			wantErr: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.description, func(t *testing.T) {
			_, err := entities.NewProject(tC.project)
			if tC.wantErr && err == nil {
				t.Errorf("NewProject() error = %v, wantErr %v", err, tC.wantErr)
			}
			if !tC.wantErr && err != nil {
				t.Errorf("NewProject() error = %v, wantErr %v", err, tC.wantErr)
			}
		})
	}
}
