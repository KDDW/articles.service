package articles_controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"kddw.article.service/src/dtos"
	"kddw.article.service/src/utils"
)

func (a *ArticleController) UpdateArticle(c *fiber.Ctx) error {

	in := &dtos.UpdateArticleDto{}

	if err := c.BodyParser(in); err != nil {
		return utils.NewAppError(400, err.Error()).ToFiberError()
	}

	if err := utils.ValidateStruct(in); err != nil {
		return utils.NewAppValidationError(400, err).ToFiberError()
	}

	id := c.Params("id")

	idNbr, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		return utils.NewAppError(400, err.Error()).ToFiberError()
	}

	entryUpdated, appErr := a.services.UpdateArticle(idNbr, in)

	if appErr != nil {
		return appErr.ToFiberError()
	}

	return c.Status(200).JSON(entryUpdated)
}
