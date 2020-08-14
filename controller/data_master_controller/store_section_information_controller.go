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

func ReturnAllStroreSectionInformation(w http.ResponseWriter, r *http.Request) {
	var store initialize.StoreSectionInformation
	var arrStoreSectionInformation []initialize.StoreSectionInformation
	var response initialize.Response

	db := db.Connect()

	rows, err := db.Query("SELECT * FROM store_section_information")

	if err != nil {
		log.Print(err)
	}
	defer db.Close()

	for rows.Next() {
		// if err := rows.Scan(&prefect.Id_prefecture, &prefect.ISO, &prefect.Prefecture_name); err != nil {
		if err := rows.Scan(&store.Id_store_section, &store.Store_section_code, &store.Store_section_name); err != nil {

			log.Fatal(err.Error())

		} else {
			arrStoreSectionInformation = append(arrStoreSectionInformation, store)
		}
	}
	response.Status = 200
	response.Message = "Success"
	response.Data = arrStoreSectionInformation

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func ReturnAllStroreSectionInformationPagination(w http.ResponseWriter, r *http.Request) {
	var store initialize.StoreSectionInformation
	var arrStoreSectionInformation []initialize.StoreSectionInformation
	var response initialize.Response
	db := db.Connect()
	defer db.Close()
	code := mux.Vars(r)

	totalDataPerPage, _ := strconv.Atoi(code["perPage"])
	page, _ := strconv.Atoi(code["page"])

	var totalData int
	err := db.QueryRow("SELECT COUNT(*) FROM store_section_information").Scan(&totalData)

	totalPage := int(math.Ceil(float64(totalData) / float64(totalDataPerPage)))

	if page > totalPage {
		page = totalPage
	}
	if page <= 0 {
		page = 1
	}

	firstIndex := (totalDataPerPage * page) - totalDataPerPage

	query := fmt.Sprintf("select id_store_section,store_section_code,store_section_name from store_section_information limit %d,%d", firstIndex, totalDataPerPage)

	rows, err := db.Query(query)
	if err != nil {
		log.Print(err)
	}
	for rows.Next() {
		if err := rows.Scan(&store.Id_store_section, &store.Store_section_code, &store.Store_section_name); err != nil {

			log.Fatal(err.Error())

		} else {
			arrStoreSectionInformation = append(arrStoreSectionInformation, store)
		}
	}
	response.Status = 200
	response.Message = "Success"
	response.Data = arrStoreSectionInformation
	response.TotalPage = totalPage
	response.CurrentPage = page

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
