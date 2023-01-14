package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/thatstoasty/health-tracker/handler"
)


func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	//// Workout
	e.POST("/workout", handler.submitWorkout)
	e.GET("/workout", handler.getWorkoutDetails)
	e.PATCH("/workout", handler.updateWorkout)
	e.DELETE("/workout", handler.deleteWorkout)

	//// Composition
	e.POST("/composition", handler.submitComposition)
	e.GET("/composition", handler.getCompositionDetails)
	e.PATCH("/composition", handler.updateComposition)
	e.DELETE("/composition", handler.deleteComposition)

	//// Exercise
	e.GET("/exercise", handler.getExercise)

	//// Composition
	e.POST("/nutrition", handler.submitNutrition)
	e.GET("/nutrition", handler.getNutritionDetails)
	e.PATCH("/nutrition", handler.updateNutrition)
	e.DELETE("/nutrition", handler.deleteNutrition)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
