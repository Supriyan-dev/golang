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
