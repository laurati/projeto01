package router

import (
	"github.com/gin-gonic/gin"
	"github.com/laurati/projeto01/internal/handler"
)

func InitializeRouter(detailsHandler *handler.DetailsHandler) *gin.Engine {
	router := gin.Default()
	// TODO:Lógica de inicialização do roteador utilizando o transactionHandler
	// ...

	return router
}
