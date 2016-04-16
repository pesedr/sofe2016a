package server

import (
	"log"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

func Config() {
	e := echo.New()

	router(e)

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "INFO: ${time_rfc3339} method=${method} uri=${uri} status=${status} ",
		Output: os.Stdout}))
	e.SetLogPrefix("INFO: ")

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	log.Println("Serving on port 3000")
	e.Run(standard.New(":3001"))
}
