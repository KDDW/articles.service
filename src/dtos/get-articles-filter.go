package dtos

type GetArticlesFilter struct {
	RecordsPerPage int    `json:"recordsPerPage"`
	Page           int    `json:"page"`
	AuthorId       int64  `json:"authorId"`
	Title          string `json:"title"`
}
