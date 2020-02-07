package persistence

// repositoryにしたいがdomain/respositoryとかぶるのでpersistence
import (
	"database/sql"
	"main/domain/model"
	"main/domain/repository"

	_ "github.com/go-sql-driver/mysql"
)

type userPersistence struct{}

// NewUserPersistence : User dataに関する persistenceを生成
func NewUserPersistence() repository.UserRepository {
	return &userPersistence{}
}

func (up userPersistence) GetUser(token string) (*model.User, error) {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/test")
	defer db.Close()
	if err != nil {
		return nil, err
	}
	user := model.User{}
	err = db.QueryRow("SELECT * FROM users WHERE token =? LIMIT 1", token).Scan(&user.ID, &user.Name, &user.Token)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
