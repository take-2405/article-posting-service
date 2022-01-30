package repository

type ArticleRepository interface {
	CreateNewArticle(articleID, title, description, content, userID string, tags, images []string) error
	FixArticle(articleID, title, description, content, userID string, tags, images []string) error
	DeleteArticle(articleID, userID string) error
	SearchArticles() error
	SendArticle() error
}
