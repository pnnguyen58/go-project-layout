package models

import (
	"github.com/google/uuid"
	"time"
)

type Outbox struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Topic     string    `json:"topic" gorm:"column:topic;type:varchar(255);default:'common'"`
	Payload   string    `json:"payload" gorm:"column:payload;type:json"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime;column:created_at;type:timestamp;not null"`
	CreatedBy string    `json:"createdBy" gorm:"column:created_by;type:varchar(255);default:'system'"`
}
