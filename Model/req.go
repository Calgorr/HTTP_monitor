package model

import "github.com/labstack/echo/v4"

type Request struct {
	URLID, StatusCode int
}

func (u *Request) Bind(c echo.Context) (*Request, error) {
	if err := c.Bind(u); err != nil {
		return nil, err
	}
	return u, nil
}
