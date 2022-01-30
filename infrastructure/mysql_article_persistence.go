package infrastructure

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"strings"

	//"errors"
	"log"
	"prac-orm-transaction/domain/repository"
	"prac-orm-transaction/infrastructure/table"
)

type articlePersistence struct {
	mysql mysqlRepository
}

func NewArticlePersistence(mysqlConn mysqlRepository) repository.ArticleRepository {
	return &articlePersistence{mysql: mysqlConn}
}

func (a *articlePersistence) CreateNewArticle(articleID, title, description, content, userID string, tags, images []string) error {
	var article table.Articles
	article.ID = articleID
	article.Title = title
	article.Description = description
	article.UserID = userID
	article.Contents = content
	article.Nice = 0

	str := "INSERT INTO article_tags (`id`,`article_id`,`tag`) VALUES"
	var s []string
	var queryEnd string

	for i, tag := range tags {
		if len(tags)-1 == i {
			queryEnd = ";"
		} else {
			queryEnd = ","
		}
		tagID, err := uuid.NewRandom()
		if err != nil {
			log.Println(err)
			return err
		}
		q := fmt.Sprintf("('%s', '%s','%s')%s", tagID.String(), articleID, tag, queryEnd)
		s = append(s, q)
	}

	query := strings.Join(s, "")

	err := a.mysql.Client.Transaction(func(tx *gorm.DB) error {
		//記事を追加
		if err := tx.Create(&article).Error; err != nil {
			log.Println(err)
			return err
		}
		//タグを追加
		if err := tx.Exec(str + query).Error; err != nil {
			log.Println(err)
			return err
		}
		// nilが返却されるとトランザクション内の全処理がコミットされる
		return nil
	})

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (a *articlePersistence) FixArticle(articleID, title, description, content, userID string, tags, images []string) error {
	var article table.Articles
	if err := a.mysql.Client.Raw("SELECT * FROM articles WHERE id = ? AND user_id = ?", articleID, userID).Scan(&article).Error; err != nil {
		log.Println(err)
		return err
	}
	if article.Title == "" {
		err := errors.New("this request user has not this article.")
		log.Println(err)
		return err
	}

	if content != "" {
		article.Contents = content
	}

	if title != "" {
		article.Title = title
	}

	if description != "" {
		article.Description = description
	}
	log.Println("aa")
	if err := a.mysql.Client.Save(&article).Error; err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (a *articlePersistence) DeleteArticle(articleID, userID string) error {
	var article table.Articles
	if err := a.mysql.Client.Raw("SELECT * FROM articles WHERE id = ? AND user_id = ?", articleID, userID).Scan(&article).Error; err != nil {
		log.Println(err)
		return err
	}
	if article.Title == "" {
		err := errors.New("this request user has not this article.")
		log.Println(err)
		return err
	}

	if err := a.mysql.Client.Delete(&article).Error; err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (a *articlePersistence) SearchArticles() error {
	return nil
}

func (a *articlePersistence) SendArticle() error {
	return nil
}
