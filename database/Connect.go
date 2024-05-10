package database

import (
	"log"
	"os"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var Handle *gorm.DB

func Connect() {
	db, err := gorm.Open(sqlite.Open(os.Getenv("DB")), &gorm.Config{})
	if err != nil {
		log.Fatal("Couldn't connect to the database")
	}

	Handle = db // Prevent unused variable
}
