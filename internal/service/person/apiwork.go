package person

import (
	"encoding/json"
	"net/http"
	"net/url"
	"testProject/internal/models"
)

func (p PersonService) fetchData(baseURL, queryKey, queryValue string, target interface{}) error {
	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		return err
	}

	values := parsedURL.Query()
	values.Add(queryKey, queryValue)
	parsedURL.RawQuery = values.Encode()

	resp, err := http.Get(parsedURL.String())
	if err != nil {
		p.errLogger.Print(err)
		return models.ErrServiceUnavailable
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.ErrBadStatusCode
	}

	return json.NewDecoder(resp.Body).Decode(target)
}

func SelectNation(probabilities []models.NationalityProbability) string {
	if len(probabilities) == 0 {
		return ""
	}

	result := probabilities[0].CountryID

	maxProb := probabilities[0].Probability

	for _, probability := range probabilities {
		if probability.Probability > maxProb {
			result = probability.CountryID
		}
	}

	return result
}
