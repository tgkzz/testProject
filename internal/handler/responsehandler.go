package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testProject/internal/models"
)

func ErrorHandler(c *gin.Context, err error, code int) {
	response := models.Response{
		Status:      models.ErrorMsg,
		Description: err.Error(),
	}

	c.JSON(code, response)
}

func SuccessHandler(c *gin.Context, operation string) {
	response := models.Response{
		Status:      models.SuccessMsg,
		Description: operation,
	}

	c.JSON(http.StatusOK, response)
}
