package shared

import (
	"time"
)

type BaseModel struct {
	CreatedDate time.Time  `json:"created_date,omitempty" gorm:"type:timestamptz;not null;default:current_timestamp"`
	CreatedBy   string     `json:"created_by,omitempty" gorm:"type:varchar(255);not null;default:'system'"`
	UpdatedDate time.Time  `json:"updated_date,omitempty" gorm:"type:timestamptz;not null;default:current_timestamp"`
	UpdatedBy   string     `json:"updated_by,omitempty" gorm:"type:varchar(255);default:'system'"`
	DeletedDate *time.Time `json:"deleted_date,omitempty" gorm:"type:timestamptz"`
	DeletedBy   *string    `json:"deleted_by,omitempty" gorm:"type:varchar(255)"`
}
