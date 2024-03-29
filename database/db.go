package database

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	model "github.com/Calgorr/IE_Backend_Fall/Model"
	_ "github.com/lib/pq"
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
	err := db.Ping()
	if err != nil {
		panic(err)
	}
}

func AddUser(user *model.User) error {
	fmt.Println(user, "moz")
	connect()
	defer db.Close()
	sqlStatement := "INSERT INTO users (created_at,username,password) VALUES ($1,$2,$3)"
	_, err := db.Exec(sqlStatement, time.Now(), user.Username, user.Password)
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
	var lsInsertID int
	sqlStatement := "INSERT INTO urls (created_at,user_id,address,treshold,failed_times,warning) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id"
	err := db.QueryRow(sqlStatement, time.Now(), url.UserID, url.Address, url.Treshold, url.FailedTimes, 0).Scan(&lsInsertID)
	url.ID = int64(lsInsertID)
	if err != nil {
		return err
	}
	return err
}

func GetURLByUser(user_id int) ([]*model.URL, error) {
	connect()
	defer db.Close()
	urls := make([]*model.URL, 20)
	sqlStatement := "SELECT id,user_id, address, treshold, failed_times FROM urls WHERE user_id=$1"
	rows, err := db.Query(sqlStatement, user_id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		url := new(model.URL)
		err = rows.Scan(&url.ID, &url.UserID, &url.Address, &url.Treshold, &url.FailedTimes)
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
func GetURLByID(id int) (*model.URL, error) {
	url := new(model.URL)
	connect()
	defer db.Close()
	sqlStatement := "SELECT id,user_id, address, treshold, failed_times FROM urls WHERE id=$1"
	err := db.QueryRow(sqlStatement, id).Scan(&url.ID, &url.UserID, &url.Address, &url.Treshold, &url.FailedTimes)
	return url, err
}

func GetIDByUsername(username string) (int, error) {
	connect()
	defer db.Close()
	sqlStatement := "SELECT userid FROM users WHERE username=$1"
	row := db.QueryRow(sqlStatement, username)
	var id int
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func GetAllURLs() ([]*model.URL, error) {
	connect()
	defer db.Close()
	urls := make([]*model.URL, 100)
	sqlstatment := "SELECT * FROM urls"
	rows, err := db.Query(sqlstatment)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		url := new(model.URL)
		err := rows.Scan(url.UserID, url.Address, url.Treshold, url.FailedTimes)
		if err != nil {
			return nil, err
		}
		urls = append(urls, url)
	}
	return urls, nil
}
func GetAlertsByUserID(ID int) ([]*model.URL, error) {
	var urls []*model.URL
	connect()
	defer db.Close()
	sqlstatement := "SELECT * FROM urls WHERE user_id=$1 AND warning=1"
	rows, err := db.Query(sqlstatement, ID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		url := new(model.URL)
		err = rows.Scan(&url.ID, url.UserID, url.Address, url.Treshold, url.FailedTimes)
		if err != nil {
			return nil, err
		}
		urls = append(urls, url)
	}
	return urls, nil
}

func IncrementFailedByOne(url *model.URL) error {
	connect()
	defer db.Close()
	sqlstatement := "UPDATE failed_times FROM urls SET failed_times=failed_times+1 WHERE id=$1"
	_, err := db.Exec(sqlstatement, url.ID)
	return err
}

func AddRequest(request *model.Request) error {
	connect()
	defer db.Close()
	sqlstatement := "INSERT INTO request (created_at,url_id,result) VALUES($1,$2$3)"
	_, err := db.Exec(sqlstatement, time.Now().Unix(), request.URLID, request.StatusCode)
	return err

}

func SetWarning(url model.URL) error {
	connect()
	defer db.Close()
	sqlstatement := "UPDATE warning FROM urls warning=1 WHERE id=$1"
	_, err := db.Exec(sqlstatement, url.ID)
	return err
}

func TresholdReached(url model.URL) bool {
	connect()
	defer db.Close()
	sqlstatement := "SELECT treshold, failed_times FROM urls WHERE id=$1"
	row := db.QueryRow(sqlstatement, url.ID)
	var treshold, failed_times int
	err := row.Scan(treshold, failed_times)
	if err == sql.ErrNoRows {
		panic("url does not exists")
	} else if failed_times > treshold {
		return false
	}
	return true
}
