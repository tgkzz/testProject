package handler

import "github.com/gin-gonic/gin"

func (h Handler) Routes() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")

	api.POST("/create", h.InsertNewPerson)

	api.DELETE("/delete/:id", h.DeleteById)

	api.GET("/get", h.GetPersonByFilter)

	api.GET("/get/:id", h.GetPersonById)

	return r
}
