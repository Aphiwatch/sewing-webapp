package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Route for /api
	e.GET("/api", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello from Go Echo!")
	})

	// Start the server on port 8080
	e.Logger.Fatal(e.Start(":8080"))
}
