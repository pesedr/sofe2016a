package routes

import (
	"log"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/labstack/echo"
	"github.com/pesedr/sofe2016a/models"
	"github.com/pesedr/sofe2016a/repo"
)

type UserController struct{}

func (u *UserController) Create(c echo.Context) error {
	user := &models.User{
		ID: bson.NewObjectId(),
	}

	//change this to u instead of user as a bug
	err := c.Bind(user)
	if err != nil {
		return err
	}

	user, err = repo.User.Create(user)
	if err != nil {

	}

	return c.JSON(http.StatusCreated, user)

}

func (u *UserController) Get(c echo.Context) error {
	id := c.Param("id")

	user, err := repo.User.Get(id)
	if err != nil {
		log.Println(err)
	}

	return c.JSON(http.StatusOK, user)
}

func (u *UserController) Update(c echo.Context) error {
	user := new(models.User)

	err := c.Bind(user)
	if err != nil {
		return err
	}

	id := c.Param("id")
	user, err = repo.User.Update(id, user)
	if err != nil {
		log.Println(err)
	}

	return c.JSON(http.StatusOK, user)
}

func (u *UserController) Delete(c echo.Context) error {
	id := c.Param("id")

	err := repo.User.Delete(id)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}
