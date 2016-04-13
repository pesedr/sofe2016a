package server

import (
	"log"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/labstack/echo"
	"github.com/pesedr/sofe2016a/errors"
	"github.com/pesedr/sofe2016a/models"
	"github.com/pesedr/sofe2016a/repo"
)

type UserController struct{}

func (u *UserController) Create(c echo.Context) error {
	user := &models.User{
		ID: bson.NewObjectId(),
	}

	log.Println("Binding user to request")
	err := c.Bind(u)
	if err != nil {
		log.Println("Binding failed", "error:", err.Error())
		return errors.NewApiError(errors.GeneraJSONError, err.Error())
	}

	u, err = repo.User.Create(user)
	if err != nil {
		return err
	}

	return serveJSON(c, user)
}

func (u *UserController) Get(c echo.Context) error {
	id := c.Param("id")

	user, err := repo.User.Get(id)
	if err != nil {
		return err
	}

	return serveJSON(c, user)
}

func (u *UserController) Update(c echo.Context) error {
	user := new(models.User)

	log.Println("Binding user to request")
	err := c.Bind(user)
	if err != nil {
		log.Println("Binding failed", "error:", err.Error())
		return errors.NewApiError(errors.GeneraJSONError, err.Error())
	}

	id := c.Param("id")
	user, err = repo.User.Update(id, user)
	if err != nil {
		return err
	}

	return serveJSON(c, user)
}

func (u *UserController) Delete(c echo.Context) error {
	id := c.Param("id")

	err := repo.User.Delete(id)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}
