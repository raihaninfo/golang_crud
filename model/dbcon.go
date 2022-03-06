package model

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB
var err error

func Dbcon() {
	db, err = sql.Open("sqlite3", "crud.db")
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxOpenConns(1)
	log.Println("db connection successful")
}
