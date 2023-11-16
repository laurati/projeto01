package database

import (
	"log"

	"github.com/laurati/projeto01/internal/domain"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	if err := db.AutoMigrate(
		&domain.Details{},
	); err != nil {
		log.Fatalln("Error when automigrate db: " + err.Error())
	}
	log.Println("Database migration completed!")
}
