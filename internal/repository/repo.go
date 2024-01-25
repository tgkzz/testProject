package repository

import (
	"database/sql"
	"testProject/internal/repository/person"
)

type Repository struct {
	person.IPersonRepo
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		IPersonRepo: person.NewPersonRepository(db),
	}
}
