package model1

import (
	"log"

	"../../db"
	"../../initialize/driver_management"
	"../../models"
)

type Models_init models.DB_init

func (model1 Models_init) ReturnAllStoreInformationModel() (arrStoreInformation []initialize.StoreInformationDriver, err error) {
	var storeInformation initialize.StoreInformationDriver
	db := db.Connect()
	rows, err := db.Query("SELECT id_code_store, code_store, store_name, latitude, longitude FROM store_information")
	if err != nil {
		log.Print(err)
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
func (model1 Models_init) SearchStoreInformationDriver(code_store, store_name string) (arrSearch []initialize.StoreInformationDriver, err error) {
	var Search *initialize.StoreInformationDriver
	// var code_store string
	// var store_name string
	db := db.Connect()
	rows, err := db.Query("SELECT id_code_store, code_store, store_name, latitude, longitude FROM store_information LIKE = ?", code_store + "%", store_name + "%")
	if err != nil {
		log.Print(err)
	}
	defer db.Close()
	for rows.Next() {
		if err := rows.Scan(&Search.Id_code_store, &Search.Code_store, &Search.Store_name, &Search.Latitude, &Search.Longitude); err != nil {
			log.Fatal(err.Error())

		} else {
			// arrSearch = append(arrSearch, Search)
		}
	}

	return arrSearch, nil
}


func (model1 Models_init) ReturnFilterStoreInformationModel(Id_code_store string) (arrEmp []initialize.Filter, err error) {
	var emp initialize.Filter
	db := db.Connect()
	rows, err := db.Query(`SELECT si.id_code_store, bi.first_name,bi.last_name, bi.employee_code FROM store_information si
	INNER JOIN general_information gi ON si.id_code_store = gi.id_store_code
	INNER JOIN basic_information bi ON bi.id_basic_information = gi.id_basic_information WHERE id_code_store = ?`, Id_code_store)
	if err != nil {
		log.Print(err)
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

func (model1 Models_init) GetIdStoreInformation(Id_code_store string) (si []initialize.StoreInformationDriver, err error) {
	var storeInformation initialize.StoreInformationDriver
	db := db.Connect()
	result, errExcuteData := db.Query("SELECT id_code_store, code_store, store_name, latitude, longitude FROM store_information WHERE id_code_store = ?", Id_code_store)
	if errExcuteData != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		errExcuteData := result.Scan( &storeInformation.Id_code_store, &storeInformation.Code_store, &storeInformation.Store_name, &storeInformation.Latitude, &storeInformation.Longitude)
		if errExcuteData != nil {
			panic(err.Error())
		} else {
			si = append(si, storeInformation)
		}
	}

	return si, nil
}

// func (model1 Models_init) InsertStoreInformation(init_insert *initialize.StoreInformationDriver) (st []initialize.StoreInformationDriver, condition string) {
// 	db := db.Connect()
// 	stmt, err := db.Prepare("INSERT INTO store_information (code_store,store_name,latitude,longitude) VALUES (?,?,?,?)")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer db.Close()

// 	result, err := stmt.Exec(init_insert.Code_store, init_insert.Store_name, init_insert.Latitude, init_insert.Longitude)
// 	log.Println(result)

// 	storeInsert := initialize.StoreInformationDriver{
// 		Code_store: init_insert.Code_store,
// 		Store_name: init_insert.Store_name,
// 		Latitude: init_insert.Latitude,
// 		Longitude: init_insert.Longitude,
// 	}

// 	st = append(st, storeInsert)

// 	return st, "Success Response"
// }

// func (model1 Models_init) UpdateStoreInformation(Update *initialize.StoreInformationDriver) (arrUpdate []initialize.StoreInformationDriver, condition string) {
// 	db := db.Connect()

// 	stmt, err := db.Prepare("UPDATE store_information SET store_name = ?, code_store = ?, latitude = ?, longitude = ? WHERE id_code_store = ? ")
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	result, err := stmt.Exec(Update.Code_store, Update.Store_name, Update.Latitude, Update.Longitude,Update.Id_code_store)
// 	log.Println(result)

// 	storeUpdate := initialize.StoreInformationDriver{
// 		Id_code_store: Update.Id_code_store,
// 		Code_store:    Update.Code_store,
// 		Store_name:    Update.Store_name,
// 		Latitude:    Update.Latitude,
// 		Longitude:    Update.Longitude,
// 	}

// 	arrUpdate = append(arrUpdate, storeUpdate)

// 	return arrUpdate, "Success Response"

// }

// func (model1 Models_init) DeleteDataStoreInformation(Id_code_store string) (arrDelete []initialize.StoreInformationDriver, err error) {
// 	var delete initialize.StoreInformationDriver
// 	db := db.Connect()
// 	stmt, err := db.Prepare("DELETE FROM store_information WHERE id_code_store = ?")
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	stmt.Exec(Id_code_store)
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	Excute := initialize.StoreInformationDriver{
// 		Id_code_store: delete.Id_code_store,
// 	}

// 	arrDelete = append(arrDelete, Excute)

// 	return arrDelete, nil
// }
