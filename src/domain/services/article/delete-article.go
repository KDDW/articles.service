package article_services

import "kddw.article.service/src/utils"

func (a *ArticleServices) DeleteArticle(id int64) *utils.AppError {

	err := a.articleRepo.DeleteArticle(id)

	if err != nil {
		if err.Error() == "Article not found" {
			return utils.NewAppError(404, err.Error())
		}

		return utils.NewAppError(500, err.Error())
	}

	return nil
}
