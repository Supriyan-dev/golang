package controller

import (
	initialize2 "Go_DX_Services/initialize/map"
	"encoding/json"
	"log"
	"net/http"
	"Go_DX_Services/db"
	"Go_DX_Services/initialize"
)

func ReturnAllCashClaimJoin(w http.ResponseWriter, r *http.Request) {
	var cashJoin initialize2.CashClaimJoin
	var arrCashClaimJoin []initialize2.CashClaimJoin
	var response initialize.Response

	db := db.Connect()

	rows, err := db.Query("SELECT * FROM cash_claim_join")
	if err != nil {
		log.Print(err)
	}
	defer db.Close()

	for rows.Next() {
		if err := rows.Scan(&cashJoin.Id_join, &cashJoin.Created_date, &cashJoin.Created_time, &cashJoin.Employee_code, &cashJoin.Id_cash_claim); err != nil {
			log.Fatal(err.Error())

		} else {
			arrCashClaimJoin = append(arrCashClaimJoin, cashJoin)
		}
	}
	response.Status = 200
	response.Message = "Success"
	response.Data = arrCashClaimJoin

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
