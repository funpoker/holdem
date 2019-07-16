package models

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"

	"github.com/funpoker/holdem/models/sqlite"
)

var dbClient *gorm.DB

func init() {
	var err error
	dbClient, err = sqlite.New()
	if err != nil {
		logrus.Panic(err)
	}
}
