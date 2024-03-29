package data_mocks

import (
	"fmt"

	"github.com/Grupo-38-Orange-Juice/orange-portfolio-back/domain/aggregates"
	"github.com/Grupo-38-Orange-Juice/orange-portfolio-back/domain/entities"
	usecases "github.com/Grupo-38-Orange-Juice/orange-portfolio-back/domain/use-cases"
)

type MockUserRepository struct {
	foundedUser bool
}

func NewMockUserRepository(foundedUser bool) *MockUserRepository {
	return &MockUserRepository{
		foundedUser: foundedUser,
	}
}
func (m *MockUserRepository) CreateUser(user *entities.User) error {
	return nil
}

func (m *MockUserRepository) FindUserByEmail(email string) (*entities.User, error) {
	if m.foundedUser {
		return &entities.User{
			ID: "validID",
		}, nil
	}
	return nil, fmt.Errorf("user not found")
}

func (m *MockUserRepository) FindUserById(id string) (*entities.User, error) {
	if m.foundedUser {
		return &entities.User{
			ID: "validID",
		}, nil
	}
	return nil, fmt.Errorf("user not found")
}

type MockCrypto struct {
	matchedHashes bool
}

func NewMockCrypto(matchedHashes bool) *MockCrypto {
	return &MockCrypto{
		matchedHashes: matchedHashes,
	}
}

func (m *MockCrypto) HashPassword(password string) (string, error) {
	return "hashedPassword", nil
}

func (m *MockCrypto) CompareHashAndPassword(hash string, password string) error {
	if m.matchedHashes {
		return nil
	}
	return fmt.Errorf("passwords don't match")
}

type MockProjectRepository struct {
	foundedProject bool
}

func NewMockProjectRepository(foundedProject bool) usecases.ProjectRepository {
	return &MockProjectRepository{
		foundedProject: foundedProject,
	}
}

func (m *MockProjectRepository) CreateProject(project *entities.Project, userId string, tagsId []int) error {
	return nil
}

func (m *MockProjectRepository) FindAllProjects() ([]aggregates.Project, error) {
	var projects []aggregates.Project
	if m.foundedProject {
		return []aggregates.Project{
			{
				Project: entities.Project{
					ID: "validID",
				},
			},
		}, nil
	}
	return projects, nil
}

func (m *MockProjectRepository) FindProjectsByUserId(userId string) ([]aggregates.Project, error) {
	var projects []aggregates.Project

	if m.foundedProject {
		return []aggregates.Project{
			{
				Project: entities.Project{
					ID: "validID",
				},
			},
		}, nil
	}
	return projects, nil
}

func (m *MockProjectRepository) FindProjectByNameAndUserId(name, userId string) (*entities.Project, error) {
	if m.foundedProject {
		return &entities.Project{
			ID: "validID",
		}, nil
	}
	return nil, nil
}

func (m *MockProjectRepository) UpdateProject(project *entities.Project, tagsId []int) (*aggregates.Project, error) {
	if m.foundedProject {
		return &aggregates.Project{
			Project: entities.Project{
				ID: "validID",
			},
		}, nil
	}
	return nil, nil
}

func (m *MockProjectRepository) DeleteProject(projectId, userId string) (bool, error) {
	if m.foundedProject {
		return true, nil
	}
	return false, fmt.Errorf("project not found")
}
