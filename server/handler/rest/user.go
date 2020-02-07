package rest

import (
	"encoding/json"
	"main/usecase"
	"net/http"
)

// UserHandler : UserのHandlerのインターフェース
type UserHandler interface {
	GetUser(http.ResponseWriter, *http.Request)
}

type userHandler struct {
	userUseCase usecase.UserUseCase
}

func NewUserHandler(ur usecase.UserUseCase) UserHandler {
	return &userHandler{
		userUseCase: ur,
	}
}

func (uh userHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	type userField struct {
		Name string `json:"name"`
	}
	type response struct {
		User userField
	}
	user, err := uh.userUseCase.GetUser(ctx)
	if err != nil {
		// TODO: エラーハンドリングをきちんとする
		http.Error(w, "Internal Server Error", 500)
		return
	}
	res := new(userField)
	res.Name = user.Name
	// クライアントにレスポンスを返却
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(res); err != nil {
		// TODO: エラーハンドリングをきちんとする
		http.Error(w, "Internal Server Error", 500)
		return
	}
}
