package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)


func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	//// Workout
	e.POST("/workout", handler.submit_workout)
	e.GET("/workout", handler.get_workout_details)
	e.PATCH("/workout", handler.update_workout)
	e.DELETE("/workout", handler.delete_workout)

	//// Composition
	e.POST("/composition", handler.submit_composition)
	e.GET("/composition", handler.get_composition_details)
	e.PATCH("/composition", handler.update_composition)
	e.DELETE("/composition", handler.delete_composition)

	//// Exercise
	e.GET("/exercise", handler.get_exercise)

	//// Composition
	e.POST("/nutrition", handler.submit_nutrition)
	e.GET("/nutrition", handler.get_nutrition_details)
	e.PATCH("/nutrition", handler.update_nutrition)
	e.DELETE("/nutrition", handler.delete_nutrition)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
