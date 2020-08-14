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

func ReturnAllBank(w http.ResponseWriter, r *http.Request) {
	var bank initialize.Bank
	var arrBank []initialize.Bank
	var response initialize.Response

	db := db.Connect()

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

	db := db.Connect()
	defer db.Close()
	code := mux.Vars(r)

	totalDataPerPage, _ := strconv.Atoi(code["perPage"])
	page, _ := strconv.Atoi(code["page"])

	var totalData int
	err := db.QueryRow("SELECT COUNT(*) FROM bank").Scan(&totalData)

	totalPage := int(math.Ceil(float64(totalData) / float64(totalDataPerPage)))

	if page > totalPage {
		page = totalPage
	}
	if page <= 0 {
		page = 1
	}

	firstIndex := (totalDataPerPage * page) - totalDataPerPage

	query := fmt.Sprintf("select id_bank,bank_code,bank_name,branch_code,branch_name,special FROM bank limit %d,%d", firstIndex, totalDataPerPage)

	rows, err := db.Query(query)
	if err != nil {
		log.Print(err)
	}
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
	response.TotalPage = totalPage
	response.CurrentPage = page

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
