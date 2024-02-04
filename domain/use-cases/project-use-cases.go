package usecases

import (
	"fmt"

	"github.com/Grupo-38-Orange-Juice/orange-portfolio-back/domain/aggregates"
	"github.com/Grupo-38-Orange-Juice/orange-portfolio-back/domain/entities"
)

type ProjectUseCase struct {
	projectRepository ProjectRepository
	UserRepository    UserRepository
}

func NewProjectUseCase(projectRepository ProjectRepository, UserRepository UserRepository) ProjectUseCase {
	return ProjectUseCase{
		projectRepository: projectRepository,
		UserRepository:    UserRepository,
	}
}

func (p ProjectUseCase) CreateProject(project *entities.Project, userId string, tagsId []int) error {
	fmt.Println(tagsId)
	if (tagsId == nil) || (len(tagsId) > 2 || len(tagsId) == 0) {
		return invalidLength("Tags", 1, 2)
	}

	user, err := p.UserRepository.FindUserById(userId)
	if user == nil {
		return entityNotFound("Usuário", "id", userId)
	}
	if err != nil {
		return err
	}

	founded, err := p.projectRepository.FindProjectByNameAndUserId(project.Name, userId)
	if err != nil {
		return err
	}
	if founded != nil {
		return entityAlreadyExists("Projeto", "Nome", project.Name)
	}

	err = p.projectRepository.CreateProject(project, userId, tagsId)
	if err != nil {
		return err
	}
	return nil
}

func (p ProjectUseCase) FindAllProjects() ([]aggregates.Project, error) {
	projects, err := p.projectRepository.FindAllProjects()
	if err != nil {
		return nil, err
	}
	return projects, nil
}

func (p ProjectUseCase) FindProjectsByUserId(userId string) ([]aggregates.Project, error) {
	usr, err := p.UserRepository.FindUserById(userId)
	if err != nil {
		return nil, err
	}
	if usr == nil {
		return nil, entityNotFound("Usuário", "id", userId)
	}

	projects, err := p.projectRepository.FindProjectsByUserId(userId)

	if err != nil {
		return nil, err
	}
	return projects, nil
}

func (p ProjectUseCase) UpdateProject(project *entities.Project, userId string, tagsId []int) error {

	if (tagsId == nil) || (len(tagsId) == 0 || len(tagsId) < 2) {
		return invalidLength("Tags", 1, 2)
	}

	user, err := p.UserRepository.FindUserById(userId)
	if user == nil {
		return entityNotFound("Usuário", "id", userId)
	}
	if err != nil {
		return err
	}

	founded, err := p.projectRepository.FindProjectByNameAndUserId(project.Name, userId)
	if err != nil {
		return err
	}
	if founded != nil && founded.ID != project.ID {
		return entityAlreadyExists("Projeto", "Nome", project.Name)
	}

	_, err = p.projectRepository.UpdateProject(project, tagsId)
	if err != nil {
		return err
	}
	return nil
}

func (p ProjectUseCase) DeleteProject(projectId, userId string) error {
	deleted, err := p.projectRepository.DeleteProject(projectId, userId)
	if err != nil {
		return err
	}
	if !deleted {
		return entityNotFound("Projeto", "id", projectId)
	}
	return nil
}
