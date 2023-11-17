package database

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
)

func TestConnectDatabase(t *testing.T) {

	dialector := sqlite.Open(":memory:")
	db := ConnectDatabase(context.Background(), dialector)
	assert.NotNil(t, db)
	db2 := ConnectDatabase(context.Background(), dialector)
	assert.Equal(t, db, db2)
}
