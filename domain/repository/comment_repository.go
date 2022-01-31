package repository

import "prac-orm-transaction/infrastructure/table"

type CommentRepository interface {
	CreateComment(article table.Articles) error
	DeleteArticle() error
	//SearchComments() error
	//SendComments() error
}
