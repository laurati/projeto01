package domain

import (
	"log"
	"time"

	"github.com/laurati/projeto01/internal/utils"
	"gorm.io/gorm"
)

type Details struct {
	ID         string `gorm:"primary_key"`
	Version    int64
	BundleType string
	Status     string
	CreatedBy  string
	CreatedAt  time.Time
}

func (d *Details) BeforeCreate(tx *gorm.DB) (err error) {
	if d.ID == "" {
		id, err := utils.GenerateUlid(tx.Statement.Context)
		if err != nil {
			log.Println("error generating details ulid ", err.Error())
			return err
		}
		d.ID = id.String()
	}

	if d.Version == 1 {
		d.Status = "new"
	} else {
		d.Status = "not new"
	}
	return nil
}
