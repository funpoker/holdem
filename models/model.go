package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

const (
	GameStatusEnd   = 0
	GameStatusJoin  = 1
	GameStatusStart = 2
)

type Room struct {
	gorm.Model
	Name          string
	BackgroundUrl string
	Status        int
}

type Player struct {
	gorm.Model
	Username  string
	Password  string
	AvatarUrl string
	chip      uint
	Status    int
}

type Game struct {
	gorm.Model
	RoomId    uint
	Period    uint
	Round     uint
	Flop      string
	Turn      string
	River     string
	BeginTime time.Time
	EndTime   time.Time
	Status    int
}

type RelationGamePlayer struct {
	gorm.Model
	GameId         uint
	PlayerId       uint
	PlayerRole     int
	PlayerPosition int
	Hands          string
	Status         int // 0 end, 1 join, 2 start
}

type GameRecord struct {
	gorm.Model
	GameId   uint
	PlayerId uint
	Round    uint
	Phase    int
	Amount   uint
}
