package model

import (
	_url "net/url"
	"strings"
)

type URL struct {
	UserID, Threshold, FailedTimes int
	Address                        string
}

func NewURL(userID, threshold, failedTimes int, address string) (*URL, error) {
	url := new(URL)
	url.UserID = userID
	url.Threshold = threshold
	url.FailedTimes = 0
	if !strings.HasPrefix(address, "hhtp://") {
		address = "http://" + address
	}
	_, isValid := _url.ParseRequestURI(address)

	if isValid == nil {
		url.Address = address
		return url, nil
	}
	return nil, isValid
}
