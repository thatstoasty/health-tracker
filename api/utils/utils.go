package utils

import (
	"fmt"
	"database/sql"

	"github.com/labstack/echo/v4"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "docker"
	password = "docker"
	dbname   = "docker"
	sslmode  = "disable"
)

func getConnectionString() string {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
    host, port, user, password, dbname, sslmode)

	return connectionString
}

func GetDBConnection(c echo.Context) (*sql.DB, error) {
	connectionString := getConnectionString()
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	return db, nil
}
