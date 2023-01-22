package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Submit workout
func SubmitWorkout(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}

// Get workout details
func GetWorkoutDetails(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}

// Delete workout
func DeleteWorkout(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}
