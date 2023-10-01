package article_services

import (
	"kddw.article.service/src/domain/entities"
	"kddw.article.service/src/utils"
)

func (a *ArticleServices) GetArticleById(id int64) (*entities.Article, *utils.AppError) {

	entry, err := a.articleRepo.GetArticleById(id)

	if err != nil {
		return nil, utils.NewAppError(404, err.Error())
	}

	if entry == nil || entry.Id == 0 {
		return nil, utils.NewAppError(404, "Article not found")
	}

	return entry, nil
}
