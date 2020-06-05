package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	DBConn *sql.DB
	DBErr  error
)

func init() {
	DBConn, DBErr = sql.Open("mysql", "root:123!@#@tcp(localhost:3306)/video_server?charset=utf8")
	if DBErr != nil {
		panic(DBErr.Error())
	}
}
