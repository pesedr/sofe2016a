package server

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}

func router(e *echo.Echo) {
	// Home
	e.Get("/", hello)

	// Auth
	authController := &AuthController{}
	e.POST("/login", authController.Login)

	// User Authentication
	u := e.Group("/users")
	u.Use(middleware.JWTAuth([]byte("secret")))

	// User controller
	userController := &UserController{}
	u.Get("", authController.Restricted)
	u.Post("/signup", userController.Create)
	u.Get("/:id", userController.Get)
	u.Put("/:id", userController.Update)
	u.Delete("/:id", userController.Delete)

}

func serveJSON(c echo.Context, response interface{}) error {
	return c.JSON(http.StatusOK, response)
}
