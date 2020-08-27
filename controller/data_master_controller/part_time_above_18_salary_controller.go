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

func ReturnAllPartTimeAbove18Salary(w http.ResponseWriter, r *http.Request) {
	var partTimeSalary initialize.PartTimeAbove18Salary
	var arrPartTimeAbove18Salary []initialize.PartTimeAbove18Salary
	var response initialize.Response

	db := db.Connect()

	rows, err := db.Query("SELECT * FROM part_time_above_18_salary")

	if err != nil {
		log.Print(err)
	}
	defer db.Close()

	for rows.Next() {
		if err := rows.Scan(&partTimeSalary.Id_part_time_above_18_salary, &partTimeSalary.Id_code_store, &partTimeSalary.Day_salary, &partTimeSalary.Night_salary, &partTimeSalary.Morning_salary, &partTimeSalary.Peek_time_salary); err != nil {

			log.Fatal(err.Error())

		} else {
			arrPartTimeAbove18Salary = append(arrPartTimeAbove18Salary, partTimeSalary)
		}
	}
	response.Status = 200
	response.Message = "Success"
	response.Data = arrPartTimeAbove18Salary

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func ReturnAllPartTimeAbove18SalaryPagination(w http.ResponseWriter, r *http.Request) {
	var partTimeSalary initialize.PartTimeAbove18Salary
	var arrPartTimeAbove18Salary []initialize.PartTimeAbove18Salary
	var response initialize.Response

	db := db.Connect()
	defer db.Close()
	code := mux.Vars(r)

	totalDataPerPage, _ := strconv.Atoi(code["perPage"])
	page, _ := strconv.Atoi(code["page"])

	var totalData int
	err := db.QueryRow("SELECT COUNT(*) FROM part_time_above_18_salary").Scan(&totalData)

	totalPage := int(math.Ceil(float64(totalData) / float64(totalDataPerPage)))

	if page > totalPage {
		page = totalPage
	}
	if page <= 0 {
		page = 1
	}

	firstIndex := (totalDataPerPage * page) - totalDataPerPage

	query := fmt.Sprintf("select id_part_time_above_18_salary,id_code_store,day_salary,night_salary,morning_salary,peek_time_salary from part_time_above_18_salary limit %d,%d", firstIndex, totalDataPerPage)

	rows, err := db.Query(query)
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&partTimeSalary.Id_part_time_above_18_salary, &partTimeSalary.Id_code_store, &partTimeSalary.Day_salary, &partTimeSalary.Night_salary, &partTimeSalary.Morning_salary, &partTimeSalary.Peek_time_salary); err != nil {

			log.Fatal(err.Error())

		} else {
			arrPartTimeAbove18Salary = append(arrPartTimeAbove18Salary, partTimeSalary)
		}
	}
	response.Status = 200
	response.Message = "Success"
	response.Data = arrPartTimeAbove18Salary
	response.TotalPage = totalPage
	response.CurrentPage = page

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func GetPartTimeAbove18Salary(w http.ResponseWriter, r *http.Request) {
	var partTimeAbove18Salary initialize.PartTimeAbove18Salary
	var arrPartTimeAbove18Salary []initialize.PartTimeAbove18Salary
	var response initialize.Response

	db := db.Connect()
	code := mux.Vars(r)

	result, err := db.Query("SELECT id_part_time_above_18_salary, id_code_store, day_salary,night_salary, morning_salary, peek_time_salary FROM part_time_above_18_salary WHERE id_part_time_above_18_salary = ?", code["id_part_time_above_18_salary"])
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {

		err := result.Scan(&partTimeAbove18Salary.Id_part_time_above_18_salary, &partTimeAbove18Salary.Id_code_store, &partTimeAbove18Salary.Day_salary, &partTimeAbove18Salary.Night_salary, &partTimeAbove18Salary.Morning_salary, &partTimeAbove18Salary.Peek_time_salary)
		if err != nil {
			panic(err.Error())
		} else {
			arrPartTimeAbove18Salary = append(arrPartTimeAbove18Salary, partTimeAbove18Salary)
		}
	}
	response.Status = 200
	response.Message = "Success"
	response.Data = arrPartTimeAbove18Salary

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func CreatePartTimeAbove18Salary(w http.ResponseWriter, r *http.Request) {
	var err error
	var response initialize.Response

	db := db.Connect()
	stmt, err := db.Prepare("INSERT INTO part_time_above_18_salary (id_code_store, day_salary, night_salary, morning_salary, peek_time_salary) VALUES(?,?,?,?,?)")
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
	DaySalary := keyVal["day_salary"]
	NightSalary := keyVal["night_salary"]
	MorningSalary := keyVal["morning_salary"]
	PeekTimeSalary := keyVal["peek_time_salary"]

	result, err := stmt.Exec(IdCodeStore, DaySalary, NightSalary, MorningSalary, PeekTimeSalary)
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

func UpdatePartTimeAbove18Salary(w http.ResponseWriter, r *http.Request) {
	var response initialize.Response

	db := db.Connect()

	stmt, err := db.Prepare("UPDATE part_time_above_18_salary SET id_code_store = ?, day_salary = ?, night_salary = ?, morning_salary = ?, peek_time_salary = ? WHERE id_part_time_above_18_salary = ?")
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	idPartAboveSalary := keyVal["id_part_time_above_18_salary"]
	newIdCodeStore := keyVal["id_code_store"]
	newDaySalary := keyVal["day_salary"]
	newNightSalary := keyVal["night_salary"]
	newMorningSalary := keyVal["morning_salary"]
	NewPeekTimeSalary := keyVal["peek_time_salary"]

	id, err := strconv.Atoi(idPartAboveSalary)

	result, err := stmt.Exec(newIdCodeStore, newDaySalary, newNightSalary, newMorningSalary, NewPeekTimeSalary, id)
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

func DeletePartTimeAbove18Salary(w http.ResponseWriter, r *http.Request) {

	db := db.Connect()
	params := mux.Vars(r)
	stmt, err := db.Prepare("DELETE FROM part_time_above_18_salary WHERE id_part_time_above_18_salary = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(params["id_part_time_above_18_salary"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Data Sudah Terhapus Dengan ID = ")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(params["id_part_time_above_18_salary"])

}
