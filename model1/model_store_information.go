package model1

import (
	"../db"
	"../initialize"
	"../models"
)

type Models_init models.DB_init

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

// func (model models_init) InsertStoreInformation(storeS *initialize.StoreInformation) (st []initialize.StoreInformation, condition string) {
// 	db := db.Connect()
// 	stmt, err := db.Prepare("INSERT INTO store_information (code_store,store_name) VALUES (?,?)")
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	defer db.Close()
// 	execute, err1 := stmt.Exec(&storeS.Code_store, &storeS.Store_name)
// 	log.Print(execute)
// 	log.Print(err1)

// 	if err1 != nil {
// 		log.Print(err1)
// 		return nil, "missing required field in body request"
// 	}

// 	storeInsert := initialize.StoreInformation{
// 		Code_store: storeS.Code_store,
// 		Store_name: storeS.Store_name,
// 	}

// 	st = append(st, storeInsert)

// 	return st, "Success Response"
// }
