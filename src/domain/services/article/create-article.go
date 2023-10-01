package article_services

import (
	"kddw.article.service/src/domain/entities"
	"kddw.article.service/src/dtos"
	"kddw.article.service/src/utils"
)

func (a *ArticleServices) CreateArticle(dto *dtos.CreateArticle) (*entities.Article, *utils.AppError) {

	createdEntry, err := a.articleRepo.CreateArticle(&entities.Article{
		Title:    dto.Title,
		Content:  dto.Content,
		AuthorId: dto.AuthorId,
	})

	if err != nil {
		return nil, utils.NewAppError(500, err.Error())
	}

	if createdEntry == nil || createdEntry.Id == 0 {
		return nil, utils.NewAppError(500, "Article cannot be created")
	}

	return createdEntry, nil
}
