package handler

import "github.com/gin-gonic/gin"

func ErrorHandler(c *gin.Context, err error, code int) {
	response := map[string]string{
		"status": "error",
		"desc":   err.Error(),
	}

	c.JSON(code, response)
}
