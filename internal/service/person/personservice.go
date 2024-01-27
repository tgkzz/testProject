package person

import (
	"log"
	"testProject/internal/models"
	"testProject/internal/repository/person"
)

type PersonService struct {
	repo       person.IPersonRepo
	URLs       models.Url
	infoLogger *log.Logger
	errLogger  *log.Logger
}

type IPersonService interface {
	CreateNewUser(person models.Person) error
	DeletePersonById(id string) error
	GetUserByFilter(filter models.Filter) ([]models.Person, error)
	GetUserById(id string) (models.Person, error)
	UpdateUserById(id string, person models.Person) error
}

func NewPersonService(repo person.IPersonRepo, url models.Url, infoLogger, errLogger *log.Logger) *PersonService {
	return &PersonService{
		repo:       repo,
		URLs:       url,
		infoLogger: infoLogger,
		errLogger:  errLogger,
	}
}
