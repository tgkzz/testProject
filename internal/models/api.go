package models

type Url struct {
	AgeURL         string
	NationalityURL string
	GenderURL      string
}

type Age struct {
	Count int    `json:"count"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
}

type Gender struct {
	Count       int     `json:"count"`
	Name        string  `json:"name"`
	Gender      string  `json:"gender"`
	Probability float32 `json:"probability"`
}

type NationalityProbability struct {
	CountryID   string  `json:"country_id"`
	Probability float64 `json:"probability"`
}

type Nationality struct {
	Count  int                      `json:"count"`
	Name   string                   `json:"name"`
	Nation []NationalityProbability `json:"country"`
}
