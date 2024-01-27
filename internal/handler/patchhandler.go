package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"testProject/internal/models"
)

func (h Handler) UpdateById(c *gin.Context) {
	id := c.Param("id")
	var person models.Person
	if err := c.BindJSON(&person); err != nil {
		ErrorHandler(c, err, http.StatusInternalServerError)
		return
	}

	if err := h.service.Person.UpdateUserById(id, person); err != nil {
		if errors.Is(err, models.ErrInvalidUpdateParams) {
			ErrorHandler(c, err, http.StatusBadRequest)
		} else if errors.Is(err, models.ErrSqlNoRows) || errors.Is(err, models.ErrAtoi) {
			ErrorHandler(c, err, http.StatusNotFound)
		} else {
			ErrorHandler(c, err, http.StatusInternalServerError)
		}
		return
	}

	SuccessHandler(c, models.SuccessPatchOperation)
}
