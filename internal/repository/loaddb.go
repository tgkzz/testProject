package repository

import (
	"database/sql"
	_ "github.com/lib/pq"
	"os"
	"strings"
)

const dbpath = "./database/init.sql"

func LoadDB(driverName, dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return db, err
	}

	if err := db.Ping(); err != nil {
		return db, err
	}

	if err := createTable(db); err != nil {
		return db, err
	}

	return db, err
}

func createTable(db *sql.DB) error {
	fileSql, err := os.ReadFile(dbpath)
	if err != nil {
		return err
	}

	requests := strings.Split(string(fileSql), ";")
	for _, request := range requests {
		_, err = db.Exec(request)
		if err != nil {
			return err
		}
	}

	return nil
}
