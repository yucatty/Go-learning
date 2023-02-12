package services

import "github.com/yucatty/Go-learning/myapi/models"

type MyAppServicer interface {
	PostArticleService(article models.Article) (models.Article, error)
	ArticleListService(page int) ([]models.Article, error)
	GetArticleService(articleID int) (models.Article, error)
	PostNiceService(article models.Article) (models.Article, error)

	PostCommentService(comment models.Comment) (models.Comment, error)
}
