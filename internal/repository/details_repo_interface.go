package repository

import (
	"context"

	"github.com/laurati/projeto01/internal/entity"
)

type Repo interface {
	CreateBundleDetails(ctx context.Context, BundleDetail *entity.DetailsDtoInput) (map[string]interface{}, error)
	// ReadAll(ctx context.Context) ([]map[string]interface{}, error)
}
