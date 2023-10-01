package articles_controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"kddw.article.service/src/utils"
)

func (a *ArticleController) GetArticleById(c *fiber.Ctx) error {

	id := c.Params("id")

	idNbr, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		return utils.NewAppError(400, err.Error()).ToFiberError()
	}

	entryCreated, appErr := a.services.GetArticleById(idNbr)

	if appErr != nil {
		return appErr.ToFiberError()
	}

	return c.JSON(entryCreated)
}
