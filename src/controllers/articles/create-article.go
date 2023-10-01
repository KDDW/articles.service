package articles_controller

import (
	"github.com/gofiber/fiber/v2"
	"kddw.article.service/src/dtos"
	"kddw.article.service/src/utils"
)

func (a *ArticleController) CreateArticle(c *fiber.Ctx) error {

	in := &dtos.CreateArticle{}

	if err := c.BodyParser(in); err != nil {
		return utils.NewAppError(400, err.Error()).ToFiberError()
	}

	if err := utils.ValidateStruct(in); err != nil {
		return utils.NewAppValidationError(400, err).ToFiberError()
	}

	entryCreated, err := a.services.CreateArticle(in)

	if err != nil {
		return err.ToFiberError()
	}

	return c.JSON(entryCreated)
}
