package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jeffri/golang-test/db"
	"github.com/jeffri/golang-test/initialize"
)

func ReturnAllCategory_136(w http.ResponseWriter, r *http.Request) {
	var cat134 initialize.Category_136
	var arrCategory_136 []initialize.Category_136
	var response initialize.Response

	db := db.Connect()

	rows, err := db.Query("SELECT * FROM category_136")
	if err != nil {
		log.Print(err)
	}
	defer db.Close()

	for rows.Next() {
		if err := rows.Scan(&cat134.Id_data, &cat134.Purpose, &cat134.Created_date, &cat134.Created_time, &cat134.Id_cash_claim); err != nil {
			log.Fatal(err.Error())

		} else {
			arrCategory_136 = append(arrCategory_136, cat134)
		}
	}
	response.Status = 200
	response.Message = "Success"
	response.Data = arrCategory_136

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
