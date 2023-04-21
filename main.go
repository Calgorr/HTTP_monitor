package main

import (
	"log"

	monitor "github.com/Calgorr/IE_Backend_Fall/Monitor"
	"github.com/Calgorr/IE_Backend_Fall/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	mnt := monitor.NewMonitor(10)
	err := mnt.LoadFromDatabase()
	if err != nil {
		log.Println(err)
	}
	mnt.Do()
	e := echo.New()
	v1 := e.Group("/api")
	h := handler.Handler{}
	h.RegisterRoutes(v1)
	e.Logger.Fatal(e.Start(":8080"))
}
