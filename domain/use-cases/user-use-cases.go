package usecases

import "github.com/Grupo-38-Orange-Juice/orange-portfolio-back/domain/entities"

type UserUseCase struct {
	userRepository UserRepository
}

func NewUserUseCase(userRepository UserRepository) UserUseCase {
	return UserUseCase{
		userRepository: userRepository,
	}
}

func (u UserUseCase) CreateUser(user *entities.User) error {
	founded, _ := u.userRepository.FindUserByEmail(user.Email)
	if founded != nil {
		return entityAlreadyExists("usuário", "e-mail", user.Email)
	}

	err := u.userRepository.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (u UserUseCase) FindUserByEmail(email string) (*entities.User, error) {
	user, err := u.userRepository.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u UserUseCase) FindUserById(id string) (*entities.User, error) {
	user, err := u.userRepository.FindUserById(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, entityNotFound("usuário", "id", id)
	}

	return user, nil
}
