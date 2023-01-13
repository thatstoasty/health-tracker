package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Workout functions
func submit_workout(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}

func get_workout_details(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}

func update_workout(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}

func delete_workout(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}

// Composition functions
func submit_composition(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}

func get_composition_details(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}

func update_composition(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}

func delete_composition(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	//// Workout
	e.POST("/workout", submit_workout)
	e.GET("/workout", get_workout_details)
	e.PATCH("/workout", update_workout)
	e.DELETE("/workout", delete_workout)

	//// Composition
	e.POST("/composition", submit_composition)
	e.GET("/composition", get_composition_details)
	e.PATCH("/composition", update_composition)
	e.DELETE("/composition", delete_composition)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
