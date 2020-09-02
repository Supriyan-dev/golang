package controller

import (
	initialize2 "Go_DX_Services/initialize/map"
	"encoding/json"
	"log"
	"net/http"
	"Go_DX_Services/db"
	"Go_DX_Services/initialize"
)

func ReturnAllCashClaim(w http.ResponseWriter, r *http.Request) {
	var cash initialize2.CashClaim
	var arrCashClaim []initialize2.CashClaim
	var response initialize.Response

	db := db.Connect()

	rows, err := db.Query("SELECT * FROM cash_claim")
	if err != nil {
		log.Print(err)
	}
	defer db.Close()

	for rows.Next() {
		if err := rows.Scan(&cash.Id_cash_claim, &cash.Code_store, &cash.Employee_code, &cash.Division, &cash.Employee_name, &cash.Description, &cash.Reason, &cash.Amount, &cash.Gbr_bill, &cash.Exp_category, &cash.Bsn_partner, &cash.Pymt_method, &cash.Ppl_joined, &cash.Join_kasumi, &cash.Created_at, &cash.Created_time, &cash.Submission_number, &cash.Approve_level_1, &cash.Approve_date_level_1, &cash.Approve_time_level_1, &cash.Approve_level_2, &cash.Approve_date_level_2, &cash.Approve_time_level_2, &cash.Approve_level_3, &cash.Approve_date_level_3, &cash.Approve_time_level_3, &cash.Reason_level_1, &cash.Reason_level_2, &cash.Reason_level_3, &cash.Approve_code_1, &cash.Approve_name_1, &cash.Approve_code_2, &cash.Approve_name_2, &cash.Approve_code_3, &cash.Approve_name_3, &cash.Name_of_representative_from_client, &cash.Number_persons_from_client, &cash.Purpose_joined, &cash.Client_name); err != nil {
			log.Fatal(err.Error())

		} else {
			arrCashClaim = append(arrCashClaim, cash)
		}
	}
	response.Status = 200
	response.Message = "Success"
	response.Data = arrCashClaim

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
