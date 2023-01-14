package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Nutrition functions
func submitNutrition(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}

func getNutritionDetails(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}

func updateNutrition(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}

func deleteNutrition(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}