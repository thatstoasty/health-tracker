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
	e.POST("/workout", handler.SubmitWorkout)
	e.GET("/workout", handler.GetWorkoutDetails)
	e.PATCH("/workout", handler.UpdateWorkout)
	e.DELETE("/workout", handler.DeleteWorkout)

	//// Composition
	e.POST("/composition", handler.SubmitComposition)
	e.GET("/composition", handler.GetCompositionDetails)
	e.PATCH("/composition", handler.UpdateComposition)
	e.DELETE("/composition", handler.DeleteComposition)

	//// Exercise
	e.GET("/exercise", handler.GetExerciseDetails)

	//// Composition
	e.POST("/nutrition", handler.SubmitNutrition)
	e.GET("/nutrition", handler.GetNutritionDetails)
	e.PATCH("/nutrition", handler.UpdateNutrition)
	e.DELETE("/nutrition", handler.DeleteNutrition)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
