package model

import (
	"net/http"
	_url "net/url"
	"strings"
)

type URL struct {
	ID          int
	UserID      int
	Address     string
	Threshold   int
	FailedTimes int
	Requests    []Request
}

func NewURL(userID, threshold, failedTimes int, address string) (*URL, error) {
	url := new(URL)
	url.UserID = userID
	url.Threshold = threshold
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
	return u.FailedTimes >= u.Threshold
}

func (u *URL) SendRequest() (*Request, error) {
	resp, err := http.Get(u.Address)
	req := new(Request)
	req.URLID = u.ID
	if err != nil {
		return nil, err
	}
	req.StatusCode = resp.StatusCode
	return req, nil
}
