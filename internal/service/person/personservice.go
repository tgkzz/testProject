package person

import (
	"testProject/internal/models"
	"testProject/internal/repository/person"
)

type PersonService struct {
	repo person.IPersonRepo
	URLs models.Url
}

type IPersonService interface {
	CreateNewUser(person models.Person) error
	DeletePersonById(id string) error
	GetUserByFilter(filter models.Filter) ([]models.Person, error)
	GetUserById(id string) (models.Person, error)
	UpdateUserById(id string, person models.Person) error
}

func NewPersonService(repo person.IPersonRepo, url models.Url) *PersonService {
	return &PersonService{
		repo: repo,
		URLs: url,
	}
}
