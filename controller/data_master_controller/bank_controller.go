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

func ReturnAllBank(w http.ResponseWriter, r *http.Request) {
	var bank initialize.Bank
	var arrBank []initialize.Bank
	var response initialize.Response

	db, err := db.Connect()

	rows, err := db.Query("SELECT * FROM bank")

	if err != nil {
		log.Print(err)
	}
	defer db.Close()

	for rows.Next() {
		if err := rows.Scan(&bank.Id_bank, &bank.Bank_code, &bank.Bank_name, &bank.Branch_code, &bank.Branch_name, &bank.Special); err != nil {

			log.Fatal(err.Error())

		} else {
			arrBank = append(arrBank, bank)
		}
	}
	response.Status = 200
	response.Message = "Success"
	response.Data = arrBank

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func ReturnAllBankPagination(w http.ResponseWriter, r *http.Request) {
	var bank initialize.Bank
	var arrBank []initialize.Bank
	var response initialize.Response
	code := mux.Vars(r)
	fmt.Fprintf(w, "Category: %v\n", code["page"])

	db, err := db.Connect()
	rows, err := db.Query("SELECT * FROM bank ORDER BY id_bank LIMIT " + code["page"] + " OFFSET 0")
	if err != nil {
		log.Print(err)
	}
	defer db.Close()

	for rows.Next() {
		if err := rows.Scan(&bank.Id_bank, &bank.Bank_code, &bank.Bank_name, &bank.Branch_code, &bank.Branch_name, &bank.Special); err != nil {

			log.Fatal(err.Error())

		} else {
			arrBank = append(arrBank, bank)
		}
	}
	response.Status = 200
	response.Message = "Success"
	response.Data = arrBank

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
