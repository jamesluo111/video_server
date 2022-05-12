package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbCoon *sql.DB
	err    error
)

func init() {
	dbCoon, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/video?charset=utf8")
	if err != nil {
		panic(err.Error())
	}

	err = dbCoon.Ping()
	if err != nil {
		panic(err.Error())
	}
}
