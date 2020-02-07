package main

import (
	"main/handler/rest"
	"main/infra/persistence"
	"main/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	userPersistence := persistence.NewUserPersistence()
	userUseCase := usecase.NewUserUseCase(userPersistence)
	userHander := rest.NewUserHandler(userUseCase)

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/user/get", userHander.GetUser)
	// //user
	// http.HandleFunc("/user/create", handler.CreateUser)
	// http.HandleFunc("/user/get", userHander.GetUser)
	// http.HandleFunc("/user/update", handler.UpdateUser)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
