package handler

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"

	"github.com/thatstoasty/health-tracker/queries"
	// "github.com/thatstoasty/health-tracker/types"
)

// func mapComposition(ob1 queries.TrackerComposition) types.Composition {
// 	ob2 := &types.Composition { 
// 		Date: queries.TrackerComposition.Date, 
// 		Weight: queries.TrackerComposition.Weight, 
// 		Bodyfat: queries.TrackerComposition.Bodyfat, 
// 	}
// 	return ob2
// }

// Submit composition entry
func SubmitComposition(c echo.Context) error {
	var requestBody queries.SubmitCompositionParams

	db, err := sql.Open("postgres", "user=docker password=docker dbname=docker sslmode=disable")
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
func GetCompositionDetails(c echo.Context) error {
	db, err := sql.Open("postgres", "user=docker password=docker dbname=docker sslmode=disable")
	if err != nil {
		log.Fatal(err)
		log.Fatal("failed to establish connection to postgres")
		return c.String(http.StatusBadRequest, "failed to establish connection to postgres")
	}

	queries := queries.New(db)
	ctx := context.Background()
	date := c.Param("date")

	composition, err := queries.GetCompositionDetails(ctx, date)
	if err != nil {
		log.Fatal(err)
		log.Fatal("failed to get composition details")
		return c.String(http.StatusBadRequest, "failed to get composition details")
	}

	// result := mapComposition(composition)

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
