package usecases

import "github.com/Grupo-38-Orange-Juice/orange-portfolio-back/domain/entities"

type UserUseCases interface {
	CreateUser(user *entities.User) error
	FindUserByEmail(email string) (*entities.User, error)
}
