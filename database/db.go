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
	connect()
	defer db.Close()
	sqlStatement := "INSERT INTO users (created_at,username,password) VALUES ($1,$2,$3)"
	_, err := db.Exec(sqlStatement, time.Now().Unix(), user.Username, user.Password)
	return err
}
func GetUserByUsername(username string) (*model.User, error) {
	connect()
	defer db.Close()
	sqlStatement := "SELECT username, password FROM users WHERE username=$1 "
	row := db.QueryRow(sqlStatement, username)
	user := new(model.User)
	err := row.Scan(&user.Username, &user.Password)
	if err == sql.ErrNoRows {
		return nil, errors.New("User does not exists")
	}
	return user, nil
}

func AddURL(url *model.URL) error {
	connect()
	defer db.Close()
	sqlStatement := "INSERT INTO url (created_at,user_id,address,threshold,failed_times) VALUES ($1,$2,$3,$4,$5)"
	_, err := db.Exec(sqlStatement, time.Now().Unix(), url.UserID, url.Address, url.Threshold, url.FailedTimes)
	return err
}

func GetURLByUser(user_id int) ([]*model.URL, error) {
	connect()
	defer db.Close()
	urls := make([]*model.URL, 20)
	sqlStatement := "SELECT user_id, address, threshold, failed_times FROM url WHERE user_id=$1"
	rows, err := db.Query(sqlStatement, user_id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		url := new(model.URL)
		err = rows.Scan(&url.UserID, &url.Address, &url.Threshold, &url.FailedTimes)
		if err != nil {
			panic(err)
		}
		urls = append(urls, url)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return urls, nil
}

func GetAllURLs() []*model.URL {
	connect()
	defer db.Close()
	urls := make([]*model.URL, 100)
	sqlstatment := "SELECT * FROM url"
	rows, err := db.Query(sqlstatment)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		url := new(model.URL)
		rows.Scan(url.UserID, url.Address, url.Threshold, url.FailedTimes)
		urls = append(urls, url)
	}
	return urls
}

func IncrementFailedByOne(url *model.URL) error {
	connect()
	defer db.Close()
	sqlstatement := "UPDATE failed_times FROM url SET failed_times=failed_times+1 WHERE url_id=$1"
	_, err := db.Exec(sqlstatement, url.URLID)
	return err
}
