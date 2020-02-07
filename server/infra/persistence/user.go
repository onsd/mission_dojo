package persistence

// repositoryにしたいがdomain/respositoryとかぶるのでpersistence
import (
	"context"
	"main/domain/model"
	"main/domain/repository"
)

type userPersistence struct{}

// NewUserPersistence : User dataに関する persistenceを生成
func NewUserPersistence() repository.UserRepository {
	return &userPersistence{}
}

func (up userPersistence) GetUser(context.Context) (*model.User, error) {
	user1 := model.User{}
	user1.ID = "1"
	user1.Token = "AAABBBCCCDDD"
	user1.Name = "Taka"
	return &user1, nil
}
