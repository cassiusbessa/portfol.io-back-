package http

import "github.com/Grupo-38-Orange-Juice/orange-portfolio-back/domain/aggregates"

func projectAggregateToDTO(project aggregates.Project) ProjectDTO {
	return ProjectDTO{
		Project: ProjectInfo{
			ID:          project.Project.ID,
			Name:        project.Project.Name,
			Description: project.Project.Description,
			Link:        project.Project.Link,
			Image:       project.Project.Image,
			CreatedAt:   project.Project.CreatedAt,
			UpdatedAt:   project.Project.UpdatedAt,
		},
		User: UserInfo{
			ID:       project.User.ID,
			FullName: project.User.FullName,
			Email:    project.User.Email,
			Image:    project.User.Image,
		},
		Tags: project.Tags,
	}
}
