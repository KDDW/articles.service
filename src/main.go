package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"kddw.article.service/src/controllers"
	"kddw.article.service/src/db"
	"kddw.article.service/src/domain/services"
	"kddw.article.service/src/repositories"
	"kddw.article.service/src/server"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

func main() {

	database := db.NewDB()
	db.Migrate(database.Db)

	repos := repositories.NewRepositories(database)

	s := services.NewServices(repos)
	c := controllers.NewControllers(s)

	app := server.CreateServer()
	server.RegisterControllers(app, c)
	server.Listen(app)
}
