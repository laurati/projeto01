package usecase

import (
	"context"

	"github.com/laurati/projeto01/internal/domain"
)

type Usecase interface {
	CreateDetail(ctx context.Context, bundleDetailsDtoInput *domain.DetailsDtoInput) (map[string]interface{}, error)
	// ReadAll(ctx context.Context) ([]map[string]interface{}, error)
}
