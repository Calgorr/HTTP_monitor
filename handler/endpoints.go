package endpoint

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterRoutes(v echo.Group) {
	v.Use(middleware.Logger())
}
