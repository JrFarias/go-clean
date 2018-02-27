package common

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" //_ sqlite dialect
)

//DB global variable to access DB
var DB *gorm.DB

//DataBase function to open connection
func DataBase() {
	db, err := gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		log.Fatal(err)
	}

	db.DB().SetMaxIdleConns(0)
	db.DB().SetConnMaxLifetime(5 * time.Second)
	DB = db
}
