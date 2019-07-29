package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/sirupsen/logrus"
)

var (
	dbFile = "run/holdem.db"

	db *gorm.DB
)

func init() {
	var err error
	db, err = gorm.Open("sqlite3", dbFile)
	if err != nil {
		logrus.Panic(err)
	}
	db.SingularTable(true)

	db.AutoMigrate(&Room{})
	db.AutoMigrate(&Player{})
	db.AutoMigrate(&Game{})
	db.AutoMigrate(&RelationGamePlayer{})
	db.AutoMigrate(&GameRecord{})
}
