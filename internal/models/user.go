package models

type Person struct {
	Id         int
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic,omitempty"`
	Age        int
	Gender     string
	CountryId  string
}
