package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"

	"github.com/thatstoasty/health-tracker/queries"
	"github.com/thatstoasty/health-tracker/utils"
)

// Get Program
func GetProgram(c echo.Context) error {
	db, err := utils.GetDBConnection(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, GenericResponse{fmt.Sprintf("Failed to establish connection to postgres: %s", err)})
	}

	queries := queries.New(db)
	ctx := context.Background()
	name := c.Param("name")

	composition, err := queries.GetProgram(ctx, name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, GenericResponse{fmt.Sprintf("Failed to get Program: %s", err)})
	}

	return c.JSON(http.StatusOK, composition)
}

// Get Program names
func GetProgramNames(c echo.Context) error {
	db, err := utils.GetDBConnection(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, GenericResponse{fmt.Sprintf("Failed to establish connection to postgres: %s", err)})
	}

	queries := queries.New(db)
	ctx := context.Background()
	limitString := c.QueryParam("limit")
	limit, err := strconv.Atoi(limitString)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, GenericResponse{fmt.Sprintf("Failed to convert limit to integer: %s", err)})
	}

	composition, err := queries.GetProgramNames(ctx, int32(limit))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, GenericResponse{fmt.Sprintf("Failed to get Program names: %s", err)})
	}

	return c.JSON(http.StatusOK, composition)
}

// Delete Program
func DeleteProgram(c echo.Context) error {
	db, err := utils.GetDBConnection(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, GenericResponse{fmt.Sprintf("Failed to establish connection to postgres: %s", err)})
	}

	queries := queries.New(db)
	ctx := context.Background()
	name := c.Param("name")

	error := queries.DeleteProgram(ctx, name)
	if error != nil {
		return c.JSON(http.StatusInternalServerError, GenericResponse{fmt.Sprintf("Failed to delete Program: %s", error)})
	}

	return c.JSON(http.StatusOK, GenericResponse{fmt.Sprintf("Successfully deleted Program: %s", name)})
}
