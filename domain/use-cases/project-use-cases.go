package usecases

import "github.com/Grupo-38-Orange-Juice/orange-portfolio-back/domain/entities"

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

func (p ProjectUseCase) CreateProject(project *entities.Project, userId string) error {
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

	err = p.projectRepository.CreateProject(project, userId)
	if err != nil {
		return err
	}
	return nil
}

func (p ProjectUseCase) FindAllProjects() ([]entities.Project, error) {
	projects, err := p.projectRepository.FindAllProjects()
	if err != nil {
		return nil, err
	}
	return projects, nil
}

func (p ProjectUseCase) FindProjectsByUserId(userId string) ([]entities.Project, error) {
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

func (p ProjectUseCase) UpdateProject(project *entities.Project, userId string) error {
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

	_, err = p.projectRepository.UpdateProject(project)
	if err != nil {
		return err
	}
	return nil
}
