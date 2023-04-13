package model

import (
	"errors"
	"net/http"
	_url "net/url"
	"strings"

	"github.com/labstack/echo/v4"
)

type URL struct {
	ID          int64
	UserID      int    `form:"userID" json:"userID"`
	Address     string `form:"address" json:"address"`
	Treshold    int    `form:"treshold" json:"treshold"`
	FailedTimes int    `form:"failedTimes" json:"failedTimes"`
	Requests    []Request
}

func NewURL(userID, treshold, failedTimes int, address string) (*URL, error) {
	url := new(URL)
	url.UserID = userID
	url.Treshold = treshold
	url.FailedTimes = 0
	if !strings.HasPrefix(address, "http://") {
		address = "http://" + address
	}
	_, isValid := _url.ParseRequestURI(address)

	if isValid == nil {
		url.Address = address
		return url, nil
	}
	return nil, isValid
}

func (u *URL) AlarmTrigger() bool {
	return u.FailedTimes >= u.Treshold
}

func (u *URL) SendRequest() (*Request, error) {
	resp, err := http.Get(u.Address)
	req := new(Request)
	req.URLID = u.ID
	if err != nil {
		return nil, err
	}
	req.StatusCode = int64(resp.StatusCode)
	return req, nil
}

func (u *URL) Bind(c echo.Context) (*URL, error) {
	if err := c.Bind(u); err != nil {
		return nil, err
	}
	if !strings.HasPrefix(u.Address, "http://") {
		u.Address = "http://" + u.Address
	}
	_, isValid := _url.ParseRequestURI(u.Address)

	if isValid == nil {
		return u, nil
	}
	return nil, errors.New("URL is not valid")
}
