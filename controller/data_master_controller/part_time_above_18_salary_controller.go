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

func ReturnAllPartTimeAbove18Salary(w http.ResponseWriter, r *http.Request) {
	var partTimeSalary initialize.PartTimeAbove18Salary
	var arrPartTimeAbove18Salary []initialize.PartTimeAbove18Salary
	var response initialize.Response

	db := db.Connect()

	rows, err := db.Query("SELECT * FROM part_time_above_18_salary")

	if err != nil {
		log.Print(err)
	}
	defer db.Close()

	for rows.Next() {
		if err := rows.Scan(&partTimeSalary.Id_part_time_above_18_salary, &partTimeSalary.Id_code_store, &partTimeSalary.Day_salary, &partTimeSalary.Night_salary, &partTimeSalary.Morning_salary, &partTimeSalary.Peek_time_salary); err != nil {

			log.Fatal(err.Error())

		} else {
			arrPartTimeAbove18Salary = append(arrPartTimeAbove18Salary, partTimeSalary)
		}
	}
	response.Status = 200
	response.Message = "Success"
	response.Data = arrPartTimeAbove18Salary

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func ReturnAllPartTimeAbove18SalaryPagination(w http.ResponseWriter, r *http.Request) {
	var partTimeSalary initialize.PartTimeAbove18Salary
	var arrPartTimeAbove18Salary []initialize.PartTimeAbove18Salary
	var response initialize.Response

	db := db.Connect()
	defer db.Close()
	code := mux.Vars(r)

	totalDataPerPage, _ := strconv.Atoi(code["perPage"])
	page, _ := strconv.Atoi(code["page"])

	var totalData int
	err := db.QueryRow("SELECT COUNT(*) FROM part_time_above_18_salary").Scan(&totalData)

	totalPage := int(math.Ceil(float64(totalData) / float64(totalDataPerPage)))

	if page > totalPage {
		page = totalPage
	}
	if page <= 0 {
		page = 1
	}

	firstIndex := (totalDataPerPage * page) - totalDataPerPage

	query := fmt.Sprintf("select id_part_time_above_18_salary,id_code_store,day_salary,night_salary,morning_salary,peek_time_salary from part_time_above_18_salary limit %d,%d", firstIndex, totalDataPerPage)

	rows, err := db.Query(query)
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&partTimeSalary.Id_part_time_above_18_salary, &partTimeSalary.Id_code_store, &partTimeSalary.Day_salary, &partTimeSalary.Night_salary, &partTimeSalary.Morning_salary, &partTimeSalary.Peek_time_salary); err != nil {

			log.Fatal(err.Error())

		} else {
			arrPartTimeAbove18Salary = append(arrPartTimeAbove18Salary, partTimeSalary)
		}
	}
	response.Status = 200
	response.Message = "Success"
	response.Data = arrPartTimeAbove18Salary

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
