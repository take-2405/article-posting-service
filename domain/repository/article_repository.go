package repository

type ArticleRepository interface {
	CreateNewArticle(articleID, title, description, content, userID string, tags []string) error
	FixArticle() error
	DeleteArticle() error
	SearchArticles() error
	SendArticle() error
}
