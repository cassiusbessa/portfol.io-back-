package aggregates

import "github.com/Grupo-38-Orange-Juice/orange-portfolio-back/domain/entities"

type Project struct {
	entities.Project
	entities.User
	tags []string
}

func NewProject(project entities.Project, user entities.User, tags []string) Project {
	return Project{
		Project: project,
		User:    user,
		tags:    tags,
	}
}
