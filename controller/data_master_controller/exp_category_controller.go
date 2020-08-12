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

func ReturnAllExpCategory(w http.ResponseWriter, r *http.Request) {
	var exp initialize.ExpCategory
	var arrExpCategory []initialize.ExpCategory
	var response initialize.Response

	db, err := db.Connect()

	rows, err := db.Query("SELECT * FROM exp_category")

	if err != nil {
		log.Print(err)
	}
	defer db.Close()

	for rows.Next() {
		if err := rows.Scan(&exp.Id_exp, &exp.Exp_category, &exp.Created_date, &exp.Created_time, &exp.Code_category, &exp.Content, &exp.Rule_code); err != nil {

			log.Fatal(err.Error())

		} else {
			arrExpCategory = append(arrExpCategory, exp)
		}
	}
	response.Status = 200
	response.Message = "Success"
	response.Data = arrExpCategory

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func ReturnAllExpCategoryPagination(w http.ResponseWriter, r *http.Request) {
	var exp initialize.ExpCategory
	var arrExpCategory []initialize.ExpCategory
	var response initialize.Response
	code := mux.Vars(r)
	fmt.Fprintf(w, "Category: %v\n", code["page"])

	db, err := db.Connect()
	rows, err := db.Query("SELECT * FROM exp_category ORDER BY id_exp LIMIT " + code["page"] + " OFFSET 0")

	if err != nil {
		log.Print(err)
	}
	defer db.Close()

	for rows.Next() {
		if err := rows.Scan(&exp.Id_exp, &exp.Exp_category, &exp.Created_date, &exp.Created_time, &exp.Code_category, &exp.Content, &exp.Rule_code); err != nil {

			log.Fatal(err.Error())

		} else {
			arrExpCategory = append(arrExpCategory, exp)
		}
	}
	response.Status = 200
	response.Message = "Success"
	response.Data = arrExpCategory

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
