package table

//-------------------------------------//
//SQL struct
//-------------------------------------//

type Articles struct {
	ID          string
	Title       string
	Description string
	Nice        int
	Contents    string
	UserID      string
}

type ArticleImage struct {
	ID        string
	ArticleID string
	Image     string
}

type ArticleTag struct {
	ID        string
	ArticleID string
	Tag       string
}

type ArticleComment struct {
	ID       string
	Name     string
	Contents string
}

type ArticleNiceStatusTable struct {
	ID        string
	ArticleID string
	UserID    string
}
