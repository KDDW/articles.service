package server

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"kddw.article.service/src/controllers"
)

func RegisterControllers(app *fiber.App, c *controllers.Controllers) {

	router := app.Group("/api/articles")

	router.Post("", c.Article.CreateArticle)
	router.Get(":id", c.Article.GetArticleById)
	router.Get("", c.Article.ListArticles)
	router.Patch(":id", c.Article.UpdateArticle)
	router.Delete(":id", c.Article.DeleteArticle)

}

func CreateServer() *fiber.App {
	config := fiber.Config{
		AppName:           "Articles Service",
		EnablePrintRoutes: true,
	}

	app := fiber.New(config)

	return app
}

func Listen(app *fiber.App) {

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

    err := app.Listen(":" + port)
    if err != nil {
        panic(err)
    }
}
