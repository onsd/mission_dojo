package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"main/handler"
	"net/http"
)

func main() {
	checkSQLConnection()
	http.HandleFunc("/", handler.IndexHandler)

	//user
	http.HandleFunc("/user/create", handler.CreateUser)
	http.HandleFunc("/user/get", handler.GetUser)
	http.HandleFunc("/user/update", handler.UpdateUser)

	// 8080ポートで起動
	http.ListenAndServe(":8080", nil)
}

func checkSQLConnection() {
	db, err := sql.Open("mysql", "test:password@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var id int
	var name string
	rows, err := db.Query("select * from users where id < ?", 6)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println(id, name)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}
}
