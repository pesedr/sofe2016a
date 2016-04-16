package server

import (
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

func Config() {
	log.Println("Creating a new echo service")
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	router(e)
	log.Println("Serving on port 3000")
	e.Run(standard.New(":3001"))
}
