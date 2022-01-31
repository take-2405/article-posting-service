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

	bulkInsertTagQuery := "INSERT INTO article_tags (`id`,`article_id`,`tag`) VALUES"
	bulkInsertImageQuery := "INSERT INTO article_images (`id`,`article_id`,`image`) VALUES"
	var bulkInsertTag []string
	var bulkInsertImage []string
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
		bulkInsertTag = append(bulkInsertTag, q)
	}
	bulkInsertTagQuery += strings.Join(bulkInsertTag, "")

	for i, image := range images {
		if len(images)-1 == i {
			queryEnd = ";"
		} else {
			queryEnd = ","
		}
		imageID, err := uuid.NewRandom()
		if err != nil {
			log.Println(err)
			return err
		}
		q := fmt.Sprintf("('%s', '%s','%s')%s", imageID.String(), articleID, image, queryEnd)
		bulkInsertImage = append(bulkInsertImage, q)
	}
	bulkInsertImageQuery += strings.Join(bulkInsertImage, "")

	err := a.mysql.Client.Transaction(func(tx *gorm.DB) error {
		//記事を追加
		if err := tx.Create(&article).Error; err != nil {
			log.Println(err)
			return err
		}
		//タグを追加
		if err := tx.Exec(bulkInsertTagQuery).Error; err != nil {
			log.Println(err)
			return err
		}
		//画像を追加
		if err := tx.Exec(bulkInsertImageQuery).Error; err != nil {
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

	if content != "" {
		article.Contents = content
	}

	if title != "" {
		article.Title = title
	}

	if description != "" {
		article.Description = description
	}

	if article.Title == "" {
		err := errors.New("this request user has not this article.")
		log.Println(err)
		return err
	}
	if len(tags) > 0 {
		bulkInsertTagQuery := "INSERT INTO article_tags (`id`,`article_id`,`tag`) VALUES"
		var bulkInsertTag []string
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
			bulkInsertTag = append(bulkInsertTag, q)
		}
		bulkInsertTagQuery += strings.Join(bulkInsertTag, "")

		a.mysql.Client.Transaction(func(tx *gorm.DB) error {
			var deleteTag table.ArticleTag
			deleteTag.ArticleID = articleID
			if err := tx.Delete(&deleteTag).Error; err != nil {
				log.Println(err)
				return err
			}

			if err := tx.Exec(bulkInsertTagQuery).Error; err != nil {
				log.Println(err)
				return err
			}
			return nil
		})
	}

	if len(images) > 0 {
		bulkInsertImageQuery := "INSERT INTO article_images (`id`,`article_id`,`image`) VALUES"
		var bulkInsertImage []string
		var queryEnd string

		for i, image := range images {
			if len(images)-1 == i {
				queryEnd = ";"
			} else {
				queryEnd = ","
			}
			imageID, err := uuid.NewRandom()
			if err != nil {
				log.Println(err)
				return err
			}
			q := fmt.Sprintf("('%s', '%s','%s')%s", imageID.String(), articleID, image, queryEnd)
			bulkInsertImage = append(bulkInsertImage, q)
		}
		bulkInsertImageQuery += strings.Join(bulkInsertImage, "")

		a.mysql.Client.Transaction(func(tx *gorm.DB) error {
			var deleteImage table.ArticleImage
			deleteImage.ArticleID = articleID
			if err := tx.Delete(&deleteImage).Error; err != nil {
				log.Println(err)
				return err
			}
			if err := tx.Exec(bulkInsertImageQuery).Error; err != nil {
				log.Println(err)
				return err
			}
			return nil
		})
	}

	if err := a.mysql.Client.Save(&article).Error; err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (a *articlePersistence) DeleteArticle(articleID, userID string) error {
	var article table.Articles
	var image table.ArticleImage
	var tag table.ArticleTag
	image.ArticleID = articleID
	tag.ArticleID = articleID

	if err := a.mysql.Client.Raw("SELECT * FROM articles WHERE id = ? AND user_id = ?", articleID, userID).Scan(&article).Error; err != nil {
		log.Println(err)
		return err
	}
	if article.Title == "" {
		err := errors.New("this request user has not this article.")
		log.Println(err)
		return err
	}

	a.mysql.Client.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&article).Error; err != nil {
			log.Println(err)
			return err
		}
		if err := tx.Delete(&image).Error; err != nil {
			log.Println(err)
			return err
		}
		if err := tx.Delete(&tag).Error; err != nil {
			log.Println(err)
			return err
		}
		return nil
	})

	return nil
}

func (a *articlePersistence) SearchArticles() error {
	return nil
}

func (a *articlePersistence) SendArticle() error {
	return nil
}
