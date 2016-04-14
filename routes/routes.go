package routes

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

func InitServer() {
	e := echo.New()

	router(e)

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "INFO: ${time_rfc3339} method=${method} uri=${uri} status=${status} ",
		Output: os.Stdout}))
	e.SetLogPrefix("INFO: ")

	e.Use(middleware.Recover())
	e.SetHTTPErrorHandler(apiErrorHandler)

	log.Println("Serving on port 3000")
	e.Run(standard.New(":3001"))
}

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
}
