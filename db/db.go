package db

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {

	db, err := sql.Open("mysql", "root:@/kasumi_development")
	//db, err := sql.Open("mysql", "godx1:G0LangDX_1@tcp(mysql_lara:3306)/kasumi_development")

	if err != nil {
		log.Fatal(err)
	}
	//db.SetMaxIdleConns(10)
	//db.SetConnMaxLifetime(time.Second)
	//db.SetMaxIdleConns(50)
	//db.SetMaxOpenConns(50)
	//db.Ping()
	KillSleepConnection()
	return db
}

func KillSleepConnection()  {

	db, err := sql.Open("mysql", "root:@/kasumi_development")
	//db, err := sql.Open("mysql", "godx1:G0LangDX_1@tcp(mysql_lara:3306)/kasumi_development")

	if err != nil {
		log.Fatal(err)
	}
	//db.SetMaxIdleConns(10)
	var datakill string
	showsleepConnection, errshowSleepConnection := db.Query(`select id from information_schema.processlist where Command='Sleep' and USER ='godx1'`)

	if errshowSleepConnection != nil {
		//log.Println(errshowSleepConnection)
	}
	for showsleepConnection.Next() {
		showsleepConnection.Scan(&datakill)

		//log.Println(datakill)

		killExecute, errkillExecute := db.Exec(`kill ?`, datakill)

		if errkillExecute != nil {
			//log.Println(errkillExecute.Error())
		}
		_, errcheckExecute := killExecute.RowsAffected()
		if errcheckExecute != nil {
			//log.Println(errcheckExecute.Error())
		}
		//log.Println(checkExecute)

	}
}


