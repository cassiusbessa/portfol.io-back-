package data_mocks

import (
	"fmt"

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
	return &entities.User{}, nil
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
