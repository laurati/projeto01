package database

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestMigrate(t *testing.T) {

	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error opening test database:", err)
	}
	Migrate(db)
	assert.True(t, db.Migrator().HasTable(Details{}), "Details table should exist")
}

type Details struct {
	ID         string `gorm:"primary_key"`
	Version    int64
	BundleType string
}

func (Details) TableName() string {
	return "bundle_details"
}
