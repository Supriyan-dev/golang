package utils

import (
	db2 "../db"
	"log"
)

func CheckCountDataById(sql string, id interface{} ) (CountData int) {
	db := db2.Connect()
	err := db.QueryRow(sql,id).Scan(&CountData)
	if err != nil {
		log.Println(err.Error())
	}
	return CountData
}
