package usecases

import (
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

func (p ProjectUseCase) CreateProject(project *entities.Project, userId string, tagsId []string) error {
	user, err := p.UserRepository.FindUserById(userId)
	if user == nil {
		return entityNotFound("user", "id", map[string]string{userId: userId})
	}
	if err != nil {
		return err
	}

	founded, err := p.projectRepository.FindProjectByNameAndUserId(project.Name, userId)
	if err != nil {
		return err
	}
	if founded != nil {
		possibleValues := map[string]string{project.Name: project.Name}
		return entityAlreadyExists("project", "name", possibleValues)
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
		return nil, entityNotFound("user", "id", map[string]string{userId: userId})
	}

	projects, err := p.projectRepository.FindProjectsByUserId(userId)

	if err != nil {
		return nil, err
	}
	return projects, nil
}

func (p ProjectUseCase) UpdateProject(project *entities.Project, userId string, tagsId []string) error {
	user, err := p.UserRepository.FindUserById(userId)
	if user == nil {
		return entityNotFound("user", "id", map[string]string{userId: userId})
	}
	if err != nil {
		return err
	}

	founded, err := p.projectRepository.FindProjectByNameAndUserId(project.Name, userId)
	if err != nil {
		return err
	}
	if founded != nil && founded.ID != project.ID {
		possibleValues := map[string]string{project.Name: project.Name}
		return entityAlreadyExists("project", "name", possibleValues)
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
		return entityNotFound("project", "id and userId", map[string]string{projectId: projectId, userId: userId})
	}
	return nil
}
