package handler

import (
	authentication "github.com/Calgorr/IE_Backend_Fall/Authentication"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (h *Handler) RegisterRoutes(v echo.Group) {
	v.Use(middleware.Logger())
	v.Use(authentication.ValidateJWT)

	userGroup := v.Group("/users")
	userGroup.POST("", h.Signup)
	userGroup.POST("/login", h.Login)

	urlGroup := v.Group("/urls")
	urlGroup.GET("", h.GetURLs)
	urlGroup.POST("", h.NewURL)
	urlGroup.GET(":id", h.StatURL)

	alertGroup := v.Group("/alerts")
	alertGroup.POST("", h.GetAlerts)
}
