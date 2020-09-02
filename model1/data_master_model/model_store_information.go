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
	rows, err := db.Query("SELECT * FROM store_information")
	if err != nil {
		log.Print(err)
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

func (model1 Models_init) GetIdStoreInformation(Id_code_store string) (si []initialize.StoreInformation, err error) {
	var storeInformation initialize.StoreInformation
	db := db.Connect()
	result, errExcuteData := db.Query("SELECT id_code_store, code_store, store_name FROM store_information WHERE id_code_store = ?", Id_code_store)
	if errExcuteData != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		errExcuteData := result.Scan(&storeInformation.Id_code_store, &storeInformation.Code_store, &storeInformation.Store_name)
		if errExcuteData != nil {
			panic(err.Error())
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
		panic(err.Error())
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
		panic(err.Error())
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

func (model1 Models_init) DeleteDataStoreInformation(delete *initialize.StoreInformation) (arrDelete []initialize.StoreInformation, err error) {
	db := db.Connect()
	stmt, err := db.Prepare("DELETE FROM store_information WHERE id_code_store = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(delete.Id_code_store)
	if err != nil {
		panic(err.Error())
	}

	Excute := initialize.StoreInformation{
		Id_code_store: delete.Id_code_store,
	}

	arrDelete = append(arrDelete, Excute)

	return arrDelete, nil
}
