package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connect(addr string) *sql.DB {

	db, err := sql.Open("mysql", addr)

	if err != nil {
		log.Fatal(err)
	}

	return db

}
