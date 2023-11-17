package usecase

import (
	"context"

	"github.com/laurati/projeto01/internal/domain"
	"github.com/laurati/projeto01/internal/repository"
)

type DetailsUseCase struct {
	detailsRepoInterface repository.Repo
}

func NewDetailsUseCase(detailsRepo *repository.DetailsRepo) *DetailsUseCase {
	return &DetailsUseCase{detailsRepo}
}

func (d *DetailsUseCase) CreateDetail(ctx context.Context, bundleDetailsDtoInput *domain.DetailsDtoInput) (map[string]interface{}, error) {
	return d.detailsRepoInterface.Create(ctx, bundleDetailsDtoInput)
}
