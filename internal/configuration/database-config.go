package configuration

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/laurati/projeto01/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(ctx context.Context) (*gorm.DB, error) {
	consStr := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))

	db, err := gorm.Open(postgres.Open(consStr), &gorm.Config{})
	if err != nil {
		log.Fatal("Error when connect db" + consStr + " : " + err.Error())
		return nil, err
	}

	if err := db.Debug().AutoMigrate(
		domain.Details{},
	); err != nil {
		log.Fatal("Error when automigrate db: " + err.Error())
		return nil, err
	}

	log.Println("Host", os.Getenv("DB_HOST"))
	log.Println("Port", os.Getenv("DB_PORT"))
	log.Println("User", os.Getenv("DB_USER"))
	log.Println("Dbname", os.Getenv("DB_NAME"))
	log.Println("Connected to Database Successfully!")

	return db, nil
}
