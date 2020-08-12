package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jeffri/golang-test/db"
	"github.com/jeffri/golang-test/initialize"
)

func ReturnAllCategory_137(w http.ResponseWriter, r *http.Request) {
	var cat137 initialize.Category_137
	var arrCategory_137 []initialize.Category_137
	var response initialize.Response

	db := db.Connect()

	rows, err := db.Query("SELECT * FROM category_137")
	if err != nil {
		log.Print(err)
	}
	defer db.Close()

	for rows.Next() {
		if err := rows.Scan(&cat137.Id_data, &cat137.Go_to, &cat137.Created_date, &cat137.Created_time, &cat137.Parking_place, &cat137.Id_cash_claim); err != nil {
			log.Fatal(err.Error())

		} else {
			arrCategory_137 = append(arrCategory_137, cat137)
		}
	}
	response.Status = 200
	response.Message = "Success"
	response.Data = arrCategory_137

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
