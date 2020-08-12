package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jeffri/golang-test/db"
	"github.com/jeffri/golang-test/initialize"
)

func ReturnAllCommutingTrip(w http.ResponseWriter, r *http.Request) {
	var CommuntingTrip initialize.CommutingTrip
	var arrCommutingTrip []initialize.CommutingTrip
	var response initialize.Response

	db := db.Connect()

	rows, err := db.Query("SELECT * FROM commuting_trip")
	if err != nil {
		log.Print(err)
	}
	defer db.Close()

	for rows.Next() {
		if err := rows.Scan(&CommuntingTrip.Id_commuting_trip, &CommuntingTrip.Id_general_information, &CommuntingTrip.Route_profile_name, &CommuntingTrip.Date, &CommuntingTrip.Attendance_code, &CommuntingTrip.Save_draft_status, &CommuntingTrip.Status_approval, &CommuntingTrip.Reason, &CommuntingTrip.Draft, &CommuntingTrip.Save_trip, &CommuntingTrip.Submit, &CommuntingTrip.Date_submit, &CommuntingTrip.Time_submit, &CommuntingTrip.Date_time_approve, &CommuntingTrip.Code_commuting, &CommuntingTrip.Created_date, &CommuntingTrip.Created_time, &CommuntingTrip.Superior_verbal); err != nil {
			log.Fatal(err.Error())

		} else {
			arrCommutingTrip = append(arrCommutingTrip, CommuntingTrip)
		}
	}
	response.Status = 200
	response.Message = "Success"
	response.Data = arrCommutingTrip

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
