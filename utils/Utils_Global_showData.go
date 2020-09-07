package utils

import (
	db2 "../db"
	"log"
)

func GetDataByIdInt(sql string, id int ) (CountData int64) {
	db := db2.Connect()
	err := db.QueryRow(sql,id).Scan(&CountData)
	if err != nil {
		log.Println(err.Error())
	}
	defer db.Close()
	return CountData
}

func GetDataByIdFloat(sql string, id int ) (CountData float64) {
	db := db2.Connect()
	err := db.QueryRow(sql,id).Scan(&CountData)
	if err != nil {
		log.Println(err.Error())
	}
	defer db.Close()
	return CountData
}
