package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Submit composition entry
func submitComposition(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}

// Get composition entry details
func getCompositionDetails(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}

// Update composition entry details
func updateComposition(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}

// Delete composition entry
func deleteComposition(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}