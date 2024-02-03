package pkg

import (
	"strings"
	"testProject/internal/models"
)

func IsValidData(person models.Person) bool {
	return len(strings.TrimSpace(person.Name)) != 0 && len(strings.TrimSpace(person.Surname)) != 0
}

func IsValidUpdateParams(data models.Person) bool {
	return data != models.Person{}
}

func IsValidFilter(filter models.Filter) bool {
	if filter.AgeTo < 0 || filter.AgeFrom < 0 || filter.Id < 0 {
		return false
	}

	if filter.AgeFrom > filter.AgeTo {
		return false
	}

	return true
}

func IsValidPagination(filter models.Filter) bool {
	if filter.Limit < 0 || filter.Offset < 0 {
		return false
	}

	return true
}
