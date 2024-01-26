package handler

import "github.com/gin-gonic/gin"

func (h Handler) Routes() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")

	api.GET("/get", h.GetPersonByFilter)

	api.GET("/get/:id", h.GetPersonById)

	api.POST("/create", h.InsertNewPerson)

	api.PATCH("/update/:id", h.UpdateById)

	api.DELETE("/delete/:id", h.DeleteById)

	return r
}
