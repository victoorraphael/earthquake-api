package main

import (
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"os"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Receiving Query Params
func Create(c echo.Context) error {
	name := c.QueryParam("name")
	email := c.QueryParam("email")

	p := User{name, email}

	return c.JSON(http.StatusOK, p)
}

// Get router params /user/:id
func GetUser(c echo.Context) error {
	id := c.Param("id")

	return c.JSON(http.StatusOK, map[string]string{"ID": id})
}

// Get info from form
func GetUserForm(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")

	return c.String(http.StatusOK, "Nome:"+name+", Email: "+email)
}

// Get file from Form Data
func SaveUser(c echo.Context) error {
	file, err := c.FormFile("avatar")

	if err != nil {
		return err
	}

	src, errc := file.Open()
	if errc != nil {
		return errc
	}
	defer src.Close()

	dst, errd := os.Create(file.Filename)
	if errd != nil {
		return errd
	}
	defer dst.Close()

	if _, errcp := io.Copy(dst, src); errcp != nil {
		return errcp
	}

	return c.HTML(http.StatusOK, "<b>Thank you!")
}

// Binding body data
func Binding(c echo.Context) error {
	u := new(User)

	if err := c.Bind(u); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, u)
}
