package dtos

type CreateArticle struct {
	Title    string `json:"title" validate:"required"`
	Content  string `json:"content" validate:"required"`
	AuthorId int64  `json:"author_id" validate:"required,min=1"`
}
