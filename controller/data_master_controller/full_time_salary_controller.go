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

func ReturnAllFullTimeSalary(w http.ResponseWriter, r *http.Request) {
	var salary initialize.FullTimeSalary
	var arrFullTimeSalary []initialize.FullTimeSalary
	var response initialize.Response

	db, err := db.Connect()

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

	code := mux.Vars(r)
	fmt.Fprintf(w, "Category: %v\n", code["page"])

	db, err := db.Connect()
	rows, err := db.Query("SELECT * FROM full_time_salary ORDER BY id_full_time_salary LIMIT " + code["page"] + " OFFSET 0")
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
