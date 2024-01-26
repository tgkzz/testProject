package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"testProject/internal/models"
)

func FetchData(baseURL, queryKey, queryValue string, target interface{}) error {
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
