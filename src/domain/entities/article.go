package entities

import "time"

type Article struct {
	Id        int64      `json:"id" gorm:"primaryKey, autoIncrement"`
	Title     string     `json:"title", gorm:"not null default:''"`
	Content   string     `json:"content" gorm:"not null default:''"`
	AuthorId  int64      `json:"authorId" gorm:"not null"`
	CreatedAt *time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt *time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}
