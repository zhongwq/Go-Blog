package utils

import (
	"database/sql"
)

var DB *sql.DB
var err error

func init() {
	DB, err = sql.Open("mysql", "root:limzhonglin@tcp(127.0.0.1:3306)/?charset=utf8")
	if err != nil {
		panic(err)
	}
}

func GetConn() *sql.DB {
	return DB;
}

