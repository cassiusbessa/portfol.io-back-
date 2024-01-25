package usecases

import "github.com/Grupo-38-Orange-Juice/orange-portfolio-back/domain/entities"

type UserUseCase struct {
	userRepository UserRepository
	crypto         Crypto
}

func NewUserUseCase(userRepository UserRepository, crypto Crypto) UserUseCase {
	return UserUseCase{
		userRepository: userRepository,
		crypto:         crypto,
	}
}

func (u UserUseCase) CreateUser(user *entities.User) error {
	founded, _ := u.userRepository.FindUserByEmail(user.Email)
	if founded.ID != "" {
		possibleValues := map[string]string{user.Email: user.Email}
		return entityAlreadyExists("user", "email", possibleValues)
	}
	hash, err := u.crypto.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hash
	err = u.userRepository.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}
