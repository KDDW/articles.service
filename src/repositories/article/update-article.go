package article_repo

import (
	"errors"

	"kddw.article.service/src/domain/entities"
	"kddw.article.service/src/dtos"
)

func (r *ArticleRepository) UpdateArticle(id int64, dto *dtos.UpdateArticleDto) (*entities.Article, error) {

	db := r.Db.Db

	res := &entities.Article{}

	db.First(res, id)

	if res == nil || res.Id == 0 {
		return nil, errors.New("Article not found")
	}

	if dto.Title != "" {
		res.Title = dto.Title
	}

	if dto.Content != "" {
		res.Content = dto.Content
	}

	db.Save(res)

	return res, nil
}
