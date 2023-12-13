package handler

import (
	"errors"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/laurati/projeto01/internal/entity"
	"github.com/laurati/projeto01/internal/usecase"
	"github.com/laurati/projeto01/internal/utils"
)

type DetailsHandler struct {
	detailsUseCaseInterface usecase.Usecase
}

func NewDetailsHandler(usecase *usecase.DetailsUseCase) *DetailsHandler {
	return &DetailsHandler{usecase}
}

func (h *DetailsHandler) CreateBundleDetailsHandler(c *gin.Context) {

	bundleDetails := entity.DetailsDtoInput{}

	err := c.ShouldBindJSON(&bundleDetails)
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, err.Error())
		return
	}

	if reflect.ValueOf(bundleDetails).IsZero() {
		utils.HandleError(c, http.StatusBadRequest, errors.New("bundle details can't be empty").Error())
		return
	}

	newbundle, err := h.detailsUseCaseInterface.CreateBundleDetails(c.Request.Context(), &bundleDetails)
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.HandleSucces(c, newbundle)

}
