package handler

import (
	"fmt"
	"net/http"

	authentication "github.com/Calgorr/IE_Backend_Fall/Authentication"
	model "github.com/Calgorr/IE_Backend_Fall/Model"
	"github.com/Calgorr/IE_Backend_Fall/database"
	"github.com/labstack/echo/v4"
)

func (h *Handler) Signup(c echo.Context) error {
	newUser, err := new(model.User).Bind(c)
	if err != nil {
		c.String(http.StatusBadRequest, "Something went wrong")
	}
	err = database.AddUser(newUser)
	return c.String(http.StatusOK, "Signed up")
}

func (h *Handler) Login(c echo.Context) error {
	user, err := new(model.User).Bind(c)
	if err != nil {
		c.String(http.StatusBadRequest, "Something went wrong")
	}
	u, err := database.GetUserByUsername(user.Username)
	if err != nil || u.Password != user.Password {
		c.String(http.StatusUnauthorized, "Wrong username or password")
	}
	token, err := authentication.GenerateJWT()
	if err != nil {
		c.String(http.StatusInternalServerError, "Something went wrong")
	}
	c.Response().Header().Set("Authorization", token)
	user.Token = token
	return c.String(http.StatusOK, "Logged in")
}

func (h *Handler) NewURL(c echo.Context) error {
	URL := c.FormValue("URL")
	fmt.Println(URL) //update
	return c.String(http.StatusOK, "URL added")
}
func (h *Handler) GetURLs(c echo.Context) error {
	urls := make([]string, 20)
	fmt.Println(urls) //update
	return c.JSONPretty(http.StatusOK, urls, " ")
}

func (h *Handler) StatURL(c echo.Context) error {
	x := 0                                     //update
	return c.JSONPretty(http.StatusOK, x, " ") //update
}

func (h *Handler) WanrURL(c echo.Context) error {
	warning := "warning" //update
	return c.JSON(http.StatusOK, warning)
}

func (h *Handler) Authenticate(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	fmt.Println(token) //update
	return nil
}

func (h *Handler) GetAlerts(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	fmt.Println(token) //update
	return nil
}
