package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Composition functions
func submitComposition(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}

func getCompositionDetails(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}

func updateComposition(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}

func deleteComposition(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}