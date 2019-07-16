package sqlite

import (
	"io/ioutil"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

const (
	dbFile  = "run/holdem.db"
	sqlFile = "sqlite.sql"
)

func New() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", dbFile)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadFile(sqlFile)
	if err != nil {
		return nil, err
	}
	db.Debug().Exec(string(data))

	return db, nil
}
