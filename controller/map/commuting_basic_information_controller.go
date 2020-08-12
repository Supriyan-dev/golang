package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jeffri/golang-test/db"
	"github.com/jeffri/golang-test/initialize"
)

func ReturnAllCommutingBasicInformation(w http.ResponseWriter, r *http.Request) {
	var CommuntingBasic initialize.CommutingBasicInformation
	var arrCommutingBasicInformation []initialize.CommutingBasicInformation
	var response initialize.Response

	db := db.Connect()

	rows, err := db.Query("SELECT * FROM commuting_basic_information")
	if err != nil {
		log.Print(err)
	}
	defer db.Close()

	for rows.Next() {
		if err := rows.Scan(&CommuntingBasic.Id_commuting_basic_information, &CommuntingBasic.Id_general_information, &CommuntingBasic.Driver_license_document, &CommuntingBasic.Driver_license_document_url, &CommuntingBasic.Driver_license_expiry_date, &CommuntingBasic.Car_insurance_document, &CommuntingBasic.Car_insurance_document_url, &CommuntingBasic.Car_insurance_document_expiry_date, &CommuntingBasic.Daily_commuting_method, &CommuntingBasic.Default_transportation, &CommuntingBasic.Permitted_to_drive, &CommuntingBasic.Insurance_company, &CommuntingBasic.Personal_injury, &CommuntingBasic.Property_damage, &CommuntingBasic.Status_approve, &CommuntingBasic.Date_approve, &CommuntingBasic.Time_approve, &CommuntingBasic.Date_submit); err != nil {
			log.Fatal(err.Error())

		} else {
			arrCommutingBasicInformation = append(arrCommutingBasicInformation, CommuntingBasic)
		}
	}
	response.Status = 200
	response.Message = "Success"
	response.Data = arrCommutingBasicInformation

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
