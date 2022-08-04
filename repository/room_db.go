package repository

import "gorm.io/gorm"

type roomRepository struct {
	db *gorm.DB
}

func NewRoomRepository(db *gorm.DB) RoomRepository {
	return roomRepository{db: db}
}

func (r roomRepository) Create(data *Room) (result *Room, err error) {
	tx := r.db.Create(&data)

	if tx.Error != nil {
		return nil, tx.Error
	}

	room := Room{}

	query := r.db.Preload("RoomType").Where("uuid = ?", data.UUID).Find(&room)

	if query.Error != nil {
		return nil, query.Error
	}
	return &room, nil
}

func (r roomRepository) GetAll() ([]Room, error) {
	rooms := []Room{}
	tx := r.db.Preload("RoomType").Find(&rooms)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return rooms, nil
}

func (r roomRepository) GetById(id string) (result *Room, err error) {
	room := Room{}

	tx := r.db.Preload("RoomType").Where("uuid=?", id).First(&room)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return &room, nil
}

func (r roomRepository) Update(id string, room *UpdateRoom) error {
	tx := r.db.Model(Room{}).Where("uuid=?", id).Updates(room)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (r roomRepository) Delete(id string) error {
	tx := r.db.Where("uuid = ?", id).Delete(Room{})

	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
