package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {


	//db, err := sql.Open("mysql", "root:@/kasumi_development")
	db, err := sql.Open("mysql", "godx1:G0LangDX_1@tcp(mysql_lara:3306)/kasumi_development")

	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxIdleConns(5)
	return db
}
