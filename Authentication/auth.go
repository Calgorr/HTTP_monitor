package authentication

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	echo "github.com/labstack/echo/v4"
)

var secretKey = []byte("calgor")

const secret = "calgor"

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

func ValidateJWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		auth := c.Request().Header.Get("Authorization")
		token, err := jwt.Parse(auth, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("there was an error")
			}
			return []byte(secret), nil
		})
		if err != nil {
			fmt.Println(err, token)
			return c.String(http.StatusUnauthorized, "unauthorized")
		}
		if token.Valid {
			return next(c)
		}
		return c.String(http.StatusUnauthorized, "unauthorized")
	}
}
