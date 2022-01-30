package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	request2 "prac-orm-transaction/interface/request"
	"prac-orm-transaction/interface/response"
	"prac-orm-transaction/usecase"
)

type UserHandler interface {
	CreateUserAccount() http.HandlerFunc
	SignIn() http.HandlerFunc
}

type userHandler struct {
	userUseCase usecase.UserUseCase
}

func NewUserHandler(userUseCase usecase.UserUseCase) *userHandler {
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
			return
		}

		token, err := uh.userUseCase.CreateUserAccount(accountInfo.ID, accountInfo.Pass)
		if err != nil {
			response.RespondError(writer, http.StatusInternalServerError, err)
			return
		}

		writer.Write([]byte(token))
	}
}

func (uh *userHandler) SignIn() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// リクエストBodyから更新後情報を取得
		var accountInfo request2.CreateUserAccountRequest
		json.NewDecoder(request.Body).Decode(&accountInfo)

		if accountInfo.ID == "" || accountInfo.Pass == "" {
			log.Println("[ERROR] request bucket is err")
			response.RespondError(writer, http.StatusBadRequest, fmt.Errorf("リクエスト情報が不足しています"))
			return
		}

		token, err := uh.userUseCase.SignIn(accountInfo.ID, accountInfo.Pass)
		if err != nil {
			response.RespondError(writer, http.StatusInternalServerError, err)
			return
		}

		writer.Write([]byte(token))
	}
}
