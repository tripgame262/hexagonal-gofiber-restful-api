package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Base struct {
	UUID      string `gorm:"primaryKey;autoIncrement:false;type:varchar(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (b *Base) BeforeCreate(tx *gorm.DB) (err error) {
	b.UUID = uuid.NewV4().String()

	return nil
}
