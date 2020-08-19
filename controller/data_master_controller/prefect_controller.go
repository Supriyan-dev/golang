package data_master_controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strconv"

	"../../db"
	"../../initialize"
	"github.com/gorilla/mux"
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

	w.Header().Set("Content-Type", "application/json", "Access-Control-Allow-Origin", "*")
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

	w.Header().Set("Content-Type", "application/json", "Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)

}

func GetPrefect(w http.ResponseWriter, r *http.Request) {
	var prefect initialize.Prefect
	var arrPrefect []initialize.Prefect
	var response initialize.Response

	db := db.Connect()
	code := mux.Vars(r)
	fmt.Fprintf(w, "Category: %v\n", code["id_prefecture"])

	result, err := db.Query("SELECT id_prefecture, ISO, prefecture_name FROM prefecture WHERE id_prefecture = ?", code["id_prefecture"])
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {

		err := result.Scan(&prefect.Id_prefecture, &prefect.ISO, &prefect.Prefecture_name)
		if err != nil {
			panic(err.Error())
		} else {
			arrPrefect = append(arrPrefect, prefect)
		}
	}
	response.Status = 200
	response.Message = "Success"
	response.Data = arrPrefect

	w.Header().Set("Content-Type", "application/json", "Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

func CreatePrefect(w http.ResponseWriter, r *http.Request) {
	var err error
	var response initialize.Response

	db := db.Connect()
	stmt, err := db.Prepare("INSERT INTO prefecture (ISO, prefecture_name) VALUES(?,?)")
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
	ISO := keyVal["ISO"]
	name := keyVal["prefecture_name"]

	result, err := stmt.Exec(ISO, name)
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
		"Data baru telah dibuat": rowsAffected,
	}

	w.Header().Set("Content-Type", "application/json", "Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)

}

func UpdatePrefect(w http.ResponseWriter, r *http.Request) {
	var response initialize.Response

	db := db.Connect()

	stmt, err := db.Prepare("UPDATE prefecture SET ISO = ?, prefecture_name = ? WHERE id_prefecture = ?")
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	idPrefecture := keyVal["id_prefecture"]
	newISO := keyVal["ISO"]
	newNamePrefecture := keyVal["prefecture_name"]

	id, err := strconv.Atoi(idPrefecture)

	result, err := stmt.Exec(newISO, newNamePrefecture, id)
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

	w.Header().Set("Content-Type", "application/json", "Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

func DeletePrefect(w http.ResponseWriter, r *http.Request) {

	db := db.Connect()
	params := mux.Vars(r)
	stmt, err := db.Prepare("DELETE FROM prefecture WHERE id_prefecture = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(params["id_prefecture"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Data Sudah Terhapus Dengan ID = ")

	w.Header().Set("Content-Type", "application/json", "Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(params["id_prefecture"])

}
