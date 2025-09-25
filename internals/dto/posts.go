package dto

import (
	"github.com/google/uuid"
	"time"
)

// Payload for creating a post
type PostCreate struct {
	Content string    `json:"content" validate:"required"`
}

// Single post response
type Post struct {
	ID        uuid.UUID  `json:"id"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	UserID    uuid.UUID  `json:"user_id"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
type Posts struct {
	Posts []Post `json:"posts"`
}

// Lightweight list view (optional)
type PostListItem struct {
	ID      uuid.UUID `json:"id" gorm:"column:id"`
	Title   string    `json:"title" gorm:"column:title"`
	Content string    `json:"content" gorm:"column:content"`
	UserID  uuid.UUID `json:"user_id" gorm:"column:user_id"`
}

