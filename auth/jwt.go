package auth

import (
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pesedr/sofe2016a/errors"
	"github.com/pesedr/sofe2016a/models"
	"github.com/pesedr/sofe2016a/repo"
)

const (
	SigningKey          = "somethingSuperSecretlol"
	tokenExpirationTime = time.Hour * 1
)

type Authorization interface {
	CheckCredentials(authCreds *models.AuthCredentials) (*models.User, error)
	CreateToken(user *models.User) (string, error)
}

type authorization struct{}

var Auth Authorization

func init() {
	Auth = &authorization{}
}

func (a *authorization) CheckCredentials(authCreds *models.AuthCredentials) (*models.User, error) {
	log.Println("Checking credentials user", "username", authCreds.Username)
	if authCreds.Username != "jonsnow" && authCreds.Password != "supersecretpassword!" {
		log.Println("invalid credentials")
		return nil, errors.NewApiError(errors.Unauthorized, "Invalid username and/or password")
	}
	log.Println("Valid credentials!")
	return repo.User.GetByUsername(authCreds.Username)
}

func (a *authorization) CreateToken(user *models.User) (string, error) {
	log.Println("creating a token for username:", user.Username)
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	token.Claims["name"] = user.Username
	token.Claims["exp"] = expireToken()

	// Generate encoded token and send it as response
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", errors.NewApiError(errors.Unauthorized, fmt.Sprintf("there was an error signing the token", "error", err.Error()))
	}
	return t, nil

}

func expireToken() int64 {
	return time.Now().Add(tokenExpirationTime).Unix()
}
