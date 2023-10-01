package articles_controller

import (
	"github.com/gofiber/fiber/v2"
	"kddw.article.service/src/dtos"
	"kddw.article.service/src/utils"
)

func (a *ArticleController) ListArticles(c *fiber.Ctx) error {

	dto := &dtos.GetArticlesFilter{}

	if err := c.QueryParser(dto); err != nil {
		return utils.NewAppError(400, err.Error()).ToFiberError()
	}

	items, appErr := a.services.GetArticles(dto)

	if appErr != nil {
		return appErr.ToFiberError()
	}

	return c.JSON(items)
}
