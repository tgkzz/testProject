package person

import (
	"fmt"
	"strings"
	"testProject/internal/models"
)

// CREATE
func (p PersonRepo) Create(user models.Person) error {
	query := "INSERT INTO person (name, surname, patronymic, age, gender, countryid) values ($1, $2, $3, $4, $5, $6)"

	if _, err := p.DB.Exec(query, user.Name, user.Surname, user.Patronymic, user.Age, user.Gender, user.CountryId); err != nil {
		return err
	}

	return nil
}

// READ

func (p PersonRepo) GetUserById(id int) (models.Person, error) {
	var result models.Person

	query := "SELECT id, name, surname, patronymic, age, gender, countryid FROM person WHERE id = $1"

	if err := p.DB.QueryRow(query, id).Scan(&result.Id, &result.Name, &result.Surname, &result.Patronymic, &result.Age, &result.Gender, &result.CountryId); err != nil {
		return models.Person{}, err
	}

	return result, nil
}

func (p PersonRepo) GetUserByFilter(filter models.Filter) ([]models.Person, error) {
	var result []models.Person
	var params []interface{}

	query := "SELECT id, name, surname, patronymic, age, gender, countryid FROM person WHERE 1=1"

	paramID := 1
	if filter.Id != 0 {
		query += fmt.Sprintf(" AND id = $%d", paramID)
		params = append(params, filter.Id)
		paramID++
	}

	if filter.Name != "" {
		query += fmt.Sprintf(" AND name = $%d", paramID)
		params = append(params, filter.Name)
		paramID++
	}
	if filter.Surname != "" {
		query += fmt.Sprintf(" AND surname = $%d", paramID)
		params = append(params, filter.Surname)
		paramID++
	}
	if filter.Patronymic != "" {
		query += fmt.Sprintf(" AND patronymic = $%d", paramID)
		params = append(params, filter.Patronymic)
		paramID++
	}
	if filter.AgeFrom != 0 {
		query += fmt.Sprintf(" AND age >= $%d", paramID)
		params = append(params, filter.AgeFrom)
		paramID++
	}

	if filter.AgeTo != 0 {
		query += fmt.Sprintf(" AND age <= $%d", paramID)
		params = append(params, filter.AgeTo)
		paramID++
	}
	if filter.Gender != "" {
		query += fmt.Sprintf(" AND gender = $%d", paramID)
		params = append(params, filter.Gender)
		paramID++
	}
	if filter.Nationality != "" {
		query += fmt.Sprintf(" AND countryid = $%d", paramID)
		params = append(params, filter.Nationality)
		paramID++
	}

	rows, err := p.DB.Query(query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var person models.Person
		if err := rows.Scan(&person.Id, &person.Name, &person.Surname, &person.Patronymic, &person.Age, &person.Gender, &person.CountryId); err != nil {
			return nil, err
		}
		result = append(result, person)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, err
}

// UPDATE

func (p PersonRepo) UpdateUserById(id int, data models.Person) error {
	query := "UPDATE person SET "

	var params []interface{}
	paramID := 1

	if data.Name != "" {
		query += fmt.Sprintf("name = $%d, ", paramID)
		params = append(params, data.Name)
		paramID++
	}

	if data.Surname != "" {
		query += fmt.Sprintf("surname = $%d, ", paramID)
		params = append(params, data.Surname)
		paramID++
	}
	if data.Patronymic != "" {
		query += fmt.Sprintf("patronymic = $%d, ", paramID)
		params = append(params, data.Patronymic)
		paramID++
	}
	if data.Age != 0 {
		query += fmt.Sprintf("age = $%d, ", paramID)
		params = append(params, data.Age)
		paramID++
	}
	if data.Gender != "" {
		query += fmt.Sprintf("gender = $%d, ", paramID)
		params = append(params, data.Gender)
		paramID++
	}
	if data.CountryId != "" {
		query += fmt.Sprintf("countryid = $%d, ", paramID)
		params = append(params, data.CountryId)
		paramID++
	}
	query = strings.TrimSuffix(query, ", ")

	query += fmt.Sprintf(" WHERE id = $%d;", paramID)
	params = append(params, id)

	_, err := p.DB.Exec(query, params...)

	return err
}

// DELETE

func (p PersonRepo) DeleteById(id int) error {
	query := "DELETE FROM person WHERE id = $1"

	res, err := p.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("no rows affected, person with id %d does not exist", id)
	}

	return nil
}
