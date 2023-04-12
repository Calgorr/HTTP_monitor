package main

import (
	authentication "github.com/Calgorr/IE_Backend_Fall/Authentication"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.POST("/signup", signup)
	e.POST("/login", login)
	e.POST("/newURL", newURL, authentication.VerifyToken)
	e.GET("/user/URL/getURLs", getURLs, authentication.VerifyToken)
	e.GET("/user/URL/URLstatistics", statURL, authentication.VerifyToken)
	e.GET("/user/warning/:id", wanrURL, authentication.VerifyToken)
	e.Logger.Fatal(e.Start(":8080"))
}
