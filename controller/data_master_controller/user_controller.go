package data_master_controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jeffri/golang-test/db"
	"github.com/jeffri/golang-test/initialize"
)

func ReturnAllUser(w http.ResponseWriter, r *http.Request) {
	var user initialize.Users
	var arrUsers []initialize.Users
	var response initialize.Response

	db, err := db.Connect()

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

	code := mux.Vars(r)
	fmt.Fprintf(w, "Category: %v\n", code["page"])

	db, err := db.Connect()
	rows, err := db.Query("SELECT * FROM user ORDER BY id_user LIMIT " + code["page"] + " OFFSET 0")
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
