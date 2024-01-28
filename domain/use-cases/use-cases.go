package usecases

import "github.com/Grupo-38-Orange-Juice/orange-portfolio-back/domain/entities"

type UserUseCases interface {
	CreateUser(user *entities.User) error
	FindUserByEmail(email string) (*entities.User, error)
}

type ProjectUseCases interface {
	CreateProject(project *entities.Project) error
	FindAllProjects() ([]entities.Project, error)
	FindProjectsByUserId(userId string) ([]entities.Project, error)
	UpdateProject(project *entities.Project) error
	DeleteProject(projectId string) error
}
