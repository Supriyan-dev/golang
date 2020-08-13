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

func ReturnAllStoreInformation(w http.ResponseWriter, r *http.Request) {
	var storeInformation initialize.StoreInformation
	var arrStoreInformation []initialize.StoreInformation
	var response initialize.Response

	db := db.Connect()

	rows, err := db.Query("SELECT * FROM store_information")
	if err != nil {
		log.Print(err)
	}
	defer db.Close()

	for rows.Next() {
		if err := rows.Scan(&storeInformation.Id_code_store, &storeInformation.Code_store, &storeInformation.Store_name); err != nil {
			log.Fatal(err.Error())

		} else {
			arrStoreInformation = append(arrStoreInformation, storeInformation)
		}
	}
	response.Status = 200
	response.Message = "Success"
	response.Data = arrStoreInformation

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func ReturnAllStoreInformationPagination(w http.ResponseWriter, r *http.Request) {
	var storeInformation initialize.StoreInformation
	var arrStoreInformation []initialize.StoreInformation
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

	query := fmt.Sprintf("select id_code_store,code_store,store_name from store_information limit %d,%d", firstIndex, totalDataPerPage)

	rows, err := db.Query(query)
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&storeInformation.Id_code_store, &storeInformation.Code_store, &storeInformation.Store_name); err != nil {
			log.Fatal(err.Error())
		} else {
			arrStoreInformation = append(arrStoreInformation, storeInformation)
		}
	}

	response.Status = 200
	response.Message = "success"
	response.Data = arrStoreInformation

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
