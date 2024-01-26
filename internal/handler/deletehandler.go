package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (h Handler) DeleteById(c *gin.Context) {
	id := c.Param("id")

	if err := h.service.Person.DeletePersonById(id); err != nil {
		h.errLogger.Printf("Delete person by id: %s", err)
		if strings.Contains(err.Error(), "strconv.Atoi") || strings.Contains(err.Error(), "no rows affected") {
			ErrorHandler(c, err, http.StatusNotFound)
		} else {
			ErrorHandler(c, err, http.StatusInternalServerError)
		}
		return
	}

	h.infoLogger.Print("successfully finished deletion operation")
	response := map[string]string{
		"status":  "success",
		"message": "successfully deleted person",
	}
	c.JSON(http.StatusOK, response)
}
