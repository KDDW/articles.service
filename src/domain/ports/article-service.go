package ports

import (
	"kddw.article.service/src/domain/entities"
	"kddw.article.service/src/dtos"
	"kddw.article.service/src/utils"
)

type ArticleServices interface {
	GetArticles(dto *dtos.GetArticlesFilter) (*dtos.GetArticlesResponse, *utils.AppError)
	GetArticleById(id int64) (*entities.Article, *utils.AppError)
	CreateArticle(dto *dtos.CreateArticle) (*entities.Article, *utils.AppError)
	UpdateArticle(id int64, dto *dtos.UpdateArticleDto) (*entities.Article, *utils.AppError)
	DeleteArticle(id int64) *utils.AppError
}
