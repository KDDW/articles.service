package dtos

type GetArticleById struct {
	Id int64 `json:"id" validate:"required"`
}
