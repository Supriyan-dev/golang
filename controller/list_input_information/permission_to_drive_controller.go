package List_input_information

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jeffri/golang-test/Go_DX_Services/db"
	"github.com/jeffri/golang-test/Go_DX_Services/initialize"
)

func PermissionToDrive(w http.ResponseWriter, r *http.Request) {
	var join initialize.Join
	var arrJoin []initialize.Join
	var response initialize.Response

	db := db.Connect()
	result, err := db.Query("SELECT store_information.code_store, basic_information.employee_code, basic_information.first_name, basic_information.last_name, commuting_basic_information.driver_license_expiry_date, commuting_basic_information.car_insurance_document_expiry_date  FROM (((general_information INNER JOIN store_information ON general_information.id_store_code = store_information.id_code_store)INNER JOIN basic_information ON general_information.id_basic_information = basic_information.id_basic_information)INNER JOIN commuting_basic_information ON commuting_basic_information.id_general_information = general_information.id_general_information)")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	for result.Next() {
		if err := result.Scan(&join.Id_store_code, &join.Employee_code, &join.First_name, &join.Last_name, &join.Driver_license_expiry_date, &join.Car_insurance_document_expiry_date); err != nil {
			log.Fatal(err.Error())

		} else {
			arrJoin = append(arrJoin, join)
		}
	}

	response.Status = 200
	response.Message = "Success"
	response.Data = arrJoin
	w.Header().Set("Content-Type", "Aplication/json")
	json.NewEncoder(w).Encode(response)

}
