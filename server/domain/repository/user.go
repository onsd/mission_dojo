package repository

import (
	"main/domain/model"
)

// UserRepository : userのrepository
type UserRepository interface {
	GetUser(token string) (*model.User, error)
	CreateUser(name, token string) (*model.User, error)
	UpdateUser(name, token string) error
}
