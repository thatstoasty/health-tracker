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

// Submit composition entry
func SubmitComposition(c echo.Context) error {
	var requestBody queries.SubmitCompositionParams

	db, err := sql.Open("postgres", "user=docker password=docker dbname=docker sslmode=disable")
	if err != nil {
		log.Println(err)
		log.Println("failed to establish connection to postgres")
		return c.String(http.StatusBadRequest, "failed to establish connection to postgres")
	}

	queries := queries.New(db)
	ctx := context.Background()

	if err := c.Bind(&requestBody); err != nil {
		log.Println(err)
		log.Println("Failed to bind request body to composition type")
		return c.String(http.StatusBadRequest, "Failed to bind request body to composition type")
	}

	composition, err := queries.SubmitComposition(ctx, requestBody)
	if err != nil {
		log.Println(err)
		log.Println("failed to get composition details")
		return c.String(http.StatusBadRequest, "failed to get composition details")
	}
	log.Println(composition)

	return c.JSON(http.StatusOK, composition)}

// Get composition entry details
func GetCompositionDetails(c echo.Context) error {
	db, err := sql.Open("postgres", "user=docker password=docker dbname=docker sslmode=disable")
	if err != nil {
		log.Println(err)
		log.Println("failed to establish connection to postgres")
		return c.String(http.StatusBadRequest, "failed to establish connection to postgres")
	}

	queries := queries.New(db)
	ctx := context.Background()
	date := c.Param("date")

	composition, err := queries.GetCompositionDetails(ctx, date)
	if err != nil {
		log.Println(err)
		log.Println("failed to get composition details")
		return c.String(http.StatusBadRequest, "failed to get composition details")
	}
	log.Println(composition)


	return c.JSON(http.StatusOK, composition)
}

// Update composition entry details
func UpdateComposition(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}

// Delete composition entry
func DeleteComposition(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}
