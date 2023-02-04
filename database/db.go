package database

import (
	"database/sql"
	"fmt"
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
