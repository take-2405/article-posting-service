package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	request2 "prac-orm-transaction/api/request"
	"prac-orm-transaction/api/response"
	"prac-orm-transaction/application"
)

type UserHandler interface {
	CreateUserAccount() http.HandlerFunc
}

type userHandler struct {
	userUseCase application.UserUseCase
}

func NewUserHandler(userUseCase application.UserUseCase) *userHandler {
	return &userHandler{userUseCase: userUseCase}
}

func (uh *userHandler) CreateUserAccount() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// リクエストBodyから更新後情報を取得
		var accountInfo request2.CreateUserAccountRequest
		json.NewDecoder(request.Body).Decode(&accountInfo)

		if accountInfo.ID == "" || accountInfo.Pass == "" {
			log.Println("[ERROR] request bucket is err")
			response.RespondError(writer, http.StatusBadRequest, fmt.Errorf("リクエスト情報が不足しています"))
		}

		token, err := uh.userUseCase.CreateUserAccount(accountInfo.ID, accountInfo.Pass)
		if err != nil {
			response.RespondError(writer, http.StatusInternalServerError, err)
		}

		if token == "" {
			response.RespondError(writer, http.StatusInternalServerError, fmt.Errorf("tokenが空です"))
		}

		writer.Write([]byte("hello new user!"))
	}
}