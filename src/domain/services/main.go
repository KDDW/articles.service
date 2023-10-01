package services

import (
	"kddw.article.service/src/domain/ports"
	article_services "kddw.article.service/src/domain/services/article"
	"kddw.article.service/src/repositories"
)

type Services struct {
	ArticleServices ports.ArticleServices
}

func NewServices(repos *repositories.Repositories) *Services {
	return &Services{
		ArticleServices: article_services.NewArticleService(repos),
	}
}
