package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	request2 "prac-orm-transaction/presentation/request"
	"prac-orm-transaction/presentation/response"
	"prac-orm-transaction/usecase"
)

type ArticleHandler interface {
	CreateArticle() http.HandlerFunc
	//FixArticle() http.HandlerFunc
	DeleteArticle() http.HandlerFunc
	//SearchArticles() http.HandlerFunc
	//SendArticles() http.HandlerFunc
}

type articleHandler struct {
	articleUseCase usecase.ArticleUseCase
}

func NewArticleHandler(articleUseCase usecase.ArticleUseCase) *articleHandler {
	return &articleHandler{articleUseCase: articleUseCase}
}

func (ah *articleHandler) CreateArticle() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// リクエストBodyから更新後情報を取得
		var createArticleRequest request2.CreateArticleRequest
		var images []string
		var tags []string

		userID := request.Header.Get("userID")
		json.NewDecoder(request.Body).Decode(&createArticleRequest)
		if createArticleRequest.Content == "" || createArticleRequest.Title == "" || createArticleRequest.Description == "" {
			log.Println("[ERROR] request bucket is err")
			response.RespondError(writer, http.StatusBadRequest, fmt.Errorf("リクエスト情報が不足しています"))
			return
		}

		for _, image := range createArticleRequest.Images {
			images = append(images, image.Image)
		}
		for _, tag := range createArticleRequest.Tags {
			tags = append(tags, tag.Tag)
		}

		articleID, err := ah.articleUseCase.CreateArticle(createArticleRequest.Title, createArticleRequest.Description,
			createArticleRequest.Content, userID, images, tags)
		if err != nil {
			response.RespondError(writer, http.StatusInternalServerError, err)
			return
		}

		writer.Write([]byte(articleID))
	}
}

func (ah *articleHandler) FixArticle() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// リクエストBodyから更新後情報を取得
		var fixArticleRequest request2.FixArticleRequest
		json.NewDecoder(request.Body).Decode(&fixArticleRequest)

		if fixArticleRequest.Id == "" {
			log.Println("[ERROR] request bucket is err")
			response.RespondError(writer, http.StatusBadRequest, fmt.Errorf("リクエスト情報が不足しています"))
			return
		}

		articleID, err := ah.articleUseCase.FixArticle(fixArticleRequest)
		if err != nil {
			response.RespondError(writer, http.StatusInternalServerError, err)
			return
		}

		writer.Write([]byte(articleID))
	}
}

func (ah *articleHandler) DeleteArticle() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// リクエストBodyから更新後情報を取得
		userID := request.Header.Get("userID")
		articleID := request.Header.Get("articleID")

		if articleID == "" {
			response.RespondError(writer, http.StatusInternalServerError, errors.New("request  is error"))
			return
		}

		err := ah.articleUseCase.DeleteArticle(articleID, userID)
		if err != nil {
			response.RespondError(writer, http.StatusInternalServerError, err)
			return
		}

		writer.Write([]byte("delete success"))
	}
}
