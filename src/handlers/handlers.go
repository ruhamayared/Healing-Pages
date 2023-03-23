package handlers

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/ruhamayared/healing-pages/src/database"
	"github.com/ruhamayared/healing-pages/src/models"
	"gorm.io/gorm"

	"github.com/labstack/echo/v4"
)

func CreateEntry(c echo.Context) error {
	// Initialize a new Entry object
	entry := new(models.Entry)

	// Bind the request body to the Entry object
	if err := c.Bind(&entry); err != nil {
		return c.String(http.StatusBadRequest, "invalid entry")
	}

	// Set the entry's CreatedAt field to the current time
	entry.CreatedAt = time.Now()

	// Use the Create method of the *gorm.DB object to insert the new entry into the database
	if err := database.DB.Create(&entry).Error; err != nil {
		return c.String(http.StatusInternalServerError, "Database error")
	}

	// Return a JSON response with the newly created entry and a status code of 201 (Created)
	return c.JSON(http.StatusCreated, entry)
}

func GetEntry(c echo.Context) error {
	// Extract the ID from the URL path parameter
	id := c.Param("id")

	// Create a new Entry object to hold the result
	var entry models.Entry

	// Query the database for the Entry with the specified ID
	result := database.DB.First(&entry, id)

	// Check if the Entry was not found
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return c.String(http.StatusNotFound, "Entry not found")
	}

	// Check for other database errors
	if result.Error != nil {
		return c.String(http.StatusInternalServerError, "Database error")
	}

	// Return the Entry as JSON
	return c.JSON(http.StatusOK, entry)
}

func UpdateEntry(c echo.Context) error {
	// Extract the ID from the URL path parameter
	id := c.Param("id")

	// Query the database for the Entry with the specified ID
	var entry models.Entry
	result := database.DB.First(&entry, id)

	// Check if the Entry was not found
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return c.String(http.StatusNotFound, "Entry not found")
	}

	// Check for other database errors
	if result.Error != nil {
		return c.String(http.StatusInternalServerError, "Database error")
	}

	// Bind the updated Entry data from the request body
	if err := c.Bind(&entry); err != nil {
		return c.String(http.StatusBadRequest, "Invalid entry")
	}

	// Update the Entry in the database
	result = database.DB.Save(&entry)

	// Check for database errors
	if result.Error != nil {
		return c.String(http.StatusInternalServerError, "database error")
	}

	// Return the updated Entry as JSON
	return c.JSON(http.StatusOK, entry)
}

func DeleteEntry(c echo.Context) error {
	// Get ID from URL parameter and convert to integer
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// Return error response if ID is not valid
		return c.String(http.StatusBadRequest, "invalid id")
	}

	// Delete entry from database with matching ID
	if err := database.DB.Where("id = ?", id).Delete(&models.Entry{}).Error; err != nil {
		// If error is "record not found" return appropriate error response
		if err == gorm.ErrRecordNotFound {
			return c.String(http.StatusNotFound, "entry not found")
		}
		// If error is not "record not found" return the error
		return err
	}

	// Return success response with no content
	return c.NoContent(http.StatusNoContent)
}

func GetAllEntries(c echo.Context) error {
	// Initialize an empty slice of Entry structs
	entries := make([]*models.Entry, 0)

	// Query the database using GORM's Find method and store the result in the entries slice
	err := database.DB.Find(&entries).Error
	if err != nil {
		return err
	}

	// Return a JSON response with the entries slice
	return c.JSON(http.StatusOK, entries)
}
