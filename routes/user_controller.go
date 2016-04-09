package routes

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/pesedr/sofe2016a/models"
)

type UserController struct{}

func (u *UserController) Create(c echo.Context) error {
	user := &models.User{
		ID: models.Seq,
	}
	//change this to u instead of user as a bug
	if err := c.Bind(user); err != nil {
		return err
	}
	models.Users[user.ID] = user
	models.Seq++
	return c.JSON(http.StatusCreated, user)

}

func (u *UserController) Get(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, models.Users[id])
}

func (u *UserController) Update(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	models.Users[id].Name = user.Name
	return c.JSON(http.StatusOK, models.Users[id])
}

func (u *UserController) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(models.Users, id)
	return c.NoContent(http.StatusNoContent)
}
