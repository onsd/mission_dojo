package main

import (
	"main/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler.IndexHandler)

	//user
	http.HandleFunc("/user/create", handler.CreateUser)
	http.HandleFunc("/user/get", handler.GetUser)
	http.HandleFunc("/user/update", handler.UpdateUser)

	// 8080ポートで起動
	http.ListenAndServe(":8080", nil)
}
