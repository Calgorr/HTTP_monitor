package model

import "errors"

type User struct {
	Username, Password string
	URLs               []URL
}

func NewUser(username, password string) (*User, error) {
	if len(password) == 0 || len(username) == 0 {
		return nil, errors.New("username or password can not be empty")
	}
	return &User{Username: username, Password: password}, nil
}
