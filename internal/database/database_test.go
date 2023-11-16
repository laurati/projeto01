package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
)

func TestConnectDatabase(t *testing.T) {

	dialector := sqlite.Open(":memory:")
	db := ConnectDatabase(dialector)
	assert.NotNil(t, db)
	db2 := ConnectDatabase(dialector)
	assert.Equal(t, db, db2)
}
