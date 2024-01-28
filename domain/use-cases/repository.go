package usecases

import (
	"github.com/Grupo-38-Orange-Juice/orange-portfolio-back/domain/aggregate"
	"github.com/Grupo-38-Orange-Juice/orange-portfolio-back/domain/entities"
)

type UserRepository interface {
	CreateUser(user *entities.User) error
	FindUserByEmail(email string) (*entities.User, error)
	FindUserById(id string) (*entities.User, error)
}

type ProjectRepository interface {
	CreateProject(project *entities.Project, userId string) error
	FindAllProjects() ([]aggregate.Project, error)
	FindProjectsByUserId(userId string) ([]aggregate.Project, error)
	FindProjectByNameAndUserId(name, userId string) (*aggregate.Project, error)
	UpdateProject(project *entities.Project) (*aggregate.Project, error)
	DeleteProject(projectId string) error
}
