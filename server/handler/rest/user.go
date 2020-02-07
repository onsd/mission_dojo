package rest

import (
	"main/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

// UserHandler : UserのHandlerのインターフェース
type UserHandler interface {
	GetUser(c echo.Context) error
}

type userHandler struct {
	userUseCase usecase.UserUseCase
}

func NewUserHandler(ur usecase.UserUseCase) UserHandler {
	return &userHandler{
		userUseCase: ur,
	}
}

// func (uh userHandler) GetUser(c echo.Context) error {
// 	return c.String(http.StatusOK, "Hello, World!")
// }

func (uh userHandler) GetUser(c echo.Context) error {
	type userField struct {
		Name string `json:"name"`
	}
	type response struct {
		User userField
	}
	user, err := uh.userUseCase.GetUser("")
	if err != nil {
		// TODO: エラーハンドリングをきちんとする
		c.NoContent(http.StatusInternalServerError)
		return err
	}
	res := new(userField)
	res.Name = user.Name
	// クライアントにレスポンスを返却
	return c.JSON(http.StatusOK, res)
}
