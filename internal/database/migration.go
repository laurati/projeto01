package database

import (
	"log"

	"github.com/laurati/projeto01/internal/entity"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	if err := db.AutoMigrate(
		&entity.Details{},
	); err != nil {
		log.Fatalln("Error when automigrate db: " + err.Error())
	}
	log.Println("Database migration completed!")
}
