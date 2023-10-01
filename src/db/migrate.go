package db

import (
	"gorm.io/gorm"
	"kddw.article.service/src/domain/entities"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&entities.Article{})
}
