package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/laurati/projeto01/internal/handler"
)

func InitializeRouter(detailsHandler *handler.DetailsHandler) *gin.Engine {
	router := gin.Default()

	// router.NoRoute(func(ctx *gin.Context) { ctx.JSON(http.StatusNotFound, gin.H{"message": "page not found"}) })
	router.GET("/", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, "up and running...") })

	// transactions := router.Group("/transactions")
	// transactions.POST("/create", transactionHandler.Create)
	// transactions.GET("", transactionHandler.GetTransactionsCurrency)
	// transactions.GET("/:id", transactionHandler.GetTransactionByIdCurrency)
	return router
}
