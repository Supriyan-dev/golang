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

func ReturnAllDepartementInformation(w http.ResponseWriter, r *http.Request) {
	var DeptInfo initialize.DepartementInformation
	var arrDepartementInformation []initialize.DepartementInformation
	var response initialize.Response

	db, err := db.Connect()

	rows, err := db.Query("SELECT * FROM department_information")

	if err != nil {
		log.Print(err)
	}
	defer db.Close()

	for rows.Next() {
		if err := rows.Scan(&DeptInfo.Id_department, &DeptInfo.Department_code, &DeptInfo.Department_name, &DeptInfo.Id_code_store); err != nil {

			log.Fatal(err.Error())

		} else {
			arrDepartementInformation = append(arrDepartementInformation, DeptInfo)
		}
	}
	response.Status = 200
	response.Message = "Success"
	response.Data = arrDepartementInformation

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func ReturnAllDepartementInformationPagination(w http.ResponseWriter, r *http.Request) {
	var DeptInfo initialize.DepartementInformation
	var arrDepartementInformation []initialize.DepartementInformation
	var response initialize.Response
	code := mux.Vars(r)
	fmt.Fprintf(w, "Category: %v\n", code["page"])

	db, err := db.Connect()
	rows, err := db.Query("SELECT * FROM department_information ORDER BY id_department LIMIT " + code["page"] + " OFFSET 0")
	if err != nil {
		log.Print(err)
	}
	defer db.Close()

	for rows.Next() {
		if err := rows.Scan(&DeptInfo.Id_department, &DeptInfo.Department_code, &DeptInfo.Department_name, &DeptInfo.Id_code_store); err != nil {

			log.Fatal(err.Error())

		} else {
			arrDepartementInformation = append(arrDepartementInformation, DeptInfo)
		}
	}
	response.Status = 200
	response.Message = "Success"
	response.Data = arrDepartementInformation

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
