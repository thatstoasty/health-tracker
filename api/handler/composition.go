package handler

import (
	"context"
	"database/sql"
	"log"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"

	"github.com/thatstoasty/health-tracker/queries"
	// "github.com/thatstoasty/health-tracker/types"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "docker"
	password = "docker"
	dbname   = "docker"
	sslmode  = "disable"
  )

func getConnectionString() string {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
    host, port, user, password, dbname, sslmode)

	return connectionString
}

// Submit composition entry
func SubmitComposition(c echo.Context) error {
	var requestBody queries.SubmitCompositionParams
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

	composition, err := queries.GetCompositionDetails(ctx, date)
	if err != nil {
		log.Fatal(err)
		log.Fatal("failed to get composition details")
		return c.String(http.StatusBadRequest, "failed to get composition details")
	}

	return c.JSON(http.StatusOK, composition)
}

// Update composition entry details
func UpdateComposition(c echo.Context) error {
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

	error := queries.UpdateComposition(ctx, date)
	if error != nil {
		log.Fatal(err)
		log.Fatal("failed to update composition")
		return c.String(http.StatusBadRequest, "failed to update composition")
	}

	return c.String(http.StatusOK, "Composition updated.")
}

// Delete composition entry
func DeleteComposition(c echo.Context) error {
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

	error := queries.DeleteComposition(ctx, date)
	if error != nil {
		log.Fatal(err)
		log.Fatal("failed to delete composition")
		return c.String(http.StatusBadRequest, "failed to delete composition")
	}

	return c.String(http.StatusOK, "Composition deleted.")
}
