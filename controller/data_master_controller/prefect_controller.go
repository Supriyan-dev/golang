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

func ReturnAllPrefect(w http.ResponseWriter, r *http.Request) {
	var prefect initialize.Prefect
	var arrPrefect []initialize.Prefect
	var response initialize.Response

	db := db.Connect()

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

	db := db.Connect()
	defer db.Close()
	code := mux.Vars(r)

	totalDataPerPage, _ := strconv.Atoi(code["perPage"])
	page, _ := strconv.Atoi(code["page"])

	var totalData int
	err := db.QueryRow("SELECT COUNT(*) FROM prefecture").Scan(&totalData)

	totalPage := int(math.Ceil(float64(totalData) / float64(totalDataPerPage)))

	if page > totalPage {
		page = totalPage
	}
	if page <= 0 {
		page = 1
	}

	firstIndex := (totalDataPerPage * page) - totalDataPerPage

	query := fmt.Sprintf("select id_prefecture,ISO,prefecture_name from prefecture limit %d,%d", firstIndex, totalDataPerPage)

	rows, err := db.Query(query)
	if err != nil {
		log.Print(err)
	}

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
	response.TotalPage = totalPage
	response.CurrentPage = page

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func CreatePrefect(w http.ResponseWriter, r *http.Request) {
	var err error
	var response initialize.Response

	db := db.Connect()
	stmt, err := db.Prepare("INSERT INTO prefecture (prefecture_name) VALUES(?)")
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
	name := keyVal["prefecture_name"]

	_, err = stmt.Exec(name)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Data baru telah dibuat")

	response.Status = 200
	response.Message = "Success"

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
