package usecases

import "github.com/Grupo-38-Orange-Juice/orange-portfolio-back/domain/entities"

type UserRepository interface {
	CreateUser(user *entities.User) error
	FindUserByEmail(email string) (*entities.User, error)
}
