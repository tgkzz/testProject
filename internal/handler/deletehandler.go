package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"testProject/internal/models"
)

func (h Handler) DeleteById(c *gin.Context) {
	id := c.Param("id")

	if err := h.service.Person.DeletePersonById(id); err != nil {
		if errors.Is(err, models.ErrBadStatusCode) || errors.Is(err, models.ErrNoRowsAffected) {
			ErrorHandler(c, err, http.StatusNotFound)
		} else {
			ErrorHandler(c, err, http.StatusInternalServerError)
		}
		return
	}

	SuccessHandler(c, models.SuccessDeleteOperation)
}
