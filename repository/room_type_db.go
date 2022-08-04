package repository

import (
	"gorm.io/gorm"
)

type roomTypeRepository struct {
	db *gorm.DB
}

func NewRoomTypeRepository(db *gorm.DB) RoomTypeRepository {
	return roomTypeRepository{db: db}
}

func (r roomTypeRepository) Create(roomType *RoomType) (result *RoomType, err error) {
	tx := r.db.Create(roomType)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return roomType, nil
}

func (r roomTypeRepository) GetAll() ([]RoomType, error) {
	rooms := []RoomType{}
	tx := r.db.Find(&rooms)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return rooms, nil
}

func (r roomTypeRepository) GetById(id string) (result *RoomType, err error) {
	room := RoomType{}

	tx := r.db.Where("uuid=?", id).First(&room)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return &room, nil
}

func (r roomTypeRepository) Update(id string, data *UpdateRoomType) error {

	tx := r.db.Model(RoomType{}).Where("uuid = ?", id).Updates(&data)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (r roomTypeRepository) Delete(id string) error {
	tx := r.db.Where("uuid = ?", id).Unscoped().Delete(&RoomType{})

	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
