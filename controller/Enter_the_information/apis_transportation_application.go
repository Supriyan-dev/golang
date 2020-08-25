package enter_the_information

import (
	"Go_DX_Services/db"
	"Go_DX_Services/initialize"
	"encoding/json"
	"io/ioutil"
	"net/http"
	_Response "Go_DX_Services/response"
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
	response.Message = "Success Response"
	response.Data = map[string]int64{
		"Data baru telah dibuat": rowsAffected,
	}

	_Response.ResponseJson(w,response.Status,response)

}

func ReturnGetByCommutingUsageRecord(w http.ResponseWriter, r *http.Request){

	var _response initialize.Response

	storeNumber := r.FormValue("store_number")
	employeeNumber := r.FormValue("employee_number")
	db:= db.Connect()

	_model :=models_init{DB: db}
	ResultData,err := _model.Model_GetIdByCodeCommuting(storeNumber,employeeNumber)

	if err != nil {
		_response.Status = http.StatusInternalServerError
		_response.Message = err.Error()
		_response.Data = ""
		_Response.ResponseJson(w,_response.Status,_response)
	}else{
		_response.Status = http.StatusOK
		_response.Message = "found data"
		_response.Data = ResultData
		_Response.ResponseJson(w,_response.Status,_response)

	}

}

func ReturnInsertUsageRecordApplyForTravelExpenses(w http.ResponseWriter, r *http.Request) {

	var init_insertUTP InsertTransportationApplication
	var _response initialize.Response
	json.NewDecoder(r.Body).Decode(&init_insertUTP)
	db := db.Connect()

	_model := models_init{DB: db}
	resultData, err := _model.Model_InsertTransportationApplication(&init_insertUTP)

	if err == "Missing required field in body request" {
		_response.Status = http.StatusBadRequest
		_response.Message = err
		_response.Data = ""
		_Response.ResponseJson(w,_response.Status,_response)
	}else if err == "Success Response"{
		_response.Status = http.StatusOK
		_response.Message = err
		_response.Data = resultData
		_Response.ResponseJson(w,_response.Status,_response)

	}else{
		_response.Status = http.StatusUnauthorized
		_response.Message = err
		_response.Data = ""
		_Response.ResponseJson(w,_response.Status,_response)
	}

}

func ReturnDetailInsertUsageRecordApplyForTravelExpenses(w http.ResponseWriter, r *http.Request) {

	var init_insertDetailUTP InsertDetailTransportationApplication
	var _response initialize.Response
	json.NewDecoder(r.Body).Decode(&init_insertDetailUTP)
	db := db.Connect()

	_model := models_init{DB: db}
	resultData, err := _model.Model_InsertDetailTransportationApplication(&init_insertDetailUTP)

	if resultData == nil && err == "Missing required field in body request" {
		_response.Status = http.StatusBadRequest
		_response.Message = err
		_response.Data = ""
		_Response.ResponseJson(w,_response.Status,_response)
	}else if err == "Success Response"{
		_response.Status = http.StatusOK
		_response.Message = err
		_response.Data = resultData
		_Response.ResponseJson(w,_response.Status,_response)

	}else{
		_response.Status = http.StatusUnauthorized
		_response.Message = err
		_response.Data = ""
		_Response.ResponseJson(w,_response.Status,_response)
	}

}
