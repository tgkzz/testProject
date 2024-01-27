package config

import (
	"fmt"
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
	DataSourceName string
}

func LoadConfig(path string) (Config, error) {
	if err := godotenv.Load(path); err != nil {
		return Config{}, err
	}

	dataSource := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DBDRIVER"), os.Getenv("DBUSER"),
		os.Getenv("DBPASS"), os.Getenv("DBHOST"),
		os.Getenv("DBPORT"), os.Getenv("DBNAME"),
	)

	cfg := Config{
		Host: os.Getenv("HOST"),
		Port: os.Getenv("PORT"),
		DB: DB{
			DriverName:     os.Getenv("DRIVERNAME"),
			DataSourceName: dataSource,
		},
		URL: models.Url{
			AgeURL:         os.Getenv("AGEURL"),
			GenderURL:      os.Getenv("GENDERURL"),
			NationalityURL: os.Getenv("NATIONURL"),
		},
	}

	return cfg, nil
}
