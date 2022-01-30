package repository

type ArticleRepository interface {
	CreateNewArticle(articleID, title, description, content, userID string, tags []string) error
	FixArticle() error
	DeleteArticle(articleID, userID string) error
	SearchArticles() error
	SendArticle() error
}
