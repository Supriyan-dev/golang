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

func ReturnAllFullTimeSalary(w http.ResponseWriter, r *http.Request) {
	var salary initialize.FullTimeSalary
	var arrFullTimeSalary []initialize.FullTimeSalary
	var response initialize.Response

	db := db.Connect()

	rows, err := db.Query("SELECT * FROM full_time_salary")

	if err != nil {
		log.Print(err)
	}
	defer db.Close()

	for rows.Next() {
		if err := rows.Scan(&salary.Id_full_time_salary, &salary.Id_code_store, &salary.Salary, &salary.Fish_section_salary); err != nil {

			log.Fatal(err.Error())

		} else {
			arrFullTimeSalary = append(arrFullTimeSalary, salary)
		}
	}
	response.Status = 200
	response.Message = "Success"
	response.Data = arrFullTimeSalary

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func ReturnAllFullTimeSalaryPagination(w http.ResponseWriter, r *http.Request) {
	var salary initialize.FullTimeSalary
	var arrFullTimeSalary []initialize.FullTimeSalary
	var response initialize.Response

	db := db.Connect()
	defer db.Close()
	code := mux.Vars(r)

	totalDataPerPage, _ := strconv.Atoi(code["perPage"])
	page, _ := strconv.Atoi(code["page"])

	var totalData int
	err := db.QueryRow("SELECT COUNT(*) FROM full_time_salary").Scan(&totalData)

	totalPage := int(math.Ceil(float64(totalData) / float64(totalDataPerPage)))

	if page > totalPage {
		page = totalPage
	}
	if page <= 0 {
		page = 1
	}

	firstIndex := (totalDataPerPage * page) - totalDataPerPage

	query := fmt.Sprintf("select id_full_time_salary,id_code_store,salary,fish_section_salary from full_time_salary limit %d,%d", firstIndex, totalDataPerPage)

	rows, err := db.Query(query)
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&salary.Id_full_time_salary, &salary.Id_code_store, &salary.Salary, &salary.Fish_section_salary); err != nil {

			log.Fatal(err.Error())

		} else {
			arrFullTimeSalary = append(arrFullTimeSalary, salary)
		}
	}
	response.Status = 200
	response.Message = "Success"
	response.Data = arrFullTimeSalary
	response.TotalPage = totalPage
	response.CurrentPage = page

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func GetFullTimeSalary(w http.ResponseWriter, r *http.Request) {
	var fullTimeSalary initialize.FullTimeSalary
	var arrFullTimeSalary []initialize.FullTimeSalary
	var response initialize.Response

	db := db.Connect()
	code := mux.Vars(r)
	fmt.Fprintf(w, "Category: %v\n", code["id_full_time_salary"])

	result, err := db.Query("SELECT id_full_time_salary, id_code_store, salary, fish_section_salary FROM full_time_salary WHERE id_full_time_salary = ?", code["id_full_time_salary"])
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {

		err := result.Scan(&fullTimeSalary.Id_full_time_salary, &fullTimeSalary.Id_code_store, &fullTimeSalary.Salary, &fullTimeSalary.Fish_section_salary)
		if err != nil {
			panic(err.Error())
		} else {
			arrFullTimeSalary = append(arrFullTimeSalary, fullTimeSalary)
		}
	}
	response.Status = 200
	response.Message = "Success"
	response.Data = arrFullTimeSalary

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func CreateFullTimeSalary(w http.ResponseWriter, r *http.Request) {
	var err error
	var response initialize.Response

	db := db.Connect()
	stmt, err := db.Prepare("INSERT INTO full_time_salary (id_code_store, salary, fish_section_salary) VALUES(?,?,?)")
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
	IdCodeStore := keyVal["id_code_store"]
	Salary := keyVal["salary"]
	FirstSectionSalary := keyVal["fish_section_salary"]

	result, err := stmt.Exec(IdCodeStore, Salary, FirstSectionSalary)
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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func UpdateFullTimeSalary(w http.ResponseWriter, r *http.Request) {
	var response initialize.Response

	db := db.Connect()

	stmt, err := db.Prepare("UPDATE full_time_salary SET id_code_store = ?, salary = ?, fish_section_salary = ? WHERE id_full_time_salary = ?")
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	idFullTimeSalary := keyVal["id_full_time_salary"]
	newIdCodeStore := keyVal["id_code_store"]
	newSalary := keyVal["salary"]
	NewFirstSectionSalary := keyVal["fish_section_salary"]

	id, err := strconv.Atoi(idFullTimeSalary)

	result, err := stmt.Exec(newIdCodeStore, newSalary, NewFirstSectionSalary, id)
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

func DeleteFullTimeSalary(w http.ResponseWriter, r *http.Request) {

	db := db.Connect()
	params := mux.Vars(r)
	stmt, err := db.Prepare("DELETE FROM full_time_salary WHERE id_full_time_salary = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(params["id_full_time_salary"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Data Sudah Terhapus Dengan ID = ")

	w.Header().Set("Content-Type", "Aplication/json")
	json.NewEncoder(w).Encode(params["id_full_time_salary"])

}
