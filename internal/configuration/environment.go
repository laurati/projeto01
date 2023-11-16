package configuration

import (
	"log"

	"github.com/joho/godotenv"
)

func EnvironmentSetup() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading env file", err)
		return
	}
	log.Println("YOUR ENV FILE IS LOADED")
}
