package database

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	model "github.com/Calgorr/IE_Backend_Fall/Model"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "calgor"
	password = "ami1r3ali"
	dbname   = "http_monitor"
)

var (
	db  *sql.DB
	err error
)

func connect() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err := db.Ping()

	if err != nil {
		panic(err)
	}
}

func AddUser(user *model.User) error {
	sqlStatement := "INSERT INTO users (created_at,username,password) VALUES ($1,$2,$3)"
	_, err := db.Exec(sqlStatement, time.Now().Unix(), user.Username, user.Password)
	return err
}
func GetUserByUsername(username string) (*model.User, error) {
	sqlStatement := "SELECT username, password FROM users WHERE username=$1 "
	row := db.QueryRow(sqlStatement, username)
	if row.Scan() == sql.ErrNoRows {
		return nil, errors.New("User does not exists")
	}
	user := new(model.User)
	row.Scan(&user.Username, user.Password)
	return user, nil
}
