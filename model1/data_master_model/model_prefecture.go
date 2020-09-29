package model1

import (
	"log"

	"../../db"
	"../../initialize"
	"../../models"
)

type ModelPref_init models.DB_init

func (model1 ModelPref_init) ReturnAllDataPrefecture() (arrAll []initialize.Prefect, err error) {
	var all initialize.Prefect
	db := db.Connect()

	rows, err := db.Query("SELECT id_prefecture, ISO, prefecture_name FROM prefecture")
	if err != nil {
		log.Print(err)
	}
	defer db.Close()

	for rows.Next() {
		if err := rows.Scan(&all.Id_prefecture, &all.ISO, &all.Prefecture_name); err != nil {
			log.Println(err.Error())

		} else {
			arrAll = append(arrAll, all)
		}
	}

	return arrAll, nil
}

func (model1 ModelPref_init) SearchPrefectureModels(Keyword string) (arrJoin []initialize.Prefect, err error) {
	var all initialize.Prefect
	db := db.Connect()
	result, err := db.Query(`SELECT id_prefecture, ISO, prefecture_name FROM prefecture WHERE CONCAT_WS('', ISO, prefecture_name) LIKE ?`, `%` + Keyword + `%`)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(result)
	defer db.Close()
	for result.Next() {
		if err := result.Scan(&all.Id_prefecture, &all.ISO, &all.Prefecture_name); err != nil {
			log.Fatal(err.Error())
		} else {
			arrJoin = append(arrJoin, all)
		}
	}

	return arrJoin, nil
}

func (model1 ModelPref_init) GetDataPrefecture(Id_prefecture string) (arrGet []initialize.Prefect, err error) {
	var get initialize.Prefect
	db := db.Connect()

	result, err := db.Query("SELECT id_prefecture, ISO, prefecture_name FROM prefecture WHERE id_prefecture = ?", Id_prefecture)
	if err != nil {
		log.Println(err.Error())
	}
	defer result.Close()
	for result.Next() {

		err := result.Scan(&get.Id_prefecture, &get.ISO, &get.Prefecture_name)
		if err != nil {
			log.Println(err.Error())
		} else {
			arrGet = append(arrGet, get)
		}
	}
	return arrGet, nil
}

func (model1 ModelPref_init) InsertDataPrefecture(insert *initialize.Prefect) (arrInsert []initialize.Prefect, err error) {

	db := db.Connect()
	stmt, err := db.Prepare("INSERT INTO prefecture (ISO, prefecture_name) VALUES(?,?)")
	if err != nil {
		log.Println(err.Error())
	}
	defer db.Close()

	stmt.Exec(insert.ISO, insert.Prefecture_name)
	if err != nil {
		log.Println(err.Error())
	}

	Execute := initialize.Prefect{
		ISO:             insert.ISO,
		Prefecture_name: insert.Prefecture_name,
	}

	arrInsert = append(arrInsert, Execute)

	return arrInsert, nil
}

func (model1 ModelPref_init) UpdateDataprefecture(update *initialize.Prefect) (arrUpdate []initialize.Prefect, err error) {
	db := db.Connect()

	stmt, err := db.Prepare("UPDATE prefecture SET ISO = ?, prefecture_name = ? WHERE id_prefecture = ?")
	if err != nil {
		log.Println(err.Error())
	}

	stmt.Exec(update.ISO, update.Prefecture_name, update.Id_prefecture)
	if err != nil {
		log.Println(err.Error())
	}

	Execute := initialize.Prefect{
		Id_prefecture:   update.Id_prefecture,
		ISO:             update.ISO,
		Prefecture_name: update.Prefecture_name,
	}

	arrUpdate = append(arrUpdate, Execute)

	return arrUpdate, nil
}

func (model1 ModelPref_init) DeleteDataPrefecture(delete *initialize.Prefect) (arrDelete []initialize.Prefect, err error) {
	db := db.Connect()
	stmt, err := db.Prepare("DELETE FROM prefecture WHERE id_prefecture = ?")
	if err != nil {
		log.Println(err.Error())
	}

	stmt.Exec(delete.Id_prefecture)
	if err != nil {
		log.Println(err.Error())
	}

	Execute := initialize.Prefect{
		Id_prefecture: delete.Id_prefecture,
	}

	arrDelete = append(arrDelete, Execute)

	return arrDelete, nil

}
