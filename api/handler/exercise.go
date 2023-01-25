package handler

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"

	"github.com/thatstoasty/health-tracker/queries"
	"github.com/thatstoasty/health-tracker/utils"
)

// Get exercise details
func GetExercise(c echo.Context) error {
	connectionString := utils.GetConnectionString()
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
		log.Fatal("failed to establish connection to postgres")
		return c.String(http.StatusBadRequest, "failed to establish connection to postgres")
	}

	queries := queries.New(db)
	ctx := context.Background()
	exercise := c.Param("exercise")

	composition, err := queries.GetExercise(ctx, exercise)
	if err != nil {
		log.Fatal(err)
		log.Fatal("failed to get exercise details")
		return c.String(http.StatusBadRequest, "failed to get exercise details")
	}

	return c.JSON(http.StatusOK, composition)
}

// Get exercise names
func GetExerciseNames(c echo.Context) error {
	connectionString := utils.GetConnectionString()
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
		log.Fatal("failed to establish connection to postgres")
		return c.String(http.StatusBadRequest, "failed to establish connection to postgres")
	}

	queries := queries.New(db)
	ctx := context.Background()
	limitString := c.QueryParam("limit")

	limit, err := strconv.Atoi(limitString)
    if err != nil {
        log.Fatal(err)
		log.Fatal("failed to convert to int")
		return c.String(http.StatusBadRequest, "failed to convert to int")
    }

	exerciseNames, err := queries.GetExercises(ctx, int32(limit))
	if err != nil {
		log.Fatal(err)
		log.Fatal("failed to get exercise names")
		return c.String(http.StatusBadRequest, "failed to get exercise names")
	}

	return c.JSON(http.StatusOK, exerciseNames)
}

// Delete program
func DeleteExercise(c echo.Context) error {
	connectionString := utils.GetConnectionString()
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
		log.Fatal("failed to establish connection to postgres")
		return c.String(http.StatusBadRequest, "failed to establish connection to postgres")
	}

	queries := queries.New(db)
	ctx := context.Background()
	name := c.Param("name")

	error := queries.DeleteExercise(ctx, name)
	if error != nil {
		log.Fatal(err)
		log.Fatal("failed to delete exercise")
		return c.String(http.StatusBadRequest, "failed to delete exercise")
	}

	return c.String(http.StatusOK, "exercise deleted.")
}
