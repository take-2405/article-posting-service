package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	request2 "prac-orm-transaction/presentation/request"
	"prac-orm-transaction/presentation/response"
	"prac-orm-transaction/usecase"
)

type AuthHandler interface {
	SignUp() http.HandlerFunc
	SignIn() http.HandlerFunc
}

type authHandler struct {
	authUseCase usecase.AuthUseCase
}

func NewUserHandler(authUseCase usecase.AuthUseCase) *authHandler {
	return &authHandler{authUseCase: authUseCase}
}

func (uh *authHandler) SignUp() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// リクエストBodyから更新後情報を取得
		var accountInfo request2.CreateUserAccountRequest
		json.NewDecoder(request.Body).Decode(&accountInfo)

		if accountInfo.ID == "" || accountInfo.Pass == "" {
			log.Println("[ERROR] request bucket is err")
			response.RespondError(writer, http.StatusBadRequest, fmt.Errorf("リクエスト情報が不足しています"))
			return
		}

		token, err := uh.authUseCase.SignUp(accountInfo.ID, accountInfo.Pass)
		if err != nil {
			response.RespondError(writer, http.StatusInternalServerError, err)
			return
		}

		writer.Write([]byte(token))
	}
}

func (uh *authHandler) SignIn() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// リクエストBodyから更新後情報を取得
		var accountInfo request2.CreateUserAccountRequest
		json.NewDecoder(request.Body).Decode(&accountInfo)

		if accountInfo.ID == "" || accountInfo.Pass == "" {
			log.Println("[ERROR] request bucket is err")
			response.RespondError(writer, http.StatusBadRequest, fmt.Errorf("リクエスト情報が不足しています"))
			return
		}

		token, err := uh.authUseCase.SignIn(accountInfo.ID, accountInfo.Pass)
		if err != nil {
			response.RespondError(writer, http.StatusInternalServerError, err)
			return
		}

		writer.Write([]byte(token))
	}
}
