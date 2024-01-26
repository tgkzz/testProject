package person

import (
	"database/sql"
	"testProject/internal/models"
)

type PersonRepo struct {
	DB *sql.DB
}

type IPersonRepo interface {
	Create(user models.Person) error
	GetUserById(id int) (models.Person, error)
	GetUserByFilter(filter models.Filter) ([]models.Person, error)
	DeleteById(id int) error
	UpdateUserById(id int, data models.Person) error
}

func NewPersonRepository(DB *sql.DB) *PersonRepo {
	return &PersonRepo{DB: DB}
}
