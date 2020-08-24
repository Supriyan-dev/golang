package Enter_the_information

import (
	"Go_DX_Services/db"
	"Go_DX_Services/initialize"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ReturnCreateCommutingBasicInformation(w http.ResponseWriter, r *http.Request) {

	var err error
	var response initialize.Response

	db := db.Connect()
	stmt, err := db.Prepare("INSERT INTO commuting_basic_information (insurance_company, driver_license_expiry_date, personal_injury, property_damage, car_insurance_document_expiry_date,id_general_information) VALUES(?,?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	Insurance_company := keyVal["insurance_company"]
	Driver_license_expiry_date := keyVal["driver_license_expiry_date"]
	Personal_injury := keyVal["personal_injury"]
	Property_damage := keyVal["property_damage"]
	Car_insurance_document_expiry_date := keyVal["car_insurance_document_expiry_date"]
	Id_general_information := keyVal["id_general_information"]

	result, err := stmt.Exec(Insurance_company,Driver_license_expiry_date,Personal_injury,Property_damage,Car_insurance_document_expiry_date,Id_general_information)
	if err != nil {
		panic(err.Error())
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}

	response.Status = 200
	response.Message = "Success"
	response.Data = map[string]int64{
		"Data baru telah dibuat": rowsAffected,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}