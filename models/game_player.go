package models

func CreateRelationGamePlayer(rec *RelationGamePlayer) error {
	return db.CreateTable(rec).Error
}

type GamePlayer struct {
	Id             int
	Username       string
	AvatarUrl      string
	Amount         int
	PlayerRole     int
	PlayerPosition int
	Status         int
}

func ListGamePlayers(gameId uint) ([]*GamePlayer, error) {
	var resp []*GamePlayer
	err := db.Raw("select p.id, p.username, p.avatar_url, p.amount, r.player_role, r.player_position, r.status from player p, relation_game_player r where r.game_id = ? and p.id = r.player_id", gameId).Find(&resp).Error
	return resp, err
}

func UpdateRelationGamePlayer(updates map[string]interface{}) error {
	return db.Model(&RelationGamePlayer{}).Updates(updates).Error
}
