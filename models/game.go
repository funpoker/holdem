package models

func CreateGame(roomId uint) error {
	rec := Game{
		RoomId: roomId,
		Period: 0,
		Status: 1,
	}
	return db.Create(rec).Error
}

func GetGame(roomId uint64) (*Game, error) {
	var rec Game
	err := db.Where("room_id = ?", roomId).First(&rec).Error
	if err != nil {
		return nil, err
	}
	return &rec, nil
}

func UpdateGame(updates map[string]interface{}) error {
	return db.Model(&Game{}).Updates(updates).Error
}
