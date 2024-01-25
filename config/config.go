package config

import (
	"github.com/joho/godotenv"
	"os"
	"testProject/internal/models"
)

type Config struct {
	Host string `env:"HOST"`
	Port string `env:"PORT"`
	DB   DB
	URL  models.Url
}

type DB struct {
	DriverName     string `env:"DRIVERNAME"`
	DataSourceName string `env:"DATASOURCENAME"`
}

func LoadConfig(path string) (Config, error) {
	if err := godotenv.Load(path); err != nil {
		return Config{}, err
	}

	cfg := Config{
		Host: os.Getenv("HOST"),
		Port: os.Getenv("PORT"),
		DB: DB{
			DriverName:     os.Getenv("DRIVERNAME"),
			DataSourceName: os.Getenv("DATASOURCENAME"),
		},
		URL: models.Url{
			AgeURL:         os.Getenv("AGEURL"),
			GenderURL:      os.Getenv("GENDERURL"),
			NationalityURL: os.Getenv("NATIONURL"),
		},
	}

	return cfg, nil
}
