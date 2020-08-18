package data_master_controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func GetStoreSectionInformation(w http.ResponseWriter, r *http.Request) {
	var storeSectionInformation initialize.StoreSectionInformation
	var response initialize.Response
	var arrStoreSectionInformation []initialize.StoreSectionInformation

	db := db.Connect()
	code := mux.Vars(r)
	fmt.Fprintf(w, "Category: %v\n", code["id_store_section"])

	result, err := db.Query("SELECT id_store_section, store_section_code, store_section_name FROM store_section_information WHERE id_store_section = ?", code["id_store_section"])
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {

		err := result.Scan(&storeSectionInformation.Id_store_section, &storeSectionInformation.Store_section_code, &storeSectionInformation.Store_section_name)
		if err != nil {
			panic(err.Error())
		} else {
			arrStoreSectionInformation = append(arrStoreSectionInformation, storeSectionInformation)
		}
	}
	response.Status = 200
	response.Message = "Success"
	response.Data = arrStoreSectionInformation

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func CreateStoreSectionInformation(w http.ResponseWriter, r *http.Request) {

	var response initialize.Response

	db := db.Connect()
	stmt, err := db.Prepare("INSERT INTO store_section_information (store_section_code,store_section_name) VALUES (?,?)")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	CodeStore := keyVal["store_section_code"]
	StoreName := keyVal["store_section_name"]

	result, err := stmt.Exec(CodeStore, StoreName)
	if err != nil {
		panic(err.Error())
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}

	response.Status = 200
	response.Message = "Success"
	response.Data = map[string]int64{
		"Data Yang Behasil Di Tambahkan": rowsAffected,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func UpdateStoreSectionInformation(w http.ResponseWriter, r *http.Request) {
	var response initialize.Response

	db := db.Connect()

	stmt, err := db.Prepare("UPDATE store_section_information SET store_section_code = ?, store_section_name = ? WHERE id_store_section = ?")
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	idCodeStore := keyVal["id_store_section"]
	newCodeSection := keyVal["store_section_code"]
	newNameSection := keyVal["store_section_name"]

	id, err := strconv.Atoi(idCodeStore)

	result, err := stmt.Exec(newCodeSection, newNameSection, id)
	if err != nil {
		panic(err.Error())
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}

	response.Status = 200
	response.Message = "Success"
	response.Data = map[string]int64{
		"Data Yang Behasil Di Update": rowsAffected,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func DeleteStoreSectionInformation(w http.ResponseWriter, r *http.Request) {

	db := db.Connect()
	params := mux.Vars(r)
	stmt, err := db.Prepare("DELETE FROM store_section_information WHERE id_store_section = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(params["id_store_section"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Data Sudah Terhapus Dengan ID = ")

	w.Header().Set("Content-Type", "Aplication/json")
	json.NewEncoder(w).Encode(params["id_store_section"])

}
