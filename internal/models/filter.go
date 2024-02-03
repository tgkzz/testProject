package models

type Filter struct {
	Id          int
	Name        string
	Surname     string
	Patronymic  string
	AgeFrom     int
	AgeTo       int
	Gender      string
	Nationality string
	// pagination
	Limit  int
	Offset int
}
