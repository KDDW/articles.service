package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func newGORMDB() *gorm.DB {

	// dsn := os.Getenv("DATABASE_URL")
	//
	// if dsn == "" {
	// 	panic("DATABASE_URL is not set")
	// }

	db, err := gorm.Open(sqlite.Open("articles.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}
