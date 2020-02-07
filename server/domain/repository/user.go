package repository

import (
	"main/domain/model"
)

// UserRepository : userのrepository
type UserRepository interface {
	GetUser(token string) (*model.User, error)
	CreateUser(token string) (*model.User, error)
}
