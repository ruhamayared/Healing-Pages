package main

import (
	"github.com/labstack/echo/v4/middleware"
	"github.com/ruhamayared/healing-pages/src/database"
	"github.com/ruhamayared/healing-pages/src/handlers"

	// Import your custom JWT middleware from "src/userauth"
	"github.com/labstack/echo/v4"
)

func main() {
	// Connect to the database
	database.ConnectDB()

	// Initialize a new instance of Echo
	e := echo.New()

	// Set up CORS middleware
	e.Use(middleware.CORS())

	// Routes for CRUD operations, all protected with JWT middleware
	e.POST("/entries", handlers.CreateEntry)
	e.GET("/entries/:id", handlers.GetEntry)
	e.GET("/entries", handlers.GetAllEntries)
	e.PUT("/entries/:id", handlers.UpdateEntry)
	e.DELETE("/entries/:id", handlers.DeleteEntry)

	// Start the server and listen on port 8080
	e.Logger.Fatal(e.Start(":8080"))
}
