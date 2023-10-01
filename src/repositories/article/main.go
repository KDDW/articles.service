package article_repo

import (
	"kddw.article.service/src/db"
)

type ArticleRepository struct {
	Db *db.DB
}

/* type ArticleRepo interface {
	GetArticles(dto *dtos.GetArticlesFilter) ([]entities.Article, error)
	GetArticleById(id uint64) (*entities.Article, error)
	CreateArticle(article *entities.Article) (*entities.Article, error)
	UpdateArticle(id uint64, dto *dtos.UpdateArticleDto) (*entities.Article, error)
	DeleteArticle(id uint64) error
} */

func NewArticleRepository(db *db.DB) *ArticleRepository {
	return &ArticleRepository{
		Db: db,
	}
}
