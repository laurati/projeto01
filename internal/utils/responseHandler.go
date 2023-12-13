package utils

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func HandleSucces(c *gin.Context, data interface{}) {
	responData := Response{
		Status:  "200",
		Message: "Success",
		Data:    data,
	}
	c.JSON(http.StatusOK, responData)
}

func HandleError(c *gin.Context, status int, message string) {
	responData := Response{
		Status:  strconv.Itoa(status),
		Message: message,
	}
	c.JSON(status, responData)
}
