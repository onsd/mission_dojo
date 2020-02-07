package main

import (
	"main/handler/rest"
	"main/infra/persistence"
	"main/usecase"
	"net/http"
)

func main() {
	userPersistence := persistence.NewUserPersistence()
	userUseCase := usecase.NewUserUseCase(userPersistence)
	userHander := rest.NewUserHandler(userUseCase)
	// http.HandleFunc("/", handler.IndexHandler)

	// //user
	// http.HandleFunc("/user/create", handler.CreateUser)
	http.HandleFunc("/user/get", userHander.GetUser)
	// http.HandleFunc("/user/update", handler.UpdateUser)

	// 8080ポートで起動
	http.ListenAndServe(":8080", nil)
}
