package main

import (
	"context"
	"log"

	"github.com/ruhamayared/healing-pages/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	// Connect to the database
	conn, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	// Initialize a new instance of Echo
	e := echo.New()

	// Routes for CRUD operations
	e.POST("/entries", handlers.CreateEntry)
	e.GET("/entries/:id", handlers.GetEntry)
	e.GET("/entries", handlers.GetAllEntries)
	e.PUT("/entries/:id", handlers.UpdateEntry)
	e.DELETE("/entries/:id", handlers.DeleteEntry)

	// Start the server and listen on port 8080
	e.Logger.Fatal(e.Start(":8080"))
}
