package main

import (
	"github.com/laurati/projeto01/internal/configuration"
	"github.com/laurati/projeto01/internal/database"
	"gorm.io/driver/postgres"
)

func main() {

	configuration.EnvironmentSetup()

	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()

	dbPostgres := database.ConnectDatabase(postgres.Open(configuration.GetPostgresConnectionString()))
	database.Migrate(dbPostgres)
}
