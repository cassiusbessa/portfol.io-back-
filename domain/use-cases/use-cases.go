package usecases

import (
	"github.com/Grupo-38-Orange-Juice/orange-portfolio-back/domain/aggregates"
	"github.com/Grupo-38-Orange-Juice/orange-portfolio-back/domain/entities"
)

type UserUseCases interface {
	CreateUser(user *entities.User) error
	FindUserByEmail(email string) (*entities.User, error)
	FindUserById(id string) (*entities.User, error)
}

type ProjectUseCases interface {
	CreateProject(project *entities.Project, userId string, tagsId []int) error
	FindAllProjects() ([]aggregates.Project, error)
	FindProjectsByUserId(userId string) ([]aggregates.Project, error)
	UpdateProject(project *entities.Project, userId string, tagsId []int) error
	DeleteProject(projectId, userId string) error
}

type TagUseCases interface {
	FindAllTags() ([]entities.Tag, error)
}
