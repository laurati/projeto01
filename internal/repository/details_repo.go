package repository

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/laurati/projeto01/internal/domain"
	"gorm.io/gorm"
)

type DetailsRepo struct {
	DB *gorm.DB
}

func NewDetailsRepo(DB *gorm.DB) *DetailsRepo {
	return &DetailsRepo{DB}
}

func (b *DetailsRepo) Create(ctx context.Context, bundleDetailsDtoInput *domain.DetailsDtoInput) (map[string]interface{}, error) {

	bundleDetails := domain.Details{
		Version:    bundleDetailsDtoInput.Version,
		BundleType: bundleDetailsDtoInput.BundleType,
		CreatedBy:  bundleDetailsDtoInput.CreatedBy,
		CreatedAt:  time.Now(),
	}

	if err := b.DB.WithContext(ctx).Create(&bundleDetails).Error; err != nil {
		log.Println("error creating a bundle details ", err.Error())
		return nil, err
	}

	var detailResponse map[string]interface{}

	data, err := json.Marshal(bundleDetails)
	if err != nil {
		log.Println("error marshal bundleDetails ", err.Error())
		return nil, err
	}
	err = json.Unmarshal(data, &detailResponse)
	if err != nil {
		log.Println("error unmarshal bundleDetails ", err.Error())
		return nil, err
	}

	return detailResponse, nil
}
