package person

import (
	"fmt"
	"sync"
	"testProject/internal/models"
	"testProject/internal/pkg/api"
	"testProject/internal/pkg/other"
	"testProject/internal/pkg/validation"
)

func (p PersonService) DeletePersonById(id string) error {
	parsedId, err := other.StrictAtoi(id)
	if err != nil {
		return err
	}

	return p.repo.DeleteById(parsedId)
}

func (p PersonService) GetUserByFilter(filter models.Filter) ([]models.Person, error) {
	if !validation.IsValidFilter(filter) {
		return []models.Person{}, fmt.Errorf("invalid filter")
	}

	return p.repo.GetUserByFilter(filter)
}

func (p PersonService) GetUserById(id string) (models.Person, error) {
	parsedId, err := other.StrictAtoi(id)
	if err != nil {
		return models.Person{}, err
	}
	return p.repo.GetUserById(parsedId)
}

func (p PersonService) UpdateUserById(id string, data models.Person) error {
	parsedId, err := other.StrictAtoi(id)
	if err != nil {
		return err
	}

	if !validation.IsValidUpdateParams(data) {
		return fmt.Errorf("json may be empty or filled in incorrectly")
	}

	return p.repo.UpdateUserById(parsedId, data)
}

func (p PersonService) CreateNewUser(person models.Person) error {
	var wg sync.WaitGroup
	var age models.Age
	var gender models.Gender
	var nation models.Nationality

	if !validation.IsValidData(person) {
		return fmt.Errorf("name and surname must be not empty")
	}

	errCh := make(chan error, 3)

	wg.Add(3)

	go func() {
		defer wg.Done()
		if err := api.FetchData(p.URLs.AgeURL, "name", person.Name, &age); err != nil {
			errCh <- err
			return
		}
		person.Age = age.Age
		errCh <- nil
	}()

	go func() {
		defer wg.Done()
		if err := api.FetchData(p.URLs.GenderURL, "name", person.Name, &gender); err != nil {
			errCh <- err
			return
		}
		person.Gender = gender.Gender
		errCh <- nil
	}()

	go func() {
		defer wg.Done()

		if err := api.FetchData(p.URLs.NationalityURL, "name", person.Surname, &nation); err != nil {
			errCh <- err
			return
		}

		var err error
		person.CountryId, err = api.SelectNation(nation.Nation)
		if err != nil {
			errCh <- err
			return
		}
		errCh <- nil
	}()

	wg.Wait()
	close(errCh)

	for err := range errCh {
		if err != nil {
			return err
		}
	}

	if err := p.repo.Create(person); err != nil {
		return err
	}

	return nil
}
