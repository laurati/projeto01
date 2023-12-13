package usecase

import (
	"context"

	"github.com/laurati/projeto01/internal/entity"
	"github.com/laurati/projeto01/internal/repository"
)

type DetailsUseCase struct {
	detailsRepoInterface repository.Repo
}

func NewDetailsUseCase(detailsRepo *repository.DetailsRepo) *DetailsUseCase {
	return &DetailsUseCase{detailsRepo}
}

func (d *DetailsUseCase) CreateBundleDetails(ctx context.Context, bundleDetailsDtoInput *entity.DetailsDtoInput) (map[string]interface{}, error) {
	return d.detailsRepoInterface.CreateBundleDetails(ctx, bundleDetailsDtoInput)
}
