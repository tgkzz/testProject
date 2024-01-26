package person

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"testProject/internal/models"
)

func StrictAtoi(s string) (int, error) {
	if strings.TrimLeft(s, "0") != s || strings.Contains(s, "+") {
		return 0, errors.New("invalid format: leading zeros or plus signs are not allowed")
	}
	if num, err := strconv.Atoi(s); err == nil {
		return num, nil
	} else {
		return 0, err
	}
}

func isValidFilter(filter models.Filter) bool {
	if filter.AgeTo <= 0 || filter.AgeFrom <= 0 || filter.Id <= 0 {
		return false
	}

	return true
}

func isValidUpdateParams(data models.Person) bool {
	return data != models.Person{}
}

func fetchData(baseURL, queryKey, queryValue string, target interface{}) error {
	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		return err
	}

	values := parsedURL.Query()
	values.Add(queryKey, queryValue)
	parsedURL.RawQuery = values.Encode()

	resp, err := http.Get(parsedURL.String())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status code: %d", resp.StatusCode)
	}

	return json.NewDecoder(resp.Body).Decode(target)
}

func SelectNation(probabilities []models.NationalityProbability) (string, error) {
	if len(probabilities) == 0 {
		return "", fmt.Errorf("surname cannot be empty")
	}

	result := probabilities[0].CountryID

	maxProb := probabilities[0].Probability

	for _, probability := range probabilities {
		if probability.Probability > maxProb {
			result = probability.CountryID
		}
	}

	return result, nil
}

func dataValidation(person models.Person) bool {
	return len(strings.TrimSpace(person.Name)) != 0 && len(strings.TrimSpace(person.Surname)) != 0
}

func (p PersonService) DeletePersonById(id string) error {
	parsedId, err := StrictAtoi(id)
	if err != nil {
		return err
	}

	return p.repo.DeleteById(parsedId)
}

func (p PersonService) GetUserByFilter(filter models.Filter) ([]models.Person, error) {
	if !isValidFilter(filter) {
		return []models.Person{}, fmt.Errorf("invalid filter")
	}

	return p.repo.GetUserByFilter(filter)
}

func (p PersonService) GetUserById(id string) (models.Person, error) {
	parsedId, err := StrictAtoi(id)
	if err != nil {
		return models.Person{}, err
	}
	return p.repo.GetUserById(parsedId)
}

func (p PersonService) UpdateUserById(id string, data models.Person) error {
	parsedId, err := StrictAtoi(id)
	if err != nil {
		return err
	}

	if !isValidUpdateParams(data) {
		return fmt.Errorf("json may be empty or filled in incorrectly")
	}

	return p.repo.UpdateUserById(parsedId, data)
}

func (p PersonService) CreateNewUser(person models.Person) error {
	var wg sync.WaitGroup
	var age models.Age
	var gender models.Gender
	var nation models.Nationality

	if !dataValidation(person) {
		return fmt.Errorf("name and surname must be not empty")
	}

	errCh := make(chan error, 3)

	wg.Add(3)

	go func() {
		defer wg.Done()
		if err := fetchData(p.URLs.AgeURL, "name", person.Name, &age); err != nil {
			errCh <- err
			return
		}
		person.Age = age.Age
		errCh <- nil
	}()

	go func() {
		defer wg.Done()
		if err := fetchData(p.URLs.GenderURL, "name", person.Name, &gender); err != nil {
			errCh <- err
			return
		}
		person.Gender = gender.Gender
		errCh <- nil
	}()

	go func() {
		defer wg.Done()

		if err := fetchData(p.URLs.NationalityURL, "name", person.Surname, &nation); err != nil {
			errCh <- err
			return
		}

		var err error
		person.CountryId, err = SelectNation(nation.Nation)
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
