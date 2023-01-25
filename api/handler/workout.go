package handler

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"

	"github.com/thatstoasty/health-tracker/queries"
	"github.com/thatstoasty/health-tracker/utils"
)

// Get workout
func GetWorkout(c echo.Context) error {
	connectionString := utils.GetConnectionString()
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Println(err)
		log.Println("failed to establish connection to postgres")
		return c.String(http.StatusBadRequest, "failed to establish connection to postgres")
	}

	queries := queries.New(db)
	ctx := context.Background()
	name := c.Param("name")

	composition, err := queries.GetWorkout(ctx, name)
	if err != nil {
		log.Println(err)
		log.Println("failed to get workout")
		return c.String(http.StatusBadRequest, "failed to get workout")
	}

	return c.JSON(http.StatusOK, composition)
}

// Get workout performed
func GetWorkoutPerformed(c echo.Context) error {
	connectionString := utils.GetConnectionString()
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Println(err)
		log.Println("failed to establish connection to postgres")
		return c.String(http.StatusBadRequest, "failed to establish connection to postgres")
	}

	queries := queries.New(db)
	ctx := context.Background()
	dateString := c.Param("date")
	date, err := time.Parse("YYYY-MM-DD", dateString)
	if err != nil {
		log.Println(err)
		log.Println("failed to parse date")
		return c.String(http.StatusBadRequest, "failed to parse date")
	}

	workoutPerformed, err := queries.GetWorkoutPerformed(ctx, date)
	if err != nil {
		log.Println(err)
		log.Println("failed to get workout")
		return c.String(http.StatusBadRequest, "failed to get workout")
	}

	return c.JSON(http.StatusOK, workoutPerformed)
}

// Get workout names
func GetWorkoutNames(c echo.Context) error {
	connectionString := utils.GetConnectionString()
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Println(err)
		log.Println("failed to establish connection to postgres")
		return c.String(http.StatusBadRequest, "failed to establish connection to postgres")
	}

	queries := queries.New(db)
	ctx := context.Background()
	limitString := c.QueryParam("limit")

	limit, err := strconv.Atoi(limitString)
	if err != nil {
		log.Println(err)
		log.Println("failed to convert to int")
		return c.String(http.StatusBadRequest, "failed to convert to int")
	}

	workoutNames, err := queries.GetWorkoutNames(ctx, int32(limit))
	if err != nil {
		log.Println(err)
		log.Println("failed to get workout")
		return c.String(http.StatusBadRequest, "failed to get workout")
	}

	return c.JSON(http.StatusOK, workoutNames)
}

// Delete workout
func DeleteWorkout(c echo.Context) error {
	connectionString := utils.GetConnectionString()
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Println(err)
		log.Println("failed to establish connection to postgres")
		return c.String(http.StatusBadRequest, "failed to establish connection to postgres")
	}

	queries := queries.New(db)
	ctx := context.Background()
	name := c.Param("name")

	error := queries.DeleteWorkout(ctx, name)
	if error != nil {
		log.Println(err)
		log.Println("failed to delete workout")
		return c.String(http.StatusBadRequest, "failed to delete workout")
	}

	return c.String(http.StatusOK, "workout deleted.")
}

// Delete workout
func DeleteWorkoutPerformed(c echo.Context) error {
	connectionString := utils.GetConnectionString()
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Println(err)
		log.Println("failed to establish connection to postgres")
		return c.String(http.StatusBadRequest, "failed to establish connection to postgres")
	}

	queries := queries.New(db)
	ctx := context.Background()
	dateString := c.Param("date")
	date, err := time.Parse("YYYY-MM-DD", dateString)
	if err != nil {
		log.Println(err)
		log.Println("failed to parse date")
		return c.String(http.StatusBadRequest, "failed to parse date")
	}

	error := queries.DeleteWorkoutPerformed(ctx, date)
	if error != nil {
		log.Println(err)
		log.Println("failed to delete workout performed")
		return c.String(http.StatusBadRequest, "failed to workout performed")
	}

	return c.String(http.StatusOK, "workout performed deleted.")
}
