package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Workout functions
func submitWorkout(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}

func getWorkoutDetails(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}

func updateWorkout(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}

func deleteWorkout(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}