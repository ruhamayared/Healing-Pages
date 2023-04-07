package main

import (
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	"github.com/ruhamayared/healing-pages/src/database"
	"github.com/ruhamayared/healing-pages/src/handlers"
)

var sessionStore = sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))

func init() {
	goth.UseProviders(
		google.New(os.Getenv("GOOGLE_KEY"), os.Getenv("GOOGLE_SECRET"), "http://localhost:8080/auth/google/callback"),
	)
}

func loginHandler(c echo.Context) error {
	gothic.BeginAuthHandler(c.Response(), c.Request())
	return nil
}

func callbackHandler(c echo.Context) error {
	user, err := gothic.CompleteUserAuth(c.Response(), c.Request())
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	session, _ := sessionStore.Get(c.Request(), "user-session")
	session.Values["user_id"] = user.UserID
	session.Values["provider_name"] = user.Provider
	session.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusFound, "/")
}

func isAuthenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		session, _ := sessionStore.Get(c.Request(), "user-session")
		userID := session.Values["user_id"]

		if userID == nil {
			return echo.ErrUnauthorized
		}
		return next(c)
	}
}

func main() {
	// Connect to the database
	database.ConnectDB()

	// Initialize a new instance of Echo
	e := echo.New()

	// Set up CORS middleware
	e.Use(middleware.CORS())

	// Routes for user authentication
	e.GET("/auth/google/login", loginHandler)
	e.GET("/auth/google/callback", callbackHandler)

	// Routes for CRUD operations, all protected with isAuthenticated middleware
	e.POST("/entries", handlers.CreateEntry, isAuthenticated)
	e.GET("/entries/:id", handlers.GetEntry, isAuthenticated)
	e.GET("/entries", handlers.GetAllEntries, isAuthenticated)
	e.PUT("/entries/:id", handlers.UpdateEntry, isAuthenticated)
	e.DELETE("/entries/:id", handlers.DeleteEntry, isAuthenticated)

	// Start the server and listen on port 8080
	e.Logger.Fatal(e.Start(":8080"))
}
