package repository

import (
	"context"

	"github.com/laurati/projeto01/internal/domain"
)

type Repo interface {
	Create(ctx context.Context, BundleDetail *domain.DetailsDtoInput) (map[string]interface{}, error)
	// ReadAll(ctx context.Context) ([]map[string]interface{}, error)
}
