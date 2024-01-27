package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"testProject/internal/models"
	"testProject/internal/pkg"
)

func (h Handler) GetPersonByFilter(c *gin.Context) {
	id, err := pkg.StrictAtoi(c.Query("id"))
	if err != nil && c.Query("id") != "" {
		ErrorHandler(c, models.ErrAtoi, http.StatusNotFound)
		return
	}
	ageTo, err := pkg.StrictAtoi(c.Query("ageto"))
	if err != nil && c.Query("ageto") != "" {
		ErrorHandler(c, models.ErrAtoi, http.StatusNotFound)
		return
	}
	ageFrom, err := pkg.StrictAtoi(c.Query("agefrom"))
	if err != nil && c.Query("agefrom") != "" {
		ErrorHandler(c, models.ErrAtoi, http.StatusNotFound)
		return
	}

	filter := models.Filter{
		Id:          id,
		Name:        c.Query("name"),
		Surname:     c.Query("surname"),
		Patronymic:  c.Query("patronymic"),
		AgeTo:       ageTo,
		AgeFrom:     ageFrom,
		Gender:      c.Query("gender"),
		Nationality: c.Query("nation"),
	}

	result, err := h.service.Person.GetUserByFilter(filter)
	if err != nil {
		if errors.Is(err, models.ErrInvalidFilter) {
			ErrorHandler(c, err, http.StatusBadRequest)
		} else if errors.Is(err, models.ErrSqlNoRows) {
			ErrorHandler(c, err, http.StatusNotFound)
		} else {
			ErrorHandler(c, err, http.StatusInternalServerError)
		}
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h Handler) GetPersonById(c *gin.Context) {
	id := c.Param("id")

	result, err := h.service.Person.GetUserById(id)
	if err != nil {
		if errors.Is(err, models.ErrSqlNoRows) || errors.Is(err, models.ErrAtoi) || errors.Is(err, models.ErrEmptyResult) {
			ErrorHandler(c, err, http.StatusNotFound)
			return
		}
		ErrorHandler(c, err, http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, result)
}
