package person

import (
	"sync"
	"testProject/internal/models"
	"testProject/internal/pkg"
)

func (p PersonService) DeletePersonById(id string) error {
	parsedId, err := pkg.StrictAtoi(id)
	if err != nil {
		p.errLogger.Print(err)
		return models.ErrAtoi
	}

	if err := p.repo.DeleteById(parsedId); err != nil {
		p.errLogger.Print(err)
		return err
	}

	p.infoLogger.Print(models.SuccessDeleteOperation)
	return nil
}

func (p PersonService) GetUserByFilter(filter models.Filter) ([]models.Person, error) {
	if !pkg.IsValidFilter(filter) {
		p.errLogger.Printf(models.ErrInvalidFilter.Error())
		return []models.Person{}, models.ErrInvalidFilter
	}

	if !pkg.IsValidPagination(filter) {
		p.errLogger.Printf(models.ErrInvalidPagination.Error())
		return []models.Person{}, models.ErrInvalidPagination
	}

	result, err := p.repo.GetUserByFilter(filter)
	if err != nil {
		p.errLogger.Print(err)
		return nil, err
	}

	if len(result) == 0 {
		p.errLogger.Print(models.ErrSqlNoRows.Error())
		return nil, models.ErrSqlNoRows
	}

	p.infoLogger.Print(models.SuccessGetOperation)
	return result, nil
}

func (p PersonService) GetUserById(id string) (models.Person, error) {
	parsedId, err := pkg.StrictAtoi(id)
	if err != nil {
		p.errLogger.Print(err)
		return models.Person{}, models.ErrAtoi
	}

	result, err := p.repo.GetUserById(parsedId)
	if err != nil {
		p.errLogger.Print(err)
		return models.Person{}, err
	}

	var emptyPerson models.Person
	if result == emptyPerson {
		p.errLogger.Print(models.ErrEmptyResult)
		return models.Person{}, models.ErrEmptyResult
	}

	p.infoLogger.Print(models.SuccessGetOperation)
	return result, nil
}

func (p PersonService) UpdateUserById(id string, data models.Person) error {
	parsedId, err := pkg.StrictAtoi(id)
	if err != nil {
		p.errLogger.Print(err)
		return models.ErrAtoi
	}

	if !pkg.IsValidUpdateParams(data) {
		p.errLogger.Print(models.ErrInvalidUpdateParams.Error())
		return models.ErrInvalidUpdateParams
	}

	if _, err := p.repo.GetUserById(parsedId); err != nil {
		p.errLogger.Print(err)
		return err
	}

	if err := p.repo.UpdateUserById(parsedId, data); err != nil {
		p.errLogger.Print(err)
		return err
	}

	p.infoLogger.Print(models.SuccessPatchOperation)
	return nil
}

func (p PersonService) CreateNewUser(person models.Person) error {
	var wg sync.WaitGroup
	var age models.Age
	var gender models.Gender
	var nation models.Nationality

	if !pkg.IsValidData(person) {
		p.errLogger.Print(models.ErrEmptyNameOrSurname.Error())
		return models.ErrEmptyNameOrSurname
	}

	errCh := make(chan error, 3)

	wg.Add(3)

	go func() {
		defer wg.Done()
		if err := p.fetchData(p.URLs.AgeURL, "name", person.Name, &age); err != nil {
			errCh <- err
			return
		}
		person.Age = age.Age
		errCh <- nil
	}()

	go func() {
		defer wg.Done()
		if err := p.fetchData(p.URLs.GenderURL, "name", person.Name, &gender); err != nil {
			errCh <- err
			return
		}
		person.Gender = gender.Gender
		errCh <- nil
	}()

	go func() {
		defer wg.Done()

		if err := p.fetchData(p.URLs.NationalityURL, "name", person.Surname, &nation); err != nil {
			errCh <- err
			return
		}

		person.CountryId = SelectNation(nation.Nation)
		errCh <- nil
	}()

	wg.Wait()
	close(errCh)

	for err := range errCh {
		if err != nil {
			p.errLogger.Print(err)
			return err
		}
	}

	if err := p.repo.Create(person); err != nil {
		p.errLogger.Print(err)
		return err
	}

	p.infoLogger.Print(models.SuccessCreatedOperation)
	return nil
}
