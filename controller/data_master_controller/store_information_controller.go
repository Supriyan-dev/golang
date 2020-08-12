package data_master_controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jeffri/golang-test/db"
	"github.com/jeffri/golang-test/initialize"
)

func ReturnAllStoreInformation(w http.ResponseWriter, r *http.Request) {
	var storeInformation initialize.StoreInformation
	var arrStoreInformation []initialize.StoreInformation
	var response initialize.Response

	db, err := db.Connect()

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
	response.Status = 200
	response.Message = "Success"
	response.Data = arrStoreInformation

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func ReturnAllStoreInformationPagination(w http.ResponseWriter, r *http.Request) {
	var storeInformation initialize.StoreInformation
	var arrStoreInformation []initialize.StoreInformation
	var response initialize.Response
	code := mux.Vars(r)
	fmt.Fprintf(w, "Store_information: %v\n", code["page"])

	db, err := db.Connect()
	rows, err := db.Query("SELECT * FROM store_information ORDER BY id_code_store LIMIT " + code["page"] + " OFFSET 20")
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
	response.Status = 200
	response.Message = "Success"
	response.Data = arrStoreInformation

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
