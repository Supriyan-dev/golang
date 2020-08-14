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

func ReturnAllDepartementInformation(w http.ResponseWriter, r *http.Request) {
	var DeptInfo initialize.DepartementInformation
	var arrDepartementInformation []initialize.DepartementInformation
	var response initialize.Response

	db := db.Connect()

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
	db := db.Connect()
	defer db.Close()
	code := mux.Vars(r)

	totalDataPerPage, _ := strconv.Atoi(code["perPage"])
	page, _ := strconv.Atoi(code["page"])

	var totalData int
	err := db.QueryRow("SELECT COUNT(*) FROM store_information").Scan(&totalData)

	totalPage := int(math.Ceil(float64(totalData) / float64(totalDataPerPage)))

	if page > totalPage {
		page = totalPage
	}
	if page <= 0 {
		page = 1
	}

	firstIndex := (totalDataPerPage * page) - totalDataPerPage

	query := fmt.Sprintf("select id_department,department_code,department_name,id_code_store from department_information limit %d,%d", firstIndex, totalDataPerPage)

	rows, err := db.Query(query)
	if err != nil {
		log.Print(err)
	}
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
	response.TotalPage = totalPage
	response.CurrentPage = page

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
