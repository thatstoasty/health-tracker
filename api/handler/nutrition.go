package handler

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"

	"github.com/thatstoasty/health-tracker/queries"
)

// Submit nutrition entry
func SubmitNutrition(c echo.Context) error {
	var requestBody queries.SubmitNutritionParams
	connectionString := getConnectionString()
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
		log.Fatal("failed to establish connection to postgres")
		return c.String(http.StatusBadRequest, "failed to establish connection to postgres")
	}

	queries := queries.New(db)
	ctx := context.Background()

	// bind request body to variable given
	if err := c.Bind(&requestBody); err != nil {
		log.Fatal(err)
		log.Fatal("Failed to bind request body to nutrition type")
		return c.String(http.StatusBadRequest, "Failed to bind request body to nutrition type")
	}

	nutrition, err := queries.SubmitNutrition(ctx, requestBody)
	if err != nil {
		log.Fatal(err)
		log.Fatal("failed to get nutrition details")
		return c.String(http.StatusBadRequest, "failed to get nutrition details")
	}
	log.Println(nutrition)

	return c.JSON(http.StatusOK, nutrition)
}

// Get nutrition entry details
func GetNutrition(c echo.Context) error {
	connectionString := getConnectionString()
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
		log.Fatal("failed to establish connection to postgres")
		return c.String(http.StatusBadRequest, "failed to establish connection to postgres")
	}

	queries := queries.New(db)
	ctx := context.Background()
	date := c.Param("date")

	nutrition, err := queries.GetNutritionDetails(ctx, date)
	if err != nil {
		log.Fatal(err)
		log.Fatal("failed to get nutrition details")
		return c.String(http.StatusBadRequest, "failed to get nutrition details")
	}

	return c.JSON(http.StatusOK, nutrition)
}

// Delete nutrition entry
func DeleteNutrition(c echo.Context) error {
	connectionString := getConnectionString()
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
		log.Fatal("failed to establish connection to postgres")
		return c.String(http.StatusBadRequest, "failed to establish connection to postgres")
	}

	queries := queries.New(db)
	ctx := context.Background()
	date := c.Param("date")

	error := queries.DeleteNutrition(ctx, date)
	if error != nil {
		log.Fatal(err)
		log.Fatal("failed to delete nutrition entry")
		return c.String(http.StatusBadRequest, "failed to delete nutrition entry")
	}

	return c.String(http.StatusOK, "Nutrition entry deleted.")
}
