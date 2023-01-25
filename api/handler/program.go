package handler

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"

	"github.com/thatstoasty/health-tracker/queries"
	"github.com/thatstoasty/health-tracker/utils"
)

// Get program
func GetProgram(c echo.Context) error {
	connectionString := utils.GetConnectionString()
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Println(err)
		log.Println("failed to establish connection to postgres")
		return c.String(http.StatusBadRequest, "failed to establish connection to postgres")
	}

	queries := queries.New(db)
	ctx := context.Background()
	name := c.Param("name")

	composition, err := queries.GetProgram(ctx, name)
	if err != nil {
		log.Println(err)
		log.Println("failed to get program details")
		return c.String(http.StatusBadRequest, "failed to get program details")
	}

	return c.JSON(http.StatusOK, composition)
}

// Get workout names
func GetProgramNames(c echo.Context) error {
	connectionString := utils.GetConnectionString()
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Println(err)
		log.Println("failed to establish connection to postgres")
		return c.String(http.StatusBadRequest, "failed to establish connection to postgres")
	}

	queries := queries.New(db)
	ctx := context.Background()
	limitString := c.QueryParam("limit")
	limit, err := strconv.Atoi(limitString)
	if err != nil {
		log.Println(err)
		log.Println("failed to convert to int")
		return c.String(http.StatusBadRequest, "failed to convert to int")
	}

	composition, err := queries.GetProgramNames(ctx, int32(limit))
	if err != nil {
		log.Println(err)
		log.Println("failed to get workout")
		return c.String(http.StatusBadRequest, "failed to get workout")
	}

	return c.JSON(http.StatusOK, composition)
}

// Delete program
func DeleteProgram(c echo.Context) error {
	connectionString := utils.GetConnectionString()
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Println(err)
		log.Println("failed to establish connection to postgres")
		return c.String(http.StatusBadRequest, "failed to establish connection to postgres")
	}

	queries := queries.New(db)
	ctx := context.Background()
	name := c.Param("name")

	error := queries.DeleteProgram(ctx, name)
	if error != nil {
		log.Println(err)
		log.Println("failed to delete program")
		return c.String(http.StatusBadRequest, "failed to delete program")
	}

	return c.String(http.StatusOK, "program deleted.")
}
