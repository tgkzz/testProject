package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"testProject/internal/models"
)

func (h Handler) InsertNewPerson(c *gin.Context) {
	var person models.Person
	if err := c.BindJSON(&person); err != nil {
		ErrorHandler(c, err, http.StatusBadRequest)
		return
	}

	if err := h.service.Person.CreateNewUser(person); err != nil {
		if errors.Is(err, models.ErrBadStatusCode) || errors.Is(err, models.ErrServiceUnavailable) {
			ErrorHandler(c, err, http.StatusFailedDependency)
		} else if errors.Is(err, models.ErrEmptyNameOrSurname) {
			ErrorHandler(c, err, http.StatusBadRequest)
		} else {
			ErrorHandler(c, err, http.StatusInternalServerError)
		}
		return
	}

	SuccessHandler(c, models.SuccessCreatedOperation)
}
