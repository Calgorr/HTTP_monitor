package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	authentication "github.com/Calgorr/IE_Backend_Fall/Authentication"
	"github.com/labstack/echo/v4"
)

func signup(c echo.Context) error {
	userPass := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&userPass)
	if err != nil {
		c.String(http.StatusBadRequest, "Something went wrong")
	}
	username := userPass["username"]
	password := userPass["password"]
	fmt.Println(username, password) //update
	return c.String(http.StatusOK, "Registration done")
}

func login(c echo.Context) error {
	userPass := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&userPass)
	if err != nil {
		c.String(http.StatusBadRequest, "Something went wrong")
	}
	username := userPass["username"]
	password := userPass["password"]
	token, err := authentication.GenerateJWT()
	if err != nil {
		c.String(http.StatusInternalServerError, "Something went wrong")
	}
	c.Response().Header().Set("Authorization", token)
	fmt.Println(username, password) //update
	return c.String(http.StatusOK, "Logged in")
}

func newURL(c echo.Context) error {
	URL := c.FormValue("URL")
	fmt.Println(URL) //update
	return c.String(http.StatusOK, "URL added")
}
func getURLs(c echo.Context) error {
	urls := make([]string, 20)
	fmt.Println(urls) //update
	return c.JSONPretty(http.StatusOK, urls, " ")
}

func statURL(c echo.Context) error {
	x := 0                                     //update
	return c.JSONPretty(http.StatusOK, x, " ") //update
}

func wanrURL(c echo.Context) error {
	warning := "warning" //update
	return c.JSON(http.StatusOK, warning)
}

func authenticate(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	fmt.Println(token) //update
	return nil
}
