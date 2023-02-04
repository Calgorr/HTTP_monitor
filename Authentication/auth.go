package authentication

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

var secretKey = []byte("calgor")

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodEdDSA)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, err
}

func VerifyToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		if c.Request().Header.Get("Authorization") != "" {
			token, err := jwt.Parse(c.Request().Header.Get("Authorization"), func(token *jwt.Token) (interface{}, error) {
				_, ok := token.Method.(*jwt.SigningMethodECDSA)
				if !ok {
					c.String(http.StatusUnauthorized, "You're Unauthorized!")
				}
				return "", nil

			})
			if err != nil {
				c.String(http.StatusUnauthorized, "You're Unauthorized!")
			}
			if token.Valid {
				next(c)
			}
		}
		return c.String(http.StatusUnauthorized, "You're Unauthorized due to No token in the header")
	}
}
