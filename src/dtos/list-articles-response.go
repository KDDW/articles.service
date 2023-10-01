package dtos

import "kddw.article.service/src/domain/entities"

type GetArticlesResponse struct {
	RecordCount int                `json:"recordCount"`
	Total       int64              `json:"total"`
	Page        int                `json:"page"`
	Articles    []entities.Article `json:"articles"`
}
