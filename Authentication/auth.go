package authentication

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const secret = "calgor"

type whiteList struct {
	path   string
	method string
}

var authWhiteList []whiteList

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(24 * time.Hour).Unix()
	claims["authorized"] = true
	claims["user"] = "username"
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
func AddToWhiteList(path string, method string) {
	if authWhiteList == nil {
		authWhiteList = make([]whiteList, 0)
	}
	authWhiteList = append(authWhiteList, whiteList{path, method})
}

func skipper(c echo.Context) bool {
	for _, v := range authWhiteList {
		if c.Path() == v.path && c.Request().Method == v.method {
			return true
		}
	}
	return false
}

func ValidateJWT() echo.MiddlewareFunc {
	c := middleware.DefaultJWTConfig
	c.SigningKey = secret
	c.Skipper = skipper
	return middleware.JWTWithConfig(c)
}
