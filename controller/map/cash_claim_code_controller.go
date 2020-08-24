package controller

import (
	initialize2 "Go_DX_Services/initialize/map"
	"encoding/json"
	"log"
	"net/http"
	"Go_DX_Services/db"
	"Go_DX_Services/initialize"
)

func ReturnAllCashClaimCode(w http.ResponseWriter, r *http.Request) {
	var cashCode initialize2.CashClaim
	var arrCashClaimCode []initialize2.CashClaimCode
	var response initialize.Response

	db := db.Connect()

	rows, err := db.Query("SELECT * FROM cash_claim_code")
	if err != nil {
		log.Print(err)
	}
	defer db.Close()

	for rows.Next() {
		if err := rows.Scan(&cashCode.Id_code, &cashCode.Submission_number, &cashCode.Created_date, &cashCode.Created_time, &cashCode.File_csv, &cashCode.Submit_to_approve, &cashCode.Date_submit, &cashCode.Lock_by_1, &cashCode.Lock_by_2, &cashCode.Lock_by_3); err != nil {
			log.Fatal(err.Error())

		} else {
			arrCashClaimCode = append(arrCashClaimCode, cashCode)
		}
	}
	response.Status = 200
	response.Message = "Success"
	response.Data = arrCashClaimCode

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
