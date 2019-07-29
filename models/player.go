package models

func GetPlayer(name string) (*Player, error) {
	var rec Player
	err := db.Where("username = ?", name).First(&rec).Error
	if err != nil {
		return nil, err
	}
	return &rec, nil
}
