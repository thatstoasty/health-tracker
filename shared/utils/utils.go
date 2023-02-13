package utils

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"

	"github.com/thatstoasty/health-tracker/shared/models"
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

func GetQueryInterface() (*models.Queries, error) {
	connectionString := getConnectionString()
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	queries := models.New(db)

	return queries, nil
}

func GetRIRMapping() map[float64]int {
	// Generate mapping of RIR:Array Index
	rirMapping := make(map[float64]int)
	for i, j := 0.0, 0; i < 5.5; i, j = i + 0.5, j + 1{
		rirMapping[i] = j
	}

	return rirMapping
}

func GetWeightTable() [][]float64 {
	// x: 0 -> 4 RIR
	// y: 1 -> 10 Reps
	table := [][]float64 {
		{1.0, 0.978, 0.955, 0.939, 0.922, 0.907, 0.892, 0.878, 0.863, 0.85, 0.837},
		{0.955, 0.939, 0.922, 0.907, 0.892, 0.878, 0.863, 0.85, 0.837, 0.824, 0.811},
		{0.922, 0.907, 0.892, 0.878, 0.863, 0.85, 0.837, 0.824, 0.811, 0.799, 0.786},
		{0.892, 0.878, 0.863, 0.85, 0.837, 0.824, 0.811, 0.799, 0.786, 0.774, 0.762},
		{0.863, 0.85, 0.837, 0.824, 0.811, 0.799, 0.786, 0.774, 0.762, 0.751, 0.739},
		{0.837, 0.824, 0.811, 0.799, 0.786, 0.774, 0.762, 0.751, 0.739, 0.723, 0.707},
		{0.811, 0.799, 0.786, 0.774, 0.762, 0.751, 0.739, 0.723, 0.707, 0.694, 0.68},
		{0.786, 0.774, 0.762, 0.751, 0.739, 0.723, 0.707, 0.694, 0.68, 0.667, 0.653},
		{0.762, 0.751, 0.739, 0.723, 0.707, 0.694, 0.68, 0.667, 0.653, 0.64, 0.626},
		{0.739, 0.723, 0.707, 0.694, 0.68, 0.667, 0.653, 0.64, 0.626, 0.613, 0.599},
	}

	return table
}