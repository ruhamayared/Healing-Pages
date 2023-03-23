package main

import (
	"github.com/ruhamayared/healing-pages/src/database"
	"github.com/ruhamayared/healing-pages/src/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Connect to the database
	database.ConnectDB()

	// Initialize a new instance of Echo
	e := echo.New()

	// Add middleware for logging and recovering from panics
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes for CRUD operations
	e.POST("/entries", handlers.CreateEntry)

	e.GET("/entries/:id", func(c echo.Context) error {
		// Pass the GORM database instance to the handler
		return handlers.GetEntry(c, database.DB)
	})

	e.GET("/entries", handlers.GetAllEntries)

	e.PUT("/entries/:id", func(c echo.Context) error {
		// Pass the GORM database instance to the handler
		return handlers.UpdateEntry(c, database.DB)
	})

	e.DELETE("/entries/:id", handlers.DeleteEntry)

	// Start the server and listen on port 8080
	e.Logger.Fatal(e.Start(":8080"))
}
