package models

import (
	"time"
)

type BaseModel struct {
	ID string `gorm:"column:id;primaryKey;autoIncrement;" json:"-"`
}

type CommonTimestampsField struct {
	CreatedAt time.Time `gorm:"column:created_at;index;" json:"createdAt,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_at;index;" json:"updatedAt,omitempty"`
}

