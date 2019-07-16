package models

import "time"

type Room struct {
	Id            uint64
	Name          string
	BackgroundUrl string
	CreatedTime   time.Time
	UpdatedTime   time.Time
	Status        int
}

type Player struct {
	Id          uint64
	Username    string
	Password    string
	AvatarUrl   string
	chip        uint64
	CreatedTime time.Time
	UpdatedTime time.Time
	Status      int
}

type Game struct {
	Id        uint64
	RoomId    uint64
	Period    uint64
	Round     uint64
	Flop      string
	Turn      string
	River     string
	BeginTime time.Time
	EndTime   time.Time
	Status    int
}

type RelationGamePlayer struct {
	Id             uint64
	GameId         uint64
	PlayerId       uint64
	PlayerRole     int
	PlayerPosition int
	Hands          string
	CreatedTime    time.Time
	UpdatedTime    time.Time
	Status         int
}

type GameRecord struct {
	Id          uint64
	GameId      uint64
	PlayerId    uint64
	Round       uint64
	Phase       int
	Amount      uint64
	CreatedTime time.Time
}
