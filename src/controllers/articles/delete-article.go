package articles_controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"kddw.article.service/src/utils"
)

func (a *ArticleController) DeleteArticle(c *fiber.Ctx) error {

	id := c.Params("id")

	idNbr, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		return utils.NewAppError(400, err.Error()).ToFiberError()
	}

	appErr := a.services.DeleteArticle(idNbr)

	if appErr != nil {
		return appErr.ToFiberError()
	}

	return c.SendStatus(204)
}
