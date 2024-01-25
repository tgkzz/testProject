package handler

import (
	"log"
	"testProject/internal/service"
)

type Handler struct {
	service    *service.Service
	infoLogger *log.Logger
	errLogger  *log.Logger
}

func NewHandler(service *service.Service, infoLog, errLog *log.Logger) *Handler {
	return &Handler{
		service:    service,
		infoLogger: infoLog,
		errLogger:  errLog,
	}
}
