package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Submit composition entry
func SubmitComposition(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}

// Get composition entry details
func GetCompositionDetails(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}

// Update composition entry details
func UpdateComposition(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}

// Delete composition entry
func DeleteComposition(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}
