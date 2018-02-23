package common

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func DataBase() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		log.Fatal(err)
	}

	return db
}
