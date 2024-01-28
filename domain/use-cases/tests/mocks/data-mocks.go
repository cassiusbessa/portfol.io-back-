package data_mocks

import (
	"fmt"

	"github.com/Grupo-38-Orange-Juice/orange-portfolio-back/domain/aggregate"
	"github.com/Grupo-38-Orange-Juice/orange-portfolio-back/domain/entities"
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

func NewMockProjectRepository(foundedProject bool) *MockProjectRepository {
	return &MockProjectRepository{
		foundedProject: foundedProject,
	}
}

func (m *MockProjectRepository) CreateProject(project *entities.Project, userId string) error {
	return nil
}

func (m *MockProjectRepository) FindAllProjects() ([]aggregate.Project, error) {
	var projects []aggregate.Project
	if m.foundedProject {
		return []aggregate.Project{
			{
				Project: entities.Project{
					ID: "validID",
				},
			},
		}, nil
	}
	return projects, nil
}

func (m *MockProjectRepository) FindProjectsByUserId(userId string) ([]aggregate.Project, error) {
	var projects []aggregate.Project

	if m.foundedProject {
		return []aggregate.Project{
			{
				Project: entities.Project{
					ID: "validID",
				},
			},
		}, nil
	}
	return projects, nil
}

func (m *MockProjectRepository) FindProjectByNameAndUserId(name, userId string) (*aggregate.Project, error) {
	if m.foundedProject {
		return &aggregate.Project{
			Project: entities.Project{
				ID: "validID",
			},
		}, nil
	}
	return nil, nil
}

func (m *MockProjectRepository) UpdateProject(project *entities.Project) (*aggregate.Project, error) {
	if m.foundedProject {
		return &aggregate.Project{
			Project: entities.Project{
				ID: "validID",
			},
		}, nil
	}
	return nil, nil
}

func (m *MockProjectRepository) DeleteProject(projectId string) error {
	if m.foundedProject {
		return nil
	}
	return fmt.Errorf("project not found")
}
