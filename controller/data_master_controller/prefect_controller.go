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

func ReturnAllPrefect(w http.ResponseWriter, r *http.Request) {
	var prefect initialize.Prefect
	var arrPrefect []initialize.Prefect
	var response initialize.Response

	db, err := db.Connect()

	rows, err := db.Query("SELECT * FROM prefecture")
	if err != nil {
		log.Print(err)
	}
	defer db.Close()

	for rows.Next() {
		if err := rows.Scan(&prefect.Id_prefecture, &prefect.ISO, &prefect.Prefecture_name); err != nil {
			log.Fatal(err.Error())

		} else {
			arrPrefect = append(arrPrefect, prefect)
		}
	}
	response.Status = 200
	response.Message = "Success"
	response.Data = arrPrefect

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func ReturnAllPrefectPagination(w http.ResponseWriter, r *http.Request) {
	var prefect initialize.Prefect
	var arrPrefect []initialize.Prefect
	var response initialize.Response

	code := mux.Vars(r)
	fmt.Fprintf(w, "Category: %v\n", code["page"])

	db, err := db.Connect()
	rows, err := db.Query("SELECT * FROM prefecture ORDER BY id_prefecture LIMIT " + code["page"] + " OFFSET 0")
	if err != nil {
		log.Print(err)
	}
	defer db.Close()

	for rows.Next() {
		if err := rows.Scan(&prefect.Id_prefecture, &prefect.ISO, &prefect.Prefecture_name); err != nil {
			log.Fatal(err.Error())

		} else {
			arrPrefect = append(arrPrefect, prefect)
		}
	}
	response.Status = 200
	response.Message = "Success"
	response.Data = arrPrefect

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
