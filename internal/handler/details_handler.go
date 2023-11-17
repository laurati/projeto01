package handler

import (
	"github.com/laurati/projeto01/internal/usecase"
)

type DetailsHandler struct {
	detailsUseCaseInterface usecase.Usecase
}

func NewDetailsHandler(usecase *usecase.DetailsUseCase) *DetailsHandler {
	return &DetailsHandler{usecase}
}
