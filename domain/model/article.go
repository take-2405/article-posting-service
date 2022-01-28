package model

type Article struct {
	ID          string
	Title       string
	Description string
	Tags        []string
	Nice        int
	Comments    []comment
	Images      []string
}

type comment struct {
	ID       string
	Contents string
	Name     string
}
