package handler

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"

	"github.com/thatstoasty/health-tracker/queries"
)

// Get exercise details
func GetExerciseDetails(c echo.Context) error {
	connectionString := getConnectionString()
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
		log.Fatal("failed to establish connection to postgres")
		return c.String(http.StatusBadRequest, "failed to establish connection to postgres")
	}

	queries := queries.New(db)
	ctx := context.Background()
	exercise := c.Param("exercise")

	composition, err := queries.GetExerciseDetails(ctx, exercise)
	if err != nil {
		log.Fatal(err)
		log.Fatal("failed to get exercise details")
		return c.String(http.StatusBadRequest, "failed to get exercise details")
	}

	return c.JSON(http.StatusOK, composition)
}
