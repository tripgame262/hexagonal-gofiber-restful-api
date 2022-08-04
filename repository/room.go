package repository

import (
	"mansion/model"
	"time"
)

type Room struct {
	model.Base
	RoomID     string `gorm:"not null;unique;type:varchar(255)"`
	RoomPrice  int16  `gorm:"not null"`
	Floor      int16  `gorm:"not null"`
	RoomType   RoomType
	RoomTypeID string
}

type UpdateRoom struct {
	RoomID     string
	RoomPrice  int16
	Floor      int16
	RoomTypeID string
	UpdatedAt  time.Time
}

type RoomRepository interface {
	Create(room *Room) (result *Room, err error)
	GetAll() ([]Room, error)
	GetById(id string) (result *Room, err error)
	Update(id string, room *UpdateRoom) error
	Delete(id string) error
}
