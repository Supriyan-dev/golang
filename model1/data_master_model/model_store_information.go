package model1

import (
	"log"
	"strconv"
	"fmt"
	"../../db"
	"../../initialize"
	"../../models"
)

type Models_init models.DB_init

func (model1 Models_init) ReturnAllStoreInformationModel() (arrStoreInformation []initialize.StoreInformation, err error) {
	var storeInformation initialize.StoreInformation
	db := db.Connect()
	rows, err := db.Query("SELECT id_code_store, code_store, store_name FROM store_information")
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	for rows.Next() {
		if err := rows.Scan(&storeInformation.Id_code_store, &storeInformation.Code_store, &storeInformation.Store_name); err != nil {
			log.Fatal(err.Error())

		} else {
			arrStoreInformation = append(arrStoreInformation, storeInformation)
		}
	}

	return arrStoreInformation, nil
}


func (model1 Models_init) SortDESCStoreInformationModel(Sort, Col string) (arrStoreInformation []initialize.StoreInformation, err error) {
	var storeInformation initialize.StoreInformation
	db := db.Connect()
	// ASC := "ASC"
	// DESC := "DESC"
	// code_store := "code_store"

	qtext := fmt.Sprintf("SELECT id_code_store, code_store, store_name FROM store_information ORDER BY code_store DESC LIMIT 5")
	rows, err := db.Query(qtext)

	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	for rows.Next() {
		if err := rows.Scan(&storeInformation.Id_code_store, &storeInformation.Code_store, &storeInformation.Store_name); err != nil {
			log.Fatal(err.Error())

		} else {
			arrStoreInformation = append(arrStoreInformation, storeInformation)
		}
	}

	return arrStoreInformation, nil
}

func (model1 Models_init) SortASCStoreInformationModel(Sort, Col string) (arrStoreInformation []initialize.StoreInformation, err error) {
	var storeInformation initialize.StoreInformation
	db := db.Connect()
	code_store := "code_store"

	qtext := fmt.Sprintf("SELECT id_code_store, code_store, store_name FROM store_information ORDER BY %s ASC", code_store)
	rows, err := db.Query(qtext)

	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	for rows.Next() {
		if err := rows.Scan(&storeInformation.Id_code_store, &storeInformation.Code_store, &storeInformation.Store_name); err != nil {
			log.Fatal(err.Error())

		} else {
			arrStoreInformation = append(arrStoreInformation, storeInformation)
		}
	}

	return arrStoreInformation, nil
}

func (model1 Models_init) SearchStoreInformationModels(Keyword string, page int ,limit int) (arrSearch []initialize.StoreInformation, err error, CheckData int) {
	var Search initialize.StoreInformation
	db := db.Connect()
	querylimit := ``
	if strconv.Itoa(page) == "" && strconv.Itoa(limit) == ""{
		querylimit = ``
	}else {
		pageacheck := strconv.Itoa((page-1)*limit)
		showadata := strconv.Itoa(limit)
		querylimit = ` LIMIT `+pageacheck+`,`+showadata
	}
	db.QueryRow("SELECT count(*) FROM store_information WHERE CONCAT_WS('', code_store, store_name) LIKE ?", "%" + Keyword + "%").Scan(&CheckData)
	queryT := `SELECT id_code_store, code_store, store_name FROM store_information WHERE CONCAT_WS('', code_store, store_name) LIKE ?` +querylimit

	rows, err := db.Query(queryT, "%" + Keyword + "%")

	if err != nil {
		log.Print(err)
	}

	defer db.Close()
	for rows.Next() {
		if err := rows.Scan(&Search.Id_code_store, &Search.Code_store, &Search.Store_name); err != nil {
			log.Fatal(err.Error())
		} else {
			arrSearch = append(arrSearch, Search)
		}
	}

	return arrSearch, nil, CheckData
}

func (model1 Models_init) ReturnFilterStoreInformationModel(Id_code_store string) (arrEmp []initialize.Filter, err error) {
	var emp initialize.Filter
	db := db.Connect()
	rows, err := db.Query(`SELECT si.id_code_store, bi.first_name,bi.last_name, bi.employee_code FROM store_information si
	INNER JOIN general_information gi ON si.id_code_store = gi.id_store_code
	INNER JOIN basic_information bi ON bi.id_basic_information = gi.id_basic_information WHERE id_code_store = ?`, Id_code_store)
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	for rows.Next() {
		if err := rows.Scan(&emp.Id_code_store, &emp.First_name, &emp.Last_name, &emp.Employee_number); err != nil {
			log.Fatal(err.Error())

		} else {
			arrEmp = append(arrEmp, emp)
		}
	}

	return arrEmp, nil
}

func (model1 Models_init) GetIdStoreInformation(Id_code_store string) (si []initialize.StoreInformation, err error) {
	var storeInformation initialize.StoreInformation
	db := db.Connect()
	result, errExcuteData := db.Query("SELECT id_code_store, code_store, store_name FROM store_information WHERE id_code_store = ?", Id_code_store)
	if errExcuteData != nil {
		log.Println(err.Error())
	}
	defer result.Close()
	for result.Next() {
		errExcuteData := result.Scan(&storeInformation.Id_code_store, &storeInformation.Code_store, &storeInformation.Store_name)
		if errExcuteData != nil {
			log.Println(err.Error())
		} else {
			si = append(si, storeInformation)
		}
	}

	return si, nil
}

func (model1 Models_init) InsertStoreInformation(init_insert *initialize.StoreInformation) (st []initialize.StoreInformation, condition string) {
	db := db.Connect()
	stmt, err := db.Prepare("INSERT INTO store_information (code_store,store_name,latitude,longitude) VALUES (?,?,?,?)")
	if err != nil {
		log.Println(err.Error())
	}
	defer db.Close()

	result, err := stmt.Exec(init_insert.Code_store, init_insert.Store_name, init_insert.Latitude, init_insert.Longitude)
	log.Println(result)

	storeInsert := initialize.StoreInformation{
		Latitude: init_insert.Latitude,
		Longitude: init_insert.Longitude,
		Code_store: init_insert.Code_store,
		Store_name: init_insert.Store_name,
	}

	st = append(st, storeInsert)

	return st, "Success Response"
}

func (model1 Models_init) UpdateStoreInformation(Update *initialize.StoreInformation) (arrUpdate []initialize.StoreInformation, condition string) {
	db := db.Connect()

	stmt, err := db.Prepare("UPDATE store_information SET code_store = ?, store_name = ? WHERE id_code_store = ?")
	if err != nil {
		log.Println(err.Error())
	}

	result, err := stmt.Exec(Update.Code_store, Update.Store_name, Update.Id_code_store)
	log.Println(result)

	storeUpdate := initialize.StoreInformation{
		Id_code_store: Update.Id_code_store,
		Code_store:    Update.Code_store,
		Store_name:    Update.Store_name,
	}

	arrUpdate = append(arrUpdate, storeUpdate)

	return arrUpdate, "Success Response"

}

func (model1 Models_init) DeleteDataStoreInformation(Id_code_store string) (arrDelete []initialize.StoreInformation, err error) {
	var delete initialize.StoreInformation
	db := db.Connect()
	stmt, err := db.Prepare("DELETE FROM store_information WHERE id_code_store = ?")
	if err != nil {
		log.Println(err.Error())
	}

	stmt.Exec(Id_code_store)
	if err != nil {
		log.Println(err.Error())
	}

	Excute := initialize.StoreInformation{
		Id_code_store: delete.Id_code_store,
	}

	arrDelete = append(arrDelete, Excute)

	return arrDelete, nil
}
