package data_master_controller

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jeffri/golang-test/GO_DX_SERVICES/db"

	"github.com/jeffri/golang-test/GO_DX_SERVICES/initialize"
)

func ReturnAllUnitInformation(w http.ResponseWriter, r *http.Request) {
	var unit initialize.UnitInformation
	var arrUnitInformation []initialize.UnitInformation
	var response initialize.Response

	db := db.Connect()

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

	db := db.Connect()
	defer db.Close()
	code := mux.Vars(r)

	totalDataPerPage, _ := strconv.Atoi(code["perPage"])
	page, _ := strconv.Atoi(code["page"])

	var totalData int
	err := db.QueryRow("SELECT COUNT(*) FROM unit_information").Scan(&totalData)

	totalPage := int(math.Ceil(float64(totalData) / float64(totalDataPerPage)))

	if page > totalPage {
		page = totalPage
	}
	if page <= 0 {
		page = 1
	}

	firstIndex := (totalDataPerPage * page) - totalDataPerPage

	query := fmt.Sprintf("select id_unit,unit_code,unit_name from unit_information limit %d,%d", firstIndex, totalDataPerPage)

	rows, err := db.Query(query)
	if err != nil {
		log.Print(err)
	}
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
