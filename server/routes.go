package server

import (
	"net/http"

	"github.com/labstack/echo"
)

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
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

	// Account routes
	// e.Get("/balance", accountController.GetAccounts)
}

func serveJSON(c echo.Context, response interface{}) error {
	return c.JSON(http.StatusOK, response)
}
