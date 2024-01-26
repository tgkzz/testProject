package handler

import (
	"github.com/gin-gonic/gin"
	"log"
)

func (h Handler) Routes(infoLog *log.Logger, errLog *log.Logger) *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery())

	api := r.Group("/api")

	api.Use(h.GinLogger(infoLog, errLog))

	api.GET("/get", h.GetPersonByFilter)

	api.GET("/get/:id", h.GetPersonById)

	api.POST("/create", h.InsertNewPerson)

	api.PATCH("/update/:id", h.UpdateById)

	api.DELETE("/delete/:id", h.DeleteById)

	return r
}
