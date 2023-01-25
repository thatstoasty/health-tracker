package utils

import (
	"fmt"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "docker"
	password = "docker"
	dbname   = "docker"
	sslmode  = "disable"
)

func GetConnectionString() string {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
    host, port, user, password, dbname, sslmode)

	return connectionString
}
