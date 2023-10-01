package article_services

import (
	"kddw.article.service/src/dtos"
	"kddw.article.service/src/utils"
)

func (a *ArticleServices) GetArticles(dto *dtos.GetArticlesFilter) (*dtos.GetArticlesResponse, *utils.AppError) {

	results, err := a.articleRepo.GetArticles(dto)

	if err != nil {
		return nil, utils.NewAppError(500, err.Error())
	}

	return results, nil
}
