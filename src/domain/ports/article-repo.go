package ports

import (
	"kddw.article.service/src/domain/entities"
	"kddw.article.service/src/dtos"
)

type ArticleRepo interface {
	GetArticles(dto *dtos.GetArticlesFilter) (*dtos.GetArticlesResponse, error)
	GetArticleById(id int64) (*entities.Article, error)
	CreateArticle(article *entities.Article) (*entities.Article, error)
	UpdateArticle(id int64, dto *dtos.UpdateArticleDto) (*entities.Article, error)
	DeleteArticle(id int64) error
}
