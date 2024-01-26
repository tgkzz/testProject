package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

func (h Handler) GetPersonById(c *gin.Context) {
	id := c.Param("id")

	result, err := h.service.Person.GetUserById(id)
	if err != nil {
		if strings.Contains(err.Error(), "sql: no rows in result set") {
			h.errLogger.Printf("GetUserById %s", err.Error())
			ErrorHandler(c, fmt.Errorf("GetUserById %s", err.Error()), http.StatusNotFound)
			return
		}
		h.errLogger.Printf("GetUserById %s", err)
		ErrorHandler(c, err, http.StatusInternalServerError)
		return
	}

	var emptyPerson models.Person
	if result == emptyPerson {
		h.errLogger.Printf("GetUserById %s", "empty string")
		ErrorHandler(c, fmt.Errorf("GetUserById %s", "empty string"), http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h Handler) GetPersonByFilter(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil && c.Query("id") != "" {
		h.errLogger.Printf("Atoi : %s")
		ErrorHandler(c, err, http.StatusNotFound)
		return
	}
	ageTo, err := strconv.Atoi(c.Query("ageto"))
	if err != nil && c.Query("ageto") != "" {
		h.errLogger.Printf("Atoi : %s")
		ErrorHandler(c, err, http.StatusNotFound)
		return
	}
	ageFrom, err := strconv.Atoi(c.Query("agefrom"))
	if err != nil && c.Query("agefrom") != "" {
		h.errLogger.Printf("Atoi : %s")
		ErrorHandler(c, err, http.StatusNotFound)
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
		h.errLogger.Printf("GetUserByFilter: %s", err)
		if strings.Contains(err.Error(), "invalid filter") {
			ErrorHandler(c, err, http.StatusBadRequest)
		} else {
			ErrorHandler(c, err, http.StatusInternalServerError)
		}
		return
	}

	if len(result) == 0 {
		h.errLogger.Printf("GetUserByFilter: %s", fmt.Errorf("No records found"))
		ErrorHandler(c, fmt.Errorf("No records found"), http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, result)
}

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

func (h Handler) DeleteById(c *gin.Context) {
	id := c.Param("id")

	if err := h.service.Person.DeletePersonById(id); err != nil {
		h.errLogger.Printf("Delete person by id: %s", err)
		if strings.Contains(err.Error(), "strconv.Atoi") || strings.Contains(err.Error(), "invalid format") || strings.Contains(err.Error(), "no rows affected") {
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
