package controller

import (
	initialize2 "Go_DX_Services/initialize/map"
	"encoding/json"
	"log"
	"net/http"
	"Go_DX_Services/db"
	"Go_DX_Services/initialize"
)

func ReturnAllCategory_138(w http.ResponseWriter, r *http.Request) {
	var cat138 initialize2.Category_138
	var arrCategory_138 []initialize2.Category_138
	var response initialize.Response

	db := db.Connect()

	rows, err := db.Query("SELECT * FROM category_138")
	if err != nil {
		log.Print(err)
	}
	defer db.Close()

	for rows.Next() {
		if err := rows.Scan(&cat138.Id_data, &cat138.Go_to, &cat138.Toll_entrance, &cat138.Toll_exit, &cat138.Created_date, &cat138.Created_time, &cat138.Id_cash_claim); err != nil {
			log.Fatal(err.Error())

		} else {
			arrCategory_138 = append(arrCategory_138, cat138)
		}
	}
	response.Status = 200
	response.Message = "Success"
	response.Data = arrCategory_138

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
