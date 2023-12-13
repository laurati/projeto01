package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/laurati/projeto01/internal/handler"
)

func InitializeRouter(detailsHandler *handler.DetailsHandler) *gin.Engine {
	router := gin.Default()

	router.NoRoute(func(ctx *gin.Context) { ctx.JSON(http.StatusNotFound, gin.H{"message": "page not found"}) })
	router.GET("/", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, "up and running...") })

	bundleDetails := router.Group("/bundledetails")
	{
		bundleDetails.POST("/", detailsHandler.CreateBundleDetailsHandler)
		// bundleDetails.GET("/", handler.GetBundleDetails)
		// bundleDetails.GET("/:id", handler.GetBundleDetailsById)
	}

	return router
}
