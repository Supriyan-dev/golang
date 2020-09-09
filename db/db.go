package db

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {

	//db, err := sql.Open("mysql", "root:@/kasumi_development")
	db, err := sql.Open("mysql", "godx1:G0LangDX_1@tcp(mysql_lara:3306)/kasumi_development")

	if err != nil {
		log.Fatal(err)
	}
	//db.SetMaxIdleConns(10)
	//db.SetConnMaxLifetime(0)
	//db.SetMaxIdleConns(50)
	//db.SetMaxOpenConns(50)
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(0)
	db.SetConnMaxLifetime(time.Nanosecond)
	return db
}
