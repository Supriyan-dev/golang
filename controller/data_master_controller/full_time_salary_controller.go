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

func ReturnAllFullTimeSalary(w http.ResponseWriter, r *http.Request) {
	var salary initialize.FullTimeSalary
	var arrFullTimeSalary []initialize.FullTimeSalary
	var response initialize.Response

	db := db.Connect()

	rows, err := db.Query("SELECT * FROM full_time_salary")

	if err != nil {
		log.Print(err)
	}
	defer db.Close()

	for rows.Next() {
		if err := rows.Scan(&salary.Id_full_time_salary, &salary.Id_code_store, &salary.Salary, &salary.Fish_section_salary); err != nil {

			log.Fatal(err.Error())

		} else {
			arrFullTimeSalary = append(arrFullTimeSalary, salary)
		}
	}
	response.Status = 200
	response.Message = "Success"
	response.Data = arrFullTimeSalary

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func ReturnAllFullTimeSalaryPagination(w http.ResponseWriter, r *http.Request) {
	var salary initialize.FullTimeSalary
	var arrFullTimeSalary []initialize.FullTimeSalary
	var response initialize.Response

	db := db.Connect()
	defer db.Close()
	code := mux.Vars(r)

	totalDataPerPage, _ := strconv.Atoi(code["perPage"])
	page, _ := strconv.Atoi(code["page"])

	var totalData int
	err := db.QueryRow("SELECT COUNT(*) FROM full_time_salary").Scan(&totalData)

	totalPage := int(math.Ceil(float64(totalData) / float64(totalDataPerPage)))

	if page > totalPage {
		page = totalPage
	}
	if page <= 0 {
		page = 1
	}

	firstIndex := (totalDataPerPage * page) - totalDataPerPage

	query := fmt.Sprintf("select id_full_time_salary,id_code_store,salary,fish_section_salary from full_time_salary limit %d,%d", firstIndex, totalDataPerPage)

	rows, err := db.Query(query)
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&salary.Id_full_time_salary, &salary.Id_code_store, &salary.Salary, &salary.Fish_section_salary); err != nil {

			log.Fatal(err.Error())

		} else {
			arrFullTimeSalary = append(arrFullTimeSalary, salary)
		}
	}
	response.Status = 200
	response.Message = "Success"
	response.Data = arrFullTimeSalary
	response.TotalPage = totalPage
	response.CurrentPage = page

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
