package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID        *uuid.UUID     `json:"id" gorm:"primary_key,type:uuid"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (base *Base) BeforeCreate(tx *gorm.DB) (err error) {
	if base.ID == nil {
		id := uuid.New()
		base.ID = &id
	}
	return
}
