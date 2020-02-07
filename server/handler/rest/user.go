package rest

import (
	"fmt"
	"main/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

// UserHandler : UserのHandlerのインターフェース
type UserHandler interface {
	GetUser(c echo.Context) error
	CreateUser(c echo.Context) error
}

type userHandler struct {
	userUseCase usecase.UserUseCase
}

func NewUserHandler(ur usecase.UserUseCase) UserHandler {
	return &userHandler{
		userUseCase: ur,
	}
}

func (uh userHandler) GetUser(c echo.Context) error {
	type userField struct {
		Name string `json:"name"`
	}
	type response struct {
		User userField
	}
	token := c.Request().Header.Get("x-token")
	user, err := uh.userUseCase.GetUser(token)
	if err != nil {
		// TODO: エラーハンドリングをきちんとする
		fmt.Errorf("%v", err)
		c.NoContent(http.StatusInternalServerError)
		return err
	}
	res := new(userField)
	res.Name = user.Name
	// クライアントにレスポンスを返却
	return c.JSON(http.StatusOK, res)
}

func (uh userHandler) CreateUser(c echo.Context) error {
	type request struct {
		Name string `json:"name"`
	}
	type response struct {
		Token string `json:"token"`
	}
	req := new(request)
	if err := c.Bind(req); err != nil {
		return err
	}

	user, err := uh.userUseCase.CreateUser(req.Name)
	if err != nil {
		// TODO: エラーハンドリングをきちんとする
		fmt.Errorf("%v", err)
		c.NoContent(http.StatusInternalServerError)
		return err
	}

	res := new(response)
	res.Token = user.Token
	return c.JSON(http.StatusOK, res)
}
