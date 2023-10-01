package article_repo

import (
	"kddw.article.service/src/domain/entities"
	"kddw.article.service/src/dtos"
)

func (r *ArticleRepository) GetArticles(dto *dtos.GetArticlesFilter) (*dtos.GetArticlesResponse, error) {

	db := r.Db.Db

	result := []entities.Article{}

	page := dto.Page
	recordsPerPage := dto.RecordsPerPage

	if page <= 0 {
		page = 1
	}

	if recordsPerPage <= 0 {
		recordsPerPage = 20
	}

	offset := (page - 1) * recordsPerPage

	query := db.Offset(offset).Limit(recordsPerPage)

	if dto.Title != "" {
		// insensitive like search
		query = query.Where("title LIKE ?", "%"+dto.Title+"%")
	}

	if dto.AuthorId != 0 {
		query = query.Where("author_id = ?", dto.AuthorId)
	}

	var totalCount int64 = 0
	query.Find(&result).Count(&totalCount)

	return &dtos.GetArticlesResponse{
		RecordCount: len(result),
		Total:       totalCount,
		Page:        page,
		Articles:    result,
	}, nil
}
