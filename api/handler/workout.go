package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Submit workout
func submitWorkout(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}

// Get workout details
func getWorkoutDetails(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}

// Update workout details
func updateWorkout(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}

// Delete workout
func deleteWorkout(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}