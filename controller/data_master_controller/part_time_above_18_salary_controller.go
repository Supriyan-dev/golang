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

func ReturnAllPartTimeAbove18Salary(w http.ResponseWriter, r *http.Request) {
	var partTimeSalary initialize.PartTimeAbove18Salary
	var arrPartTimeAbove18Salary []initialize.PartTimeAbove18Salary
	var response initialize.Response

	db, err := db.Connect()

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

	code := mux.Vars(r)
	fmt.Fprintf(w, "Category: %v\n", code["page"])

	db, err := db.Connect()
	rows, err := db.Query("SELECT * FROM part_time_above_18_salary ORDER BY id_part_time_above_18_salary LIMIT " + code["page"] + " OFFSET 0")
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
