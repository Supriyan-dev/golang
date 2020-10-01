package model1

import (
	"log"
	"strconv"
	"../../db"
	"../../initialize"
	"../../models"
)

type ModelSection_init models.DB_init

func (model1 ModelSection_init) ReturnAllStoreSectionInformation() (arrAll []initialize.StoreSectionInformation, err error) {
	var all initialize.StoreSectionInformation
	db := db.Connect()

	rows, err := db.Query("SELECT id_store_section, store_section_code, store_section_name FROM store_section_information")
	if err != nil {
		log.Println(err.Error())
	}
	defer db.Close()

	for rows.Next() {
		if err := rows.Scan(&all.Id_store_section, &all.Store_section_code, &all.Store_section_name); err != nil {
			log.Println(err.Error())
		} else {
			arrAll = append(arrAll, all)
		}
	}

	return arrAll, nil
}

func (model1 ModelSection_init) SearchStoreSectionInformationModels(Keyword string, page int ,limit int) (arrSearch []initialize.StoreSectionInformation, err error, CheckData int) {
	var Search initialize.StoreSectionInformation
	db := db.Connect()
	querylimit := ``
	if strconv.Itoa(page) == "" && strconv.Itoa(limit) == ""{
		querylimit = ``
	}else {
		pageacheck := strconv.Itoa((page-1)*limit)
		showadata := strconv.Itoa(limit)
		querylimit = ` LIMIT `+pageacheck+`,`+showadata
	}
	db.QueryRow("FROM store_section_information WHERE CONCAT_WS('', store_section_code, store_section_name) LIKE ?", "%" + Keyword + "%").Scan(&CheckData)
	queryT := `SELECT id_store_section, store_section_code, store_section_name FROM store_section_information WHERE CONCAT_WS('', store_section_code, store_section_name) LIKE ?` +querylimit

	rows, err := db.Query(queryT, "%" + Keyword + "%")

	if err != nil {
		log.Print(err)
	}

	defer db.Close()
	for rows.Next() {
		if err := rows.Scan(&Search.Id_store_section, &Search.Store_section_code, &Search.Store_section_name); err != nil {
			log.Fatal(err.Error())
		} else {
			arrSearch = append(arrSearch, Search)
		}
	}

	return arrSearch, nil, CheckData
}

func (model1 ModelSection_init) GetDataStoreSectionInformation(Id_store_section string) (arrGet []initialize.StoreSectionInformation, err error) {
	var get initialize.StoreSectionInformation
	db := db.Connect()

	result, err := db.Query("SELECT id_store_section, store_section_code, store_section_name FROM store_section_information WHERE id_store_section = ?", Id_store_section)
	if err != nil {
		log.Println(err.Error())
	}
	defer result.Close()
	for result.Next() {

		err := result.Scan(&get.Id_store_section, &get.Store_section_code, &get.Store_section_name)
		if err != nil {
			log.Println(err.Error())
		} else {
			arrGet = append(arrGet, get)
		}
	}

	return arrGet, nil
}

func (model1 ModelSection_init) InsertDataStoreSectionInformation(insert *initialize.StoreSectionInformation) (arrInsert []initialize.StoreSectionInformation, err error) {
	db := db.Connect()
	stmt, err := db.Prepare("INSERT INTO store_section_information (store_section_code,store_section_name) VALUES (?,?)")
	if err != nil {
		log.Println(err.Error())
	}
	defer db.Close()
	stmt.Exec(insert.Store_section_code, insert.Store_section_name)
	if err != nil {
		log.Println(err.Error())
	}

	Execute := initialize.StoreSectionInformation{
		Store_section_code: insert.Store_section_code,
		Store_section_name: insert.Store_section_name,
	}

	arrInsert = append(arrInsert, Execute)

	return arrInsert, nil
}

func (model1 ModelSection_init) UpdateDataStoreSectionInformation(update *initialize.StoreSectionInformation) (arrUpdate []initialize.StoreSectionInformation, err error) {

	db := db.Connect()

	stmt, err := db.Prepare("UPDATE store_section_information SET store_section_code = ?, store_section_name = ? WHERE id_store_section = ?")
	if err != nil {
		log.Println(err.Error())
	}

	stmt.Exec(update.Store_section_code, update.Store_section_name, update.Id_store_section)
	if err != nil {
		log.Println(err.Error())
	}

	Execute := initialize.StoreSectionInformation{
		Id_store_section:   update.Id_store_section,
		Store_section_code: update.Store_section_code,
		Store_section_name: update.Store_section_name,
	}

	arrUpdate = append(arrUpdate, Execute)

	return arrUpdate, nil
}

func (model1 ModelSection_init) DeleteDataStoreSectionInformation(delete *initialize.StoreSectionInformation) (arrDelete []initialize.StoreSectionInformation, err error) {

	db := db.Connect()
	stmt, err := db.Prepare("DELETE FROM store_section_information WHERE id_store_section = ?")
	if err != nil {
		log.Println(err.Error())
	}

	stmt.Exec(delete.Id_store_section)
	if err != nil {
		log.Println(err.Error())
	}

	Execute := initialize.StoreSectionInformation{
		Id_store_section: delete.Id_store_section,
	}

	arrDelete = append(arrDelete, Execute)

	return arrDelete, nil
}
