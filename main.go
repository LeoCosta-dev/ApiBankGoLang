package main

import (
	"main/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo Instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routers
	e.GET("/", handlers.Hello)

	// Start Server
	e.Logger.Fatal(e.Start(":3000"))

}
