package handler

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"

	"github.com/thatstoasty/health-tracker/queries"
	"github.com/thatstoasty/health-tracker/utils"
)

// Submit composition entry
func SubmitComposition(c echo.Context) error {
	var requestBody queries.SubmitCompositionParams
	connectionString := utils.GetConnectionString()
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
		log.Fatal("Failed to bind request body to composition type")
		return c.String(http.StatusBadRequest, "Failed to bind request body to composition type")
	}

	composition, err := queries.SubmitComposition(ctx, requestBody)
	if err != nil {
		log.Fatal(err)
		log.Fatal("failed to get composition details")
		return c.String(http.StatusBadRequest, "failed to get composition details")
	}
	log.Println(composition)

	return c.JSON(http.StatusOK, composition)}

// Get composition entry details
func GetComposition(c echo.Context) error {
	connectionString := utils.GetConnectionString()
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
		log.Fatal("failed to establish connection to postgres")
		return c.String(http.StatusBadRequest, "failed to establish connection to postgres")
	}

	queries := queries.New(db)
	ctx := context.Background()
	date := c.Param("date")

	composition, err := queries.GetComposition(ctx, date)
	if err != nil {
		log.Fatal(err)
		log.Fatal("failed to get composition details")
		return c.String(http.StatusBadRequest, "failed to get composition details")
	}

	return c.JSON(http.StatusOK, composition)
}

// Delete composition entry
func DeleteComposition(c echo.Context) error {
	connectionString := utils.GetConnectionString()
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
		log.Fatal("failed to establish connection to postgres")
		return c.String(http.StatusBadRequest, "failed to establish connection to postgres")
	}

	queries := queries.New(db)
	ctx := context.Background()
	date := c.Param("date")

	error := queries.DeleteComposition(ctx, date)
	if error != nil {
		log.Fatal(err)
		log.Fatal("failed to delete composition")
		return c.String(http.StatusBadRequest, "failed to delete composition")
	}

	return c.String(http.StatusOK, "Composition deleted.")
}
