package request

type CreateArticleRequest struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Content     string  `json:"content"`
	Images      []image `json:"images"`
	Tags        []tag   `json:"tags"`
}

type FixArticleRequest struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Content     string  `json:"content"`
	Images      []image `json:"images"`
	Tags        []tag   `json:"tags"`
}

type image struct {
	Image string `json:"image"`
}

type tag struct {
	Tag string `json:"tag"`
}
