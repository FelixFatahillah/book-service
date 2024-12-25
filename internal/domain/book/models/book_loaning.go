package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/plugin/optimisticlock"
	"time"
)

type BookLoaning struct {
	ID                 string                 `json:"id" gorm:"type:uuid;primaryKey;not null"`
	CreatedDate        time.Time              `json:"created_date,omitempty" gorm:"type:timestamptz;not null;default:current_timestamp"`
	CreatedBy          string                 `json:"created_by,omitempty" gorm:"type:varchar(255);not null;default:'system'"`
	UpdatedDate        time.Time              `json:"updated_date,omitempty" gorm:"type:timestamptz;not null;default:current_timestamp"`
	UpdatedBy          string                 `json:"updated_by,omitempty" gorm:"type:varchar(255);default:'system'"`
	CustomerName       string                 `json:"customer_name" gorm:"type:varchar(255);not null"`
	LoanDate           time.Time              `json:"loan_date" gorm:"type:timestamptz"`
	ReturnDateSchedule time.Time              `json:"return_date_schedule" gorm:"type:timestamptz"`
	ReturnDate         *time.Time             `json:"return_date" gorm:"type:timestamptz"`
	BookID             string                 `json:"book_id" gorm:"type:varchar(255);not null"`
	BookTitle          string                 `json:"book_title" gorm:"type:varchar(255);not null"`
	Version            optimisticlock.Version `json:"version,omitempty"`
}

func (u *BookLoaning) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == "" {
		u.ID = uuid.New().String() // Generate UUID
	}
	return
}
