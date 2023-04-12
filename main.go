package main

import (
	"github.com/Calgorr/IE_Backend_Fall/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	h := new(handler.Handler)
	h.RegisterRoutes(*e.Group("/api"))
	e.Logger.Fatal(e.Start(":8080"))
}
