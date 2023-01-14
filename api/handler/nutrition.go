package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Submit nutrition entry
func SubmitNutrition(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}

// Get nutrition entry details
func GetNutritionDetails(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}

// Update nutrition entry details
func UpdateNutrition(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}

// Delete nutrition entry
func DeleteNutrition(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}
