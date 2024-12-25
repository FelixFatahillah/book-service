package dtos

import (
	"book-service/internal/domain/book/models"
	"time"
)

type CreateBookDto struct {
	Title      string    `json:"title"`
	Genre      string    `json:"genre"`
	Stock      uint      `json:"stock"`
	Published  time.Time `json:"published"`
	AuthorID   string    `json:"author_id,omitempty"`
	CategoryID string    `json:"category_id,omitempty"`
}

type CreateBookResponseDto struct {
	Title     string          `json:"title"`
	Genre     string          `json:"genre"`
	Stock     uint            `json:"stock"`
	Published time.Time       `json:"published"`
	Author    models.Author   `json:"author,omitempty"`
	Category  models.Category `json:"category,omitempty"`
}

type UpdateBookDto struct {
	ID        string          `json:"id"`
	Title     string          `json:"title"`
	Genre     string          `json:"genre"`
	Stock     uint            `json:"stock"`
	Published time.Time       `json:"published"`
	Author    models.Author   `json:"author,omitempty"`
	Category  models.Category `json:"category,omitempty"`
}
