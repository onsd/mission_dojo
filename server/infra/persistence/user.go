package persistence

// repositoryにしたいがdomain/respositoryとかぶるのでpersistence
import (
	"database/sql"
	"fmt"
	"main/domain/model"
	"main/domain/repository"
	"os"

	_ "github.com/go-sql-driver/mysql" // for mysql
)

type userPersistence struct{}

// NewUserPersistence : User dataに関する persistenceを生成
func NewUserPersistence() repository.UserRepository {
	return &userPersistence{}
}
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
func (up userPersistence) GetUser(token string) (*model.User, error) {
	connectionInformation := getEnv("MYSQL_USER", "root") + ":" + getEnv("MYSQL_PASSOWRD", "password") + "@tcp(" + getEnv("MYSQL_HOST", "localhost") + ":" + getEnv("MYSQL_PORT", "3306") + ")/" + getEnv("MYSQL_DATABASE", "database")
	fmt.Println(connectionInformation)
	db, err := sql.Open("mysql", connectionInformation) //FIXME 切り出す？
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

func (up userPersistence) CreateUser(name, token string) (*model.User, error) {
	connectionInformation := getEnv("MYSQL_USER", "root") + ":" + getEnv("MYSQL_PASSOWRD", "password") + "@tcp(" + getEnv("MYSQL_HOST", "localhost") + ":" + getEnv("MYSQL_PORT", "3306") + ")/" + getEnv("MYSQL_DATABASE", "database")
	db, err := sql.Open("mysql", connectionInformation) //FIXME 切り出す？
	defer db.Close()
	if err != nil {
		return nil, err
	}

	stmt, err := db.Prepare("INSERT INTO users (name, token) values (?, ?)")
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(name, token)
	if err != nil {
		return nil, err
	}

	user := model.User{Token: token}
	return &user, nil
}
