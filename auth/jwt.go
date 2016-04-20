package auth

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

const (
	Bearer     = "Bearer"
	SigningKey = "somethingSuperSecretlol"
)

func Auth(key string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			auth := c.Request().Header().Get("Authorization")
			l := len(Bearer)
			he := echo.ErrUnauthorized

			if len(auth) > l+1 && auth[:l] == Bearer {
				t, err := jwt.Parse(auth[l+1:], func(token *jwt.Token) (interface{}, error) {

					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
					}

					return []byte(key), nil
				})
				if err == nil && t.Valid {
					c.Set("claims", t.Claims)
					return next(c)
				}
			}
			return he
		}
	}
}

func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "NO auth required for this route.\n")
}

func restricted(c echo.Context) error {
	return c.String(http.StatusOK, "Access granted with JWT.\n")
}
