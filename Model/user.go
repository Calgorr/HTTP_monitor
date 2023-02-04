package model

import "errors"

type User struct {
	Username, Password string
}

func NewUser(username, password string) (*User, error) {
	if len(username) == 0 || len(password) == 0 {
		return nil, errors.New("username or password can not be empty")
	}
	return &User{Username: username, Password: password}, nil
}
