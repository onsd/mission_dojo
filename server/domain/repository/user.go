package repository

import (
	"context"
	"main/domain/model"
)

type UserRepository interface {
	GetUser(context.Context) ([]*model.User, error)
}
