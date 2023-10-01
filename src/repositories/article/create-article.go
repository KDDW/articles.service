package article_repo

import (
	"kddw.article.service/src/domain/entities"
)

func (r *ArticleRepository) CreateArticle(article *entities.Article) (*entities.Article, error) {
	db := r.Db.Db

	// now := time.Now()
	toBeCreated := &entities.Article{
		Title:    article.Title,
		Content:  article.Content,
		AuthorId: article.AuthorId,
		// CreatedAt: &now,
		// UpdatedAt: &now,
	}

	db.Create(&toBeCreated)
	return toBeCreated, nil
}
