package article_repo

import (
	"errors"

	"kddw.article.service/src/domain/entities"
)

func (r *ArticleRepository) DeleteArticle(id int64) error {

	db := r.Db.Db

	res := &entities.Article{}

	db.First(res, id)

	if res == nil || res.Id == 0 {
		return errors.New("Article not found")
	}

	db.Delete(res, id)

	return nil
}
