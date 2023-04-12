package model

import (
	"errors"

	"github.com/labstack/echo/v4"
)

type User struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
	Token    string
	URLs     []URL
}

func NewUser(username, password string) (*User, error) {
	if len(password) == 0 || len(username) == 0 {
		return nil, errors.New("username or password can not be empty")
	}
	return &User{Username: username, Password: password}, nil
}

func (u *User) Bind(c echo.Context) (*User, error) {
	if err := c.Bind(u); err != nil {
		return nil, err
	}
	return u, nil
}
