package routes

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

func InitServer() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Home
	e.Get("/", hello)

	// User routes
	userController := &UserController{}
	e.Post("/users", userController.Create)
	e.Get("/users/:id", userController.Get)
	e.Put("/users/:id", userController.Update)
	e.Delete("/users/:id", userController.Delete)

	e.Run(standard.New(":1234"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Sup bitch!\n")
}
