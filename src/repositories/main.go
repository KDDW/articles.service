package repositories

import (
	"kddw.article.service/src/db"
	"kddw.article.service/src/domain/ports"
	article_repo "kddw.article.service/src/repositories/article"
)

type Repositories struct {
	ArticleRepo ports.ArticleRepo
}

func NewRepositories(db *db.DB) *Repositories {
	return &Repositories{
		ArticleRepo: article_repo.NewArticleRepository(db),
	}
}
