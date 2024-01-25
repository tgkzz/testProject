package service

import (
	"testProject/internal/models"
	"testProject/internal/repository"
	"testProject/internal/service/person"
)

type Service struct {
	Person person.IPersonService
	URLs   models.Url
}

func NewService(repo repository.Repository, url models.Url) *Service {
	return &Service{
		Person: person.NewPersonService(repo, url),
		URLs:   url,
	}
}
