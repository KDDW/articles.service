package articles_controller

import (
	"kddw.article.service/src/domain/ports"
	"kddw.article.service/src/domain/services"
)

type ArticleController struct {
	services ports.ArticleServices
}

func NewArticleController(s *services.Services) *ArticleController {
	return &ArticleController{
		services: s.ArticleServices,
	}
}
