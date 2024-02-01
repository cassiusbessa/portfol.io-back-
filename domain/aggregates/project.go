package aggregates

import "github.com/Grupo-38-Orange-Juice/orange-portfolio-back/domain/entities"

type Project struct {
	Project entities.Project
	User    entities.User
	Tags    []int
}

func NewProject(project entities.Project, user entities.User, tags []int) Project {
	return Project{
		Project: project,
		User:    user,
		Tags:    tags,
	}
}
