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

func ReturnAllUnitInformation(w http.ResponseWriter, r *http.Request) {
	var unit initialize.UnitInformation
	var arrUnitInformation []initialize.UnitInformation
	var response initialize.Response

	db, err := db.Connect()

	rows, err := db.Query("SELECT * FROM unit_information")

	if err != nil {
		log.Print(err)
	}
	defer db.Close()

	for rows.Next() {
		if err := rows.Scan(&unit.Id_unit, &unit.Unit_code, &unit.Unit_name); err != nil {

			log.Fatal(err.Error())

		} else {
			arrUnitInformation = append(arrUnitInformation, unit)
		}
	}
	response.Status = 200
	response.Message = "Success"
	response.Data = arrUnitInformation

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func ReturnAllUnitInformationPagination(w http.ResponseWriter, r *http.Request) {
	var unit initialize.UnitInformation
	var arrUnitInformation []initialize.UnitInformation
	var response initialize.Response

	code := mux.Vars(r)
	fmt.Fprintf(w, "Category: %v\n", code["page"])

	db, err := db.Connect()
	rows, err := db.Query("SELECT * FROM unit_information ORDER BY id_unit LIMIT " + code["page"] + " OFFSET 0")
	if err != nil {
		log.Print(err)
	}
	defer db.Close()

	for rows.Next() {
		if err := rows.Scan(&unit.Id_unit, &unit.Unit_code, &unit.Unit_name); err != nil {

			log.Fatal(err.Error())

		} else {
			arrUnitInformation = append(arrUnitInformation, unit)
		}
	}
	response.Status = 200
	response.Message = "Success"
	response.Data = arrUnitInformation

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
