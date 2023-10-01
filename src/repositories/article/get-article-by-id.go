package article_repo

import (
	"errors"

	"kddw.article.service/src/domain/entities"
)

func (r *ArticleRepository) GetArticleById(id int64) (*entities.Article, error) {

	db := r.Db.Db

	res := &entities.Article{}

	db.First(res, id)

	if res == nil || res.Id == 0 {
		return nil, errors.New("Article not found")
	}

	return res, nil
}
