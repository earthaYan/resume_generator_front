package db

import (
	"database/sql"
	"fmt"
)

const (
	dbName   = "resume"
	user     = "root"
	password = "123"
)

var db *sql.DB

func ConnectDb() *sql.DB {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", user, password, dbName))
	if err != nil {
		panic(err)
	}
	return db
}
func GetDB() *sql.DB {
	if db == nil {
		db = ConnectDb()
	}
	return db
}
func CloseDb() {
	if db != nil {
		db.Close()
	}
}
