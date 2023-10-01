package article_services

import (
	"kddw.article.service/src/domain/ports"
	"kddw.article.service/src/repositories"
)

/* type ArticleServices interface {
	GetArticles(dto *dtos.GetArticlesFilter) ([]entities.Article, error)
	GetArticleById(id uint64) (*entities.Article, error)
	CreateArticle(article *entities.Article) (*entities.Article, error)
	UpdateArticle(dto *dtos.UpdateArticleDto) (*entities.Article, error)
	DeleteArticle(id uint64) error
} */

type ArticleServices struct {
	articleRepo ports.ArticleRepo
}

func NewArticleService(repo *repositories.Repositories) *ArticleServices {
	return &ArticleServices{
		articleRepo: repo.ArticleRepo,
	}
}
