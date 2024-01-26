package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"testProject/internal/models"
)

func (h Handler) InsertNewPerson(c *gin.Context) {
	var person models.Person
	if err := c.BindJSON(&person); err != nil {
		h.errLogger.Printf("Bind JSON: %s", err)
		ErrorHandler(c, err, http.StatusBadRequest)
		return
	}

	if err := h.service.Person.CreateNewUser(person); err != nil {
		h.errLogger.Printf("Create new user: %s", err)
		if strings.Contains(err.Error(), "bad status code") || strings.Contains(err.Error(), "dial tcp:") {
			ErrorHandler(c, err, http.StatusFailedDependency)
		} else if strings.Contains(err.Error(), "name and surname must be not empty") {
			ErrorHandler(c, err, http.StatusBadRequest)
		} else {
			ErrorHandler(c, err, http.StatusInternalServerError)
		}
		return
	}

	h.infoLogger.Print("successfully finished create operation")
	response := map[string]string{
		"status":  "success",
		"message": "successfully inserted new person",
	}
	c.JSON(http.StatusOK, response)
}
