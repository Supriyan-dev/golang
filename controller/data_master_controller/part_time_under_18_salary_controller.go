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

func ReturnAllPartTimeUnder18Salary(w http.ResponseWriter, r *http.Request) {
	var partTimeUnderSalary initialize.PartTimeUnder18Salary
	var arrPartTimeUnder18Salary []initialize.PartTimeUnder18Salary
	var response initialize.Response

	db := db.Connect()

	rows, err := db.Query("SELECT * FROM part_time_under_18_salary")

	if err != nil {
		log.Print(err)
	}
	defer db.Close()

	for rows.Next() {
		if err := rows.Scan(&partTimeUnderSalary.Id_part_time_under_18_salary, &partTimeUnderSalary.Id_code_store, &partTimeUnderSalary.Salary); err != nil {

			log.Fatal(err.Error())

		} else {
			arrPartTimeUnder18Salary = append(arrPartTimeUnder18Salary, partTimeUnderSalary)
		}
	}
	response.Status = 200
	response.Message = "Success"
	response.Data = arrPartTimeUnder18Salary

	w.Header().Set("Content-Type", "application/json", "*")
	json.NewEncoder(w).Encode(response)

}

func ReturnAllPartTimeUnder18SalaryPagination(w http.ResponseWriter, r *http.Request) {
	var partTimeUnderSalary initialize.PartTimeUnder18Salary
	var arrPartTimeUnder18Salary []initialize.PartTimeUnder18Salary
	var response initialize.Response

	db := db.Connect()
	defer db.Close()
	code := mux.Vars(r)

	totalDataPerPage, _ := strconv.Atoi(code["perPage"])
	page, _ := strconv.Atoi(code["page"])

	var totalData int
	err := db.QueryRow("SELECT COUNT(*) FROM part_time_under_18_salary").Scan(&totalData)

	totalPage := int(math.Ceil(float64(totalData) / float64(totalDataPerPage)))

	if page > totalPage {
		page = totalPage
	}
	if page <= 0 {
		page = 1
	}

	firstIndex := (totalDataPerPage * page) - totalDataPerPage

	query := fmt.Sprintf("select id_part_time_under_18_salary,id_code_store,salary from part_time_under_18_salary limit %d,%d", firstIndex, totalDataPerPage)

	rows, err := db.Query(query)
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&partTimeUnderSalary.Id_part_time_under_18_salary, &partTimeUnderSalary.Id_code_store, &partTimeUnderSalary.Salary); err != nil {

			log.Fatal(err.Error())

		} else {
			arrPartTimeUnder18Salary = append(arrPartTimeUnder18Salary, partTimeUnderSalary)
		}
	}
	response.Status = 200
	response.Message = "Success"
	response.Data = arrPartTimeUnder18Salary
	response.TotalPage = totalPage
	response.CurrentPage = page

	w.Header().Set("Content-Type", "application/json", "*")
	json.NewEncoder(w).Encode(response)

}

func GetPartTimeUnder18Salary(w http.ResponseWriter, r *http.Request) {
	var partTimeUnder18Salary initialize.PartTimeUnder18Salary
	var arrPartTimeUnder18Salary []initialize.PartTimeUnder18Salary
	var response initialize.Response

	db := db.Connect()
	code := mux.Vars(r)
	fmt.Fprintf(w, "Category: %v\n", code["id_part_time_under_18_salary"])

	result, err := db.Query("SELECT id_part_time_under_18_salary, id_code_store, salary FROM part_time_under_18_salary WHERE id_part_time_under_18_salary = ?", code["id_part_time_under_18_salary"])
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {

		err := result.Scan(&partTimeUnder18Salary.Id_part_time_under_18_salary, &partTimeUnder18Salary.Id_code_store, &partTimeUnder18Salary.Salary)
		if err != nil {
			panic(err.Error())
		} else {
			arrPartTimeUnder18Salary = append(arrPartTimeUnder18Salary, partTimeUnder18Salary)
		}
	}
	response.Status = 200
	response.Message = "Success"
	response.Data = arrPartTimeUnder18Salary

	w.Header().Set("Content-Type", "application/json", "*")
	json.NewEncoder(w).Encode(response)
}

func CreatePartTimeUnder18Salary(w http.ResponseWriter, r *http.Request) {
	var err error
	var response initialize.Response

	db := db.Connect()
	stmt, err := db.Prepare("INSERT INTO part_time_under_18_salary (id_code_store, salary) VALUES(?,?)")
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

	result, err := stmt.Exec(IdCodeStore, Salary)
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

	w.Header().Set("Content-Type", "application/json", "*")
	json.NewEncoder(w).Encode(response)

}

func UpdatePartTimeUnder18Salary(w http.ResponseWriter, r *http.Request) {
	var response initialize.Response

	db := db.Connect()

	stmt, err := db.Prepare("UPDATE part_time_under_18_salary SET id_code_store = ?, salary = ? WHERE id_part_time_under_18_salary = ?")
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	idPartAboveSalary := keyVal["id_part_time_under_18_salary"]
	newIdCodeStore := keyVal["id_code_store"]
	newSalary := keyVal["salary"]

	id, err := strconv.Atoi(idPartAboveSalary)

	result, err := stmt.Exec(newIdCodeStore, newSalary, id)
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

func DeletePartTimeUnder18Salary(w http.ResponseWriter, r *http.Request) {

	db := db.Connect()
	params := mux.Vars(r)
	stmt, err := db.Prepare("DELETE FROM part_time_under_18_salary WHERE id_part_time_under_18_salary = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(params["id_part_time_under_18_salary"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Data Sudah Terhapus Dengan ID = ")

	w.Header().Set("Content-Type", "application/json", "*")
	json.NewEncoder(w).Encode(params["id_part_time_under_18_salary"])

}
