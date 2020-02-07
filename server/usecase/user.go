package usecase

import (
	"main/domain/model"
	"main/domain/repository"
)

type UserUseCase interface {
	GetUser(token string) (*model.User, error)
	CreateUser(name string) (*model.User, error)
}

type userUseCase struct {
	userRepository repository.UserRepository
}

// NewUserUseCase: UserのためのUseCase
func NewUserUseCase(ur repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: ur,
	}
}

func (ur userUseCase) GetUser(token string) (users *model.User, err error) {
	users, err = ur.userRepository.GetUser(token)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (ur userUseCase) CreateUser(name string) (user *model.User, err error) {
	user, err = ur.userRepository.CreateUser(name)
	if err != nil {
		return nil, err
	}
	return user, nil
}
