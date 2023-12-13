package main

import (
	"context"
	"log"
	"os"

	"github.com/laurati/projeto01/internal/configuration"
	"github.com/laurati/projeto01/internal/database"
	"github.com/laurati/projeto01/internal/handler"
	"github.com/laurati/projeto01/internal/repository"
	"github.com/laurati/projeto01/internal/router"
	"github.com/laurati/projeto01/internal/server"
	"github.com/laurati/projeto01/internal/usecase"
	"gorm.io/driver/postgres"
)

func main() {

	configuration.EnvironmentSetup()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	log.Printf("Server Port: %v", os.Getenv("PORT"))
	log.Println("Starting process API...")

	postgresDsn := configuration.GetPostgresConnectionString()
	dbPostgres := database.ConnectDatabase(ctx, postgres.Open(postgresDsn))
	database.Migrate(dbPostgres)

	repo := repository.NewDetailsRepo(dbPostgres)
	useCase := usecase.NewDetailsUseCase(repo)
	handler := handler.NewDetailsHandler(useCase)

	router := router.InitializeRouter(handler)
	server := server.NewServer(":"+os.Getenv("PORT"), router)
	server.Start()

}
