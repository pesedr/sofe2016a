package server

import (
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/pesedr/sofe2016a/auth"
	"github.com/pesedr/sofe2016a/errors"
	"github.com/pesedr/sofe2016a/models"
)

type AuthController struct{}

func (a *AuthController) Restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	name := user.Claims["name"].(string)
	return c.String(http.StatusOK, "Welcome "+name+"!")
}

func (a *AuthController) Login(c echo.Context) error {
	log.Println("Login Controller")
	authCreds := &models.AuthCredentials{}

	log.Println("Binding credentials")
	err := c.Bind(authCreds)
	if err != nil {
		log.Println("Binding failed", "error:", err.Error())
		return errors.NewApiError(errors.GeneraJSONError, err.Error())
	}

	user, err := auth.Auth.CheckCredentials(authCreds)
	if err != nil {
		return err
	}
	token, err := auth.Auth.CreateToken(user)
	if err != nil {
		return err
	}
	return serveJSON(c, map[string]string{
		"token": token,
	})

}
