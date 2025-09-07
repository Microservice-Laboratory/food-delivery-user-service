package entity

import (
	"time"

	"gorm.io/gorm"

	"github.com/google/uuid"
)

type BaseEntity struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (base *BaseEntity) BeforeCreate(tx *gorm.DB) (err error) {
	base.ID = uuid.New()
	return
}
