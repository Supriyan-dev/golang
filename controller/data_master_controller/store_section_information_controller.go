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

func ReturnAllStroreSectionInformation(w http.ResponseWriter, r *http.Request) {
	var store initialize.StoreSectionInformation
	var arrStoreSectionInformation []initialize.StoreSectionInformation
	var response initialize.Response

	db, err := db.Connect()

	rows, err := db.Query("SELECT * FROM store_section_information")

	if err != nil {
		log.Print(err)
	}
	defer db.Close()

	for rows.Next() {
		// if err := rows.Scan(&prefect.Id_prefecture, &prefect.ISO, &prefect.Prefecture_name); err != nil {
		if err := rows.Scan(&store.Id_store_section, &store.Store_section_code, &store.Store_section_name); err != nil {

			log.Fatal(err.Error())

		} else {
			arrStoreSectionInformation = append(arrStoreSectionInformation, store)
		}
	}
	response.Status = 200
	response.Message = "Success"
	response.Data = arrStoreSectionInformation

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func ReturnAllStroreSectionInformationPagination(w http.ResponseWriter, r *http.Request) {
	var store initialize.StoreSectionInformation
	var arrStoreSectionInformation []initialize.StoreSectionInformation
	var response initialize.Response
	code := mux.Vars(r)
	fmt.Fprintf(w, "Category: %v\n", code["page"])

	db, err := db.Connect()
	rows, err := db.Query("SELECT * FROM store_section_information ORDER BY id_store_section LIMIT " + code["page"] + " OFFSET 0")
	if err != nil {
		log.Print(err)
	}
	defer db.Close()

	for rows.Next() {
		if err := rows.Scan(&store.Id_store_section, &store.Store_section_code, &store.Store_section_name); err != nil {

			log.Fatal(err.Error())

		} else {
			arrStoreSectionInformation = append(arrStoreSectionInformation, store)
		}
	}
	response.Status = 200
	response.Message = "Success"
	response.Data = arrStoreSectionInformation

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
