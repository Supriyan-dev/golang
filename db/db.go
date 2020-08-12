package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", "root:@/kasumi_development")

	if err != nil {
		log.Fatal(err)
	}

	return
}
