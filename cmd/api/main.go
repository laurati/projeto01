package main

import (
	"context"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/laurati/projeto01/internal/configuration"
)

func main() {

	environmentSetup()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db, err := configuration.Connect(ctx)
	if err != nil {
		log.Println("error when create connection DB", err.Error())
		return
	}

	fmt.Println(db)

}

func environmentSetup() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading env file", err)
		return
	}
	log.Println("YOUR ENV FILE IS LOADED")
}
