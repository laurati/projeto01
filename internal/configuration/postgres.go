package configuration

import (
	"fmt"
	"os"
)

func GetPostgresConnectionString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_POSTGRES_HOST"),
		os.Getenv("DB_POSTGRES_PORT"),
		os.Getenv("DB_POSTGRES_USER"),
		os.Getenv("DB_POSTGRES_PASSWORD"),
		os.Getenv("DB_POSTGRES_NAME"))
}
