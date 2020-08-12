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

func ReturnAllPartTimeUnder18Salary(w http.ResponseWriter, r *http.Request) {
	var partTimeUnderSalary initialize.PartTimeUnder18Salary
	var arrPartTimeUnder18Salary []initialize.PartTimeUnder18Salary
	var response initialize.Response

	db, err := db.Connect()

	rows, err := db.Query("SELECT * FROM part_time_under_18_salary")

	if err != nil {
		log.Print(err)
	}
	defer db.Close()

	for rows.Next() {
		if err := rows.Scan(&partTimeUnderSalary.Id_part_time_under_18_salary, &partTimeUnderSalary.Id_code_store, &partTimeUnderSalary.Salary); err != nil {

			log.Fatal(err.Error())

		} else {
			arrPartTimeUnder18Salary = append(arrPartTimeUnder18Salary, partTimeUnderSalary)
		}
	}
	response.Status = 200
	response.Message = "Success"
	response.Data = arrPartTimeUnder18Salary

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func ReturnAllPartTimeUnder18SalaryPagination(w http.ResponseWriter, r *http.Request) {
	var partTimeUnderSalary initialize.PartTimeUnder18Salary
	var arrPartTimeUnder18Salary []initialize.PartTimeUnder18Salary
	var response initialize.Response
	code := mux.Vars(r)
	fmt.Fprintf(w, "Category: %v\n", code["page"])

	db, err := db.Connect()
	rows, err := db.Query("SELECT * FROM part_time_under_18_salary ORDER BY id_part_time_under_18_salary LIMIT " + code["page"] + " OFFSET 0")
	if err != nil {
		log.Print(err)
	}
	defer db.Close()

	for rows.Next() {
		if err := rows.Scan(&partTimeUnderSalary.Id_part_time_under_18_salary, &partTimeUnderSalary.Id_code_store, &partTimeUnderSalary.Salary); err != nil {

			log.Fatal(err.Error())

		} else {
			arrPartTimeUnder18Salary = append(arrPartTimeUnder18Salary, partTimeUnderSalary)
		}
	}
	response.Status = 200
	response.Message = "Success"
	response.Data = arrPartTimeUnder18Salary

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
