package usecase

import (
	"github.com/google/uuid"
	"log"
	"prac-orm-transaction/domain/repository"
	"prac-orm-transaction/presentation/request"
)

type ArticleUseCase interface {
	CreateArticle(title, description, content, userID string, images, tags []string) (string, error)
	FixArticle(request.FixArticleRequest) (string, error)
	DeleteArticle(articleID, userID string) error
	SearchArticle(id, pass string) (string, error)
	SearchArticles(id, pass string) (string, error)
	SendArticles(id, pass string) (string, error)
}

type articleUseCase struct {
	article repository.ArticleRepository
	tag     repository.TagRepository
	image   repository.ImageRepository
}

func NewArticleUseCase(article repository.ArticleRepository) *articleUseCase {
	return &articleUseCase{article: article}
}

func (au articleUseCase) CreateArticle(title, description, content, userID string, images, tags []string) (string, error) {
	var articleID string
	uuid, err := uuid.NewRandom()
	if err != nil {
		log.Println(err)
		return articleID, err
	}

	articleID = uuid.String()
	if err = au.article.CreateNewArticle(articleID, title, description, content, userID, tags); err != nil {
		log.Println(err)
		return articleID, err
	}

	return articleID, err
}

func (au articleUseCase) FixArticle(request.FixArticleRequest) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (au articleUseCase) DeleteArticle(articleID, userID string) error {
	if err := au.article.DeleteArticle(articleID, userID); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (au articleUseCase) SearchArticle(id, pass string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (au articleUseCase) SearchArticles(id, pass string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (au articleUseCase) SendArticles(id, pass string) (string, error) {
	//TODO implement me
	panic("implement me")
}
