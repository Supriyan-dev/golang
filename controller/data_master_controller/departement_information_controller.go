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

func ReturnAllDepartementInformation(w http.ResponseWriter, r *http.Request) {
	var DeptInfo initialize.DepartementInformation
	var arrDepartementInformation []initialize.DepartementInformation
	var response initialize.Response

	db := db.Connect()

	rows, err := db.Query("SELECT * FROM department_information")

	if err != nil {
		log.Print(err)
	}
	defer db.Close()

	for rows.Next() {
		if err := rows.Scan(&DeptInfo.Id_department, &DeptInfo.Department_code, &DeptInfo.Department_name, &DeptInfo.Id_code_store); err != nil {

			log.Fatal(err.Error())

		} else {
			arrDepartementInformation = append(arrDepartementInformation, DeptInfo)
		}
	}
	response.Status = 200
	response.Message = "Success"
	response.Data = arrDepartementInformation

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func ReturnAllDepartementInformationPagination(w http.ResponseWriter, r *http.Request) {
	var DeptInfo initialize.DepartementInformation
	var arrDepartementInformation []initialize.DepartementInformation
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

	query := fmt.Sprintf("select id_department,department_code,department_name,id_code_store from department_information limit %d,%d", firstIndex, totalDataPerPage)

	rows, err := db.Query(query)
	if err != nil {
		log.Print(err)
	}
	for rows.Next() {
		if err := rows.Scan(&DeptInfo.Id_department, &DeptInfo.Department_code, &DeptInfo.Department_name, &DeptInfo.Id_code_store); err != nil {

			log.Fatal(err.Error())

		} else {
			arrDepartementInformation = append(arrDepartementInformation, DeptInfo)
		}
	}
	response.Status = 200
	response.Message = "Success"
	response.Data = arrDepartementInformation
	response.TotalPage = totalPage
	response.CurrentPage = page

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func GetDepartementInformation(w http.ResponseWriter, r *http.Request) {
	var DeptInfo initialize.DepartementInformation
	var response initialize.Response
	var arrDepartementInformation []initialize.DepartementInformation

	db := db.Connect()
	code := mux.Vars(r)
	fmt.Fprintf(w, "Category: %v\n", code["id_department"])

	result, err := db.Query("SELECT id_department, department_code, department_name, id_code_store FROM department_information WHERE id_department = ?", code["id_department"])
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {

		err := result.Scan(&DeptInfo.Id_department, &DeptInfo.Department_code, &DeptInfo.Department_name, &DeptInfo.Id_code_store)
		if err != nil {
			panic(err.Error())
		} else {
			arrDepartementInformation = append(arrDepartementInformation, DeptInfo)
		}
	}
	response.Status = 200
	response.Message = "Success"
	response.Data = arrDepartementInformation

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func CreateDepartementInformation(w http.ResponseWriter, r *http.Request) {
	var response initialize.Response

	db := db.Connect()
	stmt, err := db.Prepare("INSERT INTO department_information (department_code,department_name,id_code_store) VALUES (?,?,?)")
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
	DeptCode := keyVal["department_code"]
	DeptName := keyVal["department_name"]
	IdDeptStore := keyVal["id_code_store"]

	result, err := stmt.Exec(DeptCode, DeptName, IdDeptStore)
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

func UpdateDepartementInformation(w http.ResponseWriter, r *http.Request) {
	var response initialize.Response

	db := db.Connect()

	stmt, err := db.Prepare("UPDATE department_information SET department_code = ?, department_name = ?, id_code_store = ? WHERE id_department = ?")
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	idCodeDept := keyVal["id_department"]
	newCode := keyVal["department_code"]
	newName := keyVal["department_name"]
	IdCodeStore := keyVal["id_code_store"]

	id, err := strconv.Atoi(idCodeDept)

	result, err := stmt.Exec(newCode, newName, IdCodeStore, id)
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

func DeleteDepartementInformation(w http.ResponseWriter, r *http.Request) {

	db := db.Connect()
	params := mux.Vars(r)
	stmt, err := db.Prepare("DELETE FROM department_information WHERE id_department = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(params["id_department"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Data Sudah Terhapus Dengan ID = ")

	w.Header().Set("Content-Type", "Aplication/json")
	json.NewEncoder(w).Encode(params["id_department"])

}
