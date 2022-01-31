package repository

type TagRepository interface {
	CreateTag(articleID string, tag []string) error
	DeleteTag() error
	//SearchTags() error
	//SendTags() error
}
