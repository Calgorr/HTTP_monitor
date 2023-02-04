package model

import (
	URL "net/url"
	"strings"
)

type url struct {
	UserID, Threshold, FailedTimes int
	Address                        string
}

func NewURL(userID, threshold, failedTimes int, address string) (*url, error) {
	url := new(url)
	url.UserID = userID
	url.Threshold = threshold
	url.FailedTimes = 0
	if !strings.HasPrefix(address, "hhtp://") {
		address = "http://" + address
	}
	_, isValid := URL.ParseRequestURI(address)

	if isValid == nil {
		url.Address = address
		return url, nil
	}
	return nil, isValid
}
