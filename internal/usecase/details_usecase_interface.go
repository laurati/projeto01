package usecase

import (
	"context"

	"github.com/laurati/projeto01/internal/entity"
)

type Usecase interface {
	CreateBundleDetails(ctx context.Context, bundleDetailsDtoInput *entity.DetailsDtoInput) (map[string]interface{}, error)
	// ReadAll(ctx context.Context) ([]map[string]interface{}, error)
}
