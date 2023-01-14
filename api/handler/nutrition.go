package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Submit nutrition entry
func submitNutrition(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}

// Get nutrition entry details
func getNutritionDetails(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}

// Update nutrition entry details
func updateNutrition(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}

// Delete nutrition entry
func deleteNutrition(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}