package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Base struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	CreatedAt time.Time      `json:"createdAt" gorm:"autoCreateTime;column:created_at;type:timestamp;not null"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"autoUpdateTime;column:updated_at;type:timestamp;default:now();not null"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"column:deleted_at;type:timestamp"`
	CreatedBy string         `json:"createdBy" gorm:"column:created_by;type:varchar(255);default:'system';index:idx_created_by"`
	UpdatedBy string         `json:"updatedBy" gorm:"column:updated_by;type:varchar(255);default:'system'"`
	DeletedBy *string        `json:"deletedBy" gorm:"column:deleted_by;type:varchar(255)"`
}
