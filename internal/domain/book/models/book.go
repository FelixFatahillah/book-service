package models

import (
	"book-service/internal/shared"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/plugin/optimisticlock"
	"time"
)

type Book struct {
	ID string `json:"id" gorm:"type:uuid;primaryKey;not null"`
	shared.BaseModel
	Title     string                 `json:"title" gorm:"type:varchar(255);not null"`
	Genre     string                 `json:"genre" gorm:"type:varchar(255);not null"`
	Stock     uint                   `json:"stock" gorm:"type:integer;not null;default:0"`
	Published time.Time              `json:"published" gorm:"type:timestamptz"`
	Author    Author                 `json:"author,omitempty" gorm:"embedded;embeddedPrefix:author_"`
	Category  Category               `json:"category,omitempty" gorm:"embedded;embeddedPrefix:category_"`
	Version   optimisticlock.Version `json:"version,omitempty"`
}

type Author struct {
	FirstName   string  `json:"first_name,omitempty" gorm:"type:varchar(255)"`
	LastName    *string `json:"last_name,omitempty" gorm:"type:varchar(255)"`
	PhoneNumber *string `json:"phone_number,omitempty" gorm:"type:varchar(50)"`
	Email       string  `json:"email,omitempty" gorm:"type:varchar(255)"`
}

type Category struct {
	Name        string `json:"name,omitempty" gorm:"type:varchar(255)"`
	Description string `json:"description,omitempty" gorm:"type:varchar(255)"`
}

func (u *Book) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == "" {
		u.ID = uuid.New().String() // Generate UUID
	}
	return
}
