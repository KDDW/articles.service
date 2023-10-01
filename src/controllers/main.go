package controllers

import (
	articles_controller "kddw.article.service/src/controllers/articles"
	"kddw.article.service/src/domain/services"
)

type Controllers struct {
	Article *articles_controller.ArticleController
}

func NewControllers(s *services.Services) *Controllers {
	return &Controllers{
		Article: articles_controller.NewArticleController(s),
	}
}
