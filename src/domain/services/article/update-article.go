package article_services

import (
	"kddw.article.service/src/domain/entities"
	"kddw.article.service/src/dtos"
	"kddw.article.service/src/utils"
)

func (a *ArticleServices) UpdateArticle(id int64, dto *dtos.UpdateArticleDto) (*entities.Article, *utils.AppError) {

	updatedEntry, err := a.articleRepo.UpdateArticle(id, dto)

	if err != nil {
		if err.Error() == "Article not found" {
			return nil, utils.NewAppError(404, err.Error())
		}

		return nil, utils.NewAppError(500, err.Error())
	}

	return updatedEntry, nil
}
