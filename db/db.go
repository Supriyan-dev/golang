package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	var err error
	db, err := sql.Open("mysql", "root:@/kasumi_development")

	if err != nil {
		log.Fatal(err)
	}

	return db
}
