package repository

import (
	"mansion/model"
	"time"
)

type RoomType struct {
	model.Base
	Name string `gorm:"not null; type:varchar(255)"`
}

type UpdateRoomType struct {
	Name      string
	UpdatedAt time.Time
}

type RoomTypeRepository interface {
	Create(room *RoomType) (result *RoomType, err error)
	GetAll() ([]RoomType, error)
	GetById(id string) (result *RoomType, err error)
	Update(id string, room *UpdateRoomType) error
	Delete(id string) error
}
