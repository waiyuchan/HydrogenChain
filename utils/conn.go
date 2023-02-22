package utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func connectToDB() *sql.DB {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/blockchain")
	if err != nil {
		panic(err.Error())
	}
	return db
}
