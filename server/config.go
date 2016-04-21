package server

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

func Config() {
	// log.Println("Creating a new echo service")
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	router(e)
	// log.Println("Serving on port 3000")
	e.Run(standard.New(":3001"))
}

func router(e *echo.Echo) {
	// Home
	e.Get("/", hello)

	// User routes
	userController := &UserController{}
	e.Post("/users", userController.Create)
	e.Get("/users/:id", userController.Get)
	e.Put("/users/:id", userController.Update)
	e.Delete("/users/:id", userController.Delete)
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}

func serveJSON(c echo.Context, response interface{}) error {
	return c.JSON(http.StatusOK, response)
}
