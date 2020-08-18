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

func ReturnAllUser(w http.ResponseWriter, r *http.Request) {
	var user initialize.Users
	var arrUsers []initialize.Users
	var response initialize.Response

	db := db.Connect()

	rows, err := db.Query("SELECT * FROM user")
	if err != nil {
		log.Print(err)
	}
	defer db.Close()

	for rows.Next() {
		if err := rows.Scan(&user.Id_user, &user.First_name, &user.Last_name, &user.Employee_number, &user.Id_code_store, &user.Password, &user.Id_role, &user.Email, &user.Recovery_pin, &user.Photo_url, &user.Photo_name); err != nil {
			log.Fatal(err.Error())

		} else {
			arrUsers = append(arrUsers, user)
		}
	}
	response.Status = 200
	response.Message = "Success"
	response.Data = arrUsers

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func ReturnAllUserPagination(w http.ResponseWriter, r *http.Request) {
	var user initialize.Users
	var arrUsers []initialize.Users
	var response initialize.Response

	db := db.Connect()
	defer db.Close()
	code := mux.Vars(r)

	totalDataPerPage, _ := strconv.Atoi(code["perPage"])
	page, _ := strconv.Atoi(code["page"])

	var totalData int
	err := db.QueryRow("SELECT COUNT(*) FROM user").Scan(&totalData)

	totalPage := int(math.Ceil(float64(totalData) / float64(totalDataPerPage)))

	if page > totalPage {
		page = totalPage
	}
	if page <= 0 {
		page = 1
	}

	firstIndex := (totalDataPerPage * page) - totalDataPerPage

	query := fmt.Sprintf("select Id_user,first_name,last_name,employee_number,id_code_store,password,id_role,email,recovery_pin,photo_url,photo_name from user limit %d,%d", firstIndex, totalDataPerPage)

	rows, err := db.Query(query)
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&user.Id_user, &user.First_name, &user.Last_name, &user.Employee_number, &user.Id_code_store, &user.Password, &user.Id_role, &user.Email, &user.Recovery_pin, &user.Photo_url, &user.Photo_name); err != nil {
			log.Fatal(err.Error())

		} else {
			arrUsers = append(arrUsers, user)
		}
	}
	response.Status = 200
	response.Message = "Success"
	response.Data = arrUsers
	response.TotalPage = totalPage
	response.CurrentPage = page

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func GetUser(w http.ResponseWriter, r *http.Request) {
	var user initialize.Users
	var arrUser []initialize.Users
	var response initialize.Response

	db := db.Connect()
	code := mux.Vars(r)
	fmt.Fprintf(w, "Category: %v\n", code["id_user"])

	result, err := db.Query("SELECT id_user, first_name, last_name, employee_number, id_code_store, password, id_role, email, recovery_pin, photo_url, photo_name FROM user WHERE id_user = ?", code["id_user"])
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {

		err := result.Scan(&user.Id_user, &user.First_name, &user.Last_name, &user.Employee_number, &user.Id_code_store, &user.Password, &user.Id_role, &user.Email, &user.Recovery_pin, &user.Photo_url, &user.Photo_name)
		if err != nil {
			panic(err.Error())
		} else {
			arrUser = append(arrUser, user)
		}
	}
	response.Status = 200
	response.Message = "Success"
	response.Data = arrUser

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var err error
	var response initialize.Response

	db := db.Connect()
	stmt, err := db.Prepare("INSERT INTO user (first_name, last_name, employee_number, id_code_store, password, id_role, email, recovery_pin, photo_url, photo_name) VALUES(?,?,?,?,?,?,?,?,?,?)")
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
	FirstName := keyVal["first_name"]
	LastName := keyVal["last_name"]
	EmployeeNumber := keyVal["employee_number"]
	IdCodeStore := keyVal["id_code_store"]
	Password := keyVal["password"]
	IdRole := keyVal["id_role"]
	Email := keyVal["email"]
	RecoveryPin := keyVal["recovery_pin"]
	PhotoUrl := keyVal["photo_url"]
	PhotoName := keyVal["photo_name"]

	result, err := stmt.Exec(FirstName, LastName, EmployeeNumber, IdCodeStore, Password, IdRole, Email, RecoveryPin, PhotoUrl, PhotoName)
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

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var response initialize.Response

	db := db.Connect()

	stmt, err := db.Prepare("UPDATE user SET first_name = ?, last_name = ?, employee_number = ?, id_code_store = ?, password = ?, id_role = ?, email = ?, recovery_pin = ?, photo_url = ?, photo_name = ? WHERE id_user = ?")
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	idUser := keyVal["id_user"]
	NewFirstName := keyVal["first_name"]
	NewLastName := keyVal["last_name"]
	NewEmployeeNumber := keyVal["employee_number"]
	NewIdCodeStore := keyVal["id_code_store"]
	NewPassword := keyVal["password"]
	NewIdRole := keyVal["id_role"]
	NewEmail := keyVal["email"]
	NewRecoveryPin := keyVal["recovery_pin"]
	NewPhotoUrl := keyVal["photo_url"]
	NewPhotoName := keyVal["photo_name"]
	id, err := strconv.Atoi(idUser)

	result, err := stmt.Exec(NewFirstName, NewLastName, NewEmployeeNumber, NewIdCodeStore, NewPassword, NewIdRole, NewEmail, NewRecoveryPin, NewPhotoUrl, NewPhotoName, id)
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

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	db := db.Connect()
	params := mux.Vars(r)
	stmt, err := db.Prepare("DELETE FROM user WHERE id_user = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(params["id_user"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Data Sudah Terhapus Dengan ID = ")

	w.Header().Set("Content-Type", "Aplication/json")
	json.NewEncoder(w).Encode(params["id_user"])

}
