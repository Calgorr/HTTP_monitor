package handler

import (
	authentication "github.com/Calgorr/IE_Backend_Fall/Authentication"
	"github.com/labstack/echo/v4"
)

type Handler struct {
}

func extractID(c echo.Context) int {
	token := c.Request().Header.Get("Authorization")
	claims, err := authentication.ExtractClaimsFromToken(token)
	if err != nil {
		panic(err)
	}
	id := int(claims["id"].(float64))
	return id
}
