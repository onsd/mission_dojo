package repository

import (
	"main/domain/model"
)

type UserRepository interface {
	GetUser(token string) (*model.User, error)
}
