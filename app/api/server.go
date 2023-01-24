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
	//// Exercise
	e.GET("/exercise", handler.GetExerciseNames)
	e.GET("/exercise/:name", handler.GetExerciseDetails)
	e.DELETE("/exercise/:name", handler.DeleteExercise)

	//// Workout
	e.GET("/workout", handler.GetWorkoutNames)
	e.DELETE("/workout/:name", handler.DeleteWorkout)
	e.GET("/workout/:name", handler.GetWorkoutDetails)
	e.GET("/workout/:name/:date", handler.GetWorkoutPerformed)
	e.DELETE("/workout/:name/:date", handler.DeleteWorkoutPerformed)

	//// Program
	e.GET("/exercise", handler.GetProgramNames)
	e.GET("/exercise/:name", handler.GetProgramDetails)


	//// Composition
	e.POST("/composition", handler.SubmitComposition)
	e.GET("/composition/:date", handler.GetCompositionDetails)
	e.DELETE("/composition/:date ", handler.DeleteComposition)





	//// Composition
	e.POST("/nutrition", handler.SubmitNutrition)
	e.GET("/nutrition/:date", handler.GetNutritionDetails)
	e.DELETE("/nutrition/:date", handler.DeleteNutrition)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
