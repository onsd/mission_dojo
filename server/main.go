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
	userHandler := rest.NewUserHandler(userUseCase)

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes for user
	userGroup := e.Group("/user")
	userGroup.GET("/get", userHandler.GetUser)
	userGroup.POST("/create", userHandler.CreateUser)
	// //user
	// http.HandleFunc("/user/create", handler.CreateUser)
	// http.HandleFunc("/user/get", userHander.GetUser)
	// http.HandleFunc("/user/update", handler.UpdateUser)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
