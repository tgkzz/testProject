package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"testProject/internal/models"
)

func (h Handler) UpdateById(c *gin.Context) {
	id := c.Param("id")
	var person models.Person
	if err := c.BindJSON(&person); err != nil {
		h.errLogger.Printf("Bind JSON: %s", err)
		ErrorHandler(c, err, http.StatusInternalServerError)
		return
	}

	if err := h.service.Person.UpdateUserById(id, person); err != nil {
		h.errLogger.Printf("UpdateUserById: %s", err)
		if strings.Contains(err.Error(), "json may be empty or filled in incorrectly") {
			ErrorHandler(c, err, http.StatusBadRequest)
		} else {
			ErrorHandler(c, err, http.StatusInternalServerError)

		}
		return
	}

	h.infoLogger.Print("successfully finished update operation")
	response := map[string]string{
		"status":  "success",
		"message": "successfully update person",
	}
	c.JSON(http.StatusOK, response)
}
