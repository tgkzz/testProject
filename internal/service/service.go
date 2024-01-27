package service

import (
	"log"
	"testProject/internal/models"
	"testProject/internal/repository"
	"testProject/internal/service/person"
)

type Service struct {
	Person     person.IPersonService
	URLs       models.Url
	infoLogger *log.Logger
	errLogger  *log.Logger
}

func NewService(repo repository.Repository, url models.Url, infoLogger, errLogger *log.Logger) *Service {
	return &Service{
		Person: person.NewPersonService(repo, url, infoLogger, errLogger),
		URLs:   url,
	}
}
