package usecase

import (
	"context"
	"main/domain/model"
	"main/domain/repository"
)

type UserUseCase interface {
	GetUser(context.Context) ([]*model.User, error)
}

type userUseCase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(ur repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: ur,
	}
}

func (ur userUseCase) GetUser(ctx context.Context) (users []*model.User, err error) {
	users, err = ur.userRepository.GetUser(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}
