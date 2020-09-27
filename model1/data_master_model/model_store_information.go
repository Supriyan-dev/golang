package model1

import (
	"log"

	"../../db"
	"../../initialize"
	"../../models"
)

type Models_init models.DB_init

func (model1 Models_init) ReturnAllStoreInformationModel() (arrStoreInformation []initialize.StoreInformation, err error) {
	var storeInformation initialize.StoreInformation
	db := db.Connect()
	rows, err := db.Query("SELECT id_code_store, code_store, store_name, latitude, longitude FROM store_information")
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	for rows.Next() {
		if err := rows.Scan(&storeInformation.Id_code_store, &storeInformation.Code_store, &storeInformation.Store_name, &storeInformation.Latitude, &storeInformation.Longitude); err != nil {
			log.Fatal(err.Error())

		} else {
			arrStoreInformation = append(arrStoreInformation, storeInformation)
		}
	}

	return arrStoreInformation, nil
}

func (model1 Models_init) SearchStoreInformationModels(Keyword string) (arrJoin []initialize.StoreInformation, err error) {
	var storeInformation initialize.StoreInformation
	db := db.Connect()
	result, err := db.Query(`SELECT id_code_store, code_store, store_name, latitude, longitude WHERE CONCAT_WS('', code_store, store_name, latitude, longitude) LIKE ?`, `%` + Keyword + `%`)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(result)
	defer db.Close()
	for result.Next() {
		if err := result.Scan(&storeInformation.Id_code_store, &storeInformation.Code_store, &storeInformation.Store_name, &storeInformation.Latitude, &storeInformation.Longitude); err != nil {
			log.Fatal(err.Error())
		} else {
			arrJoin = append(arrJoin, storeInformation)
		}
	}

	return arrJoin, nil
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
	result, errExcuteData := db.Query("SELECT id_code_store, code_store, store_name, latitude, longitude FROM store_information WHERE id_code_store = ?", Id_code_store)
	if errExcuteData != nil {
		log.Println(err.Error())
	}
	defer result.Close()
	for result.Next() {
		errExcuteData := result.Scan(&storeInformation.Id_code_store, &storeInformation.Code_store, &storeInformation.Store_name, &storeInformation.Latitude, &storeInformation.Longitude)
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
	stmt, err := db.Prepare("INSERT INTO store_information (code_store,store_name) VALUES (?,?)")
	if err != nil {
		log.Println(err.Error())
	}
	defer db.Close()

	result, err := stmt.Exec(init_insert.Code_store, init_insert.Store_name)
	log.Println(result)

	storeInsert := initialize.StoreInformation{
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
