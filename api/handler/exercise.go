package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Get exercise details
func getExerciseDetails(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}