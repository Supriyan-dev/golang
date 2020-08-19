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
	"../../db"
	"../../initialize"
)

func ReturnAllUnitInformation(w http.ResponseWriter, r *http.Request) {
	var unit initialize.UnitInformation
	var arrUnitInformation []initialize.UnitInformation
	var response initialize.Response

	db := db.Connect()

	rows, err := db.Query("SELECT * FROM unit_information")

	if err != nil {
		log.Print(err)
	}
	defer db.Close()

	for rows.Next() {
		if err := rows.Scan(&unit.Id_unit, &unit.Unit_code, &unit.Unit_name); err != nil {

			log.Fatal(err.Error())

		} else {
			arrUnitInformation = append(arrUnitInformation, unit)
		}
	}
	response.Status = 200
	response.Message = "Success"
	response.Data = arrUnitInformation

	w.Header().Set("Content-Type", "application/json", "*")
	json.NewEncoder(w).Encode(response)

}

func ReturnAllUnitInformationPagination(w http.ResponseWriter, r *http.Request) {
	var unit initialize.UnitInformation
	var arrUnitInformation []initialize.UnitInformation
	var response initialize.Response

	db := db.Connect()
	defer db.Close()
	code := mux.Vars(r)

	totalDataPerPage, _ := strconv.Atoi(code["perPage"])
	page, _ := strconv.Atoi(code["page"])

	var totalData int
	err := db.QueryRow("SELECT COUNT(*) FROM unit_information").Scan(&totalData)

	totalPage := int(math.Ceil(float64(totalData) / float64(totalDataPerPage)))

	if page > totalPage {
		page = totalPage
	}
	if page <= 0 {
		page = 1
	}

	firstIndex := (totalDataPerPage * page) - totalDataPerPage

	query := fmt.Sprintf("select id_unit,unit_code,unit_name from unit_information limit %d,%d", firstIndex, totalDataPerPage)

	rows, err := db.Query(query)
	if err != nil {
		log.Print(err)
	}
	for rows.Next() {
		if err := rows.Scan(&unit.Id_unit, &unit.Unit_code, &unit.Unit_name); err != nil {

			log.Fatal(err.Error())

		} else {
			arrUnitInformation = append(arrUnitInformation, unit)
		}
	}
	response.Status = 200
	response.Message = "Success"
	response.Data = arrUnitInformation
	response.TotalPage = totalPage
	response.CurrentPage = page

	w.Header().Set("Content-Type", "application/json", "*")
	json.NewEncoder(w).Encode(response)

}

func GetUnitInformation(w http.ResponseWriter, r *http.Request) {
	var unit initialize.UnitInformation
	var arrUnitInformation []initialize.UnitInformation
	var response initialize.Response

	db := db.Connect()
	code := mux.Vars(r)
	fmt.Fprintf(w, "Category: %v\n", code["id_unit"])

	result, err := db.Query("SELECT id_unit, unit_code, unit_name FROM unit_information WHERE id_unit = ?", code["id_unit"])
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {

		err := result.Scan(&unit.Id_unit, &unit.Unit_code, &unit.Unit_name)
		if err != nil {
			panic(err.Error())
		} else {
			arrUnitInformation = append(arrUnitInformation, unit)
		}
	}
	response.Status = 200
	response.Message = "Success"
	response.Data = arrUnitInformation

	w.Header().Set("Content-Type", "application/json", "*")
	json.NewEncoder(w).Encode(response)
}

func CreateUnitInformation(w http.ResponseWriter, r *http.Request) {
	var response initialize.Response

	db := db.Connect()
	stmt, err := db.Prepare("INSERT INTO unit_information (unit_code,unit_name) VALUES (?,?)")
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
	UnitCode := keyVal["unit_code"]
	UnitName := keyVal["unit_name"]

	result, err := stmt.Exec(UnitCode, UnitName)
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

	w.Header().Set("Content-Type", "application/json", "*")
	json.NewEncoder(w).Encode(response)

}

func UpdateUnitInformation(w http.ResponseWriter, r *http.Request) {
	var response initialize.Response

	db := db.Connect()

	stmt, err := db.Prepare("UPDATE unit_information SET unit_code = ?, unit_name = ? WHERE id_unit = ?")
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	idUnit := keyVal["id_unit"]
	newUnitCode := keyVal["unit_code"]
	newUnitName := keyVal["unit_name"]

	id, err := strconv.Atoi(idUnit)

	result, err := stmt.Exec(newUnitCode, newUnitName, id)
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

	w.Header().Set("Content-Type", "application/json", "*")
	json.NewEncoder(w).Encode(response)
}

func DeleteUnitInformation(w http.ResponseWriter, r *http.Request) {

	db := db.Connect()
	params := mux.Vars(r)
	stmt, err := db.Prepare("DELETE FROM unit_information WHERE id_unit = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(params["id_unit"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Data Sudah Terhapus Dengan ID = ")

	w.Header().Set("Content-Type", "application/json", "*")
	json.NewEncoder(w).Encode(params["id_unit"])

}
