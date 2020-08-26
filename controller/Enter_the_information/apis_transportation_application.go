package enter_the_information

import (
	"Go_DX_Services/db"
	"Go_DX_Services/initialize"
	_Response "Go_DX_Services/response"
	"encoding/json"
	"net/http"
)

func ReturnCreateCommutingBasicInformation(w http.ResponseWriter, r *http.Request) {

	var init_insert InsertBasicInformation
	var _response initialize.Response
	json.NewDecoder(r.Body).Decode(&init_insert)
	db := db.Connect()

	_model := models_init{DB: db}
	resultData, err := _model.Model_InsertBasicInformation(&init_insert)

	if err == "Success Response" {
		_response.Status = http.StatusOK
		_response.Message = err
		_response.Data = resultData
		_Response.ResponseJson(w, _response.Status, _response)

	} else {
		_response.Status = http.StatusBadRequest
		_response.Message = err
		_response.Data = ""
		_Response.ResponseJson(w, _response.Status, _response)
	}

}

func ReturnGetByCommutingUsageRecordGet(w http.ResponseWriter, r *http.Request) {

	var _response initialize.Response

	db := db.Connect()

	_model := models_init{DB: db}
	ResultData, err := _model.Model_GetIdByCodeCommutingGet()

	if err != nil {
		_response.Status = http.StatusInternalServerError
		_response.Message = err.Error()
		_response.Data = ""
		_Response.ResponseJson(w, _response.Status, _response)
	} else {
		_response.Status = http.StatusOK
		_response.Message = "found data"
		_response.Data = ResultData
		_Response.ResponseJson(w, _response.Status, _response)
	}

}

func ReturnGetByCommutingUsageRecord(w http.ResponseWriter, r *http.Request) {

	var _response initialize.Response

	storeNumber := r.FormValue("store_number")
	employeeNumber := r.FormValue("employee_number")
	db := db.Connect()

	_model := models_init{DB: db}
	ResultData, err := _model.Model_GetIdByCodeCommuting(storeNumber, employeeNumber)

	if err != nil {
		_response.Status = http.StatusInternalServerError
		_response.Message = err.Error()
		_response.Data = ""
		_Response.ResponseJson(w, _response.Status, _response)
	} else {
		_response.Status = http.StatusOK
		_response.Message = "found data"
		_response.Data = ResultData
		_Response.ResponseJson(w, _response.Status, _response)

	}

}

func ReturnInsertUsageRecordApplyForTravelExpenses(w http.ResponseWriter, r *http.Request) {

	var init_insertUTP InsertTransportationApplication
	var _response initialize.Response
	json.NewDecoder(r.Body).Decode(&init_insertUTP)
	db := db.Connect()

	_model := models_init{DB: db}
	resultData, err := _model.Model_InsertUsageRecordApplyForTravelExpenses(&init_insertUTP)

	if err == "Success Response" {
		_response.Status = http.StatusOK
		_response.Message = err
		_response.Data = resultData
		_Response.ResponseJson(w, _response.Status, _response)

	} else {
		_response.Status = http.StatusBadRequest
		_response.Message = err
		_response.Data = ""
		_Response.ResponseJson(w, _response.Status, _response)
	}

}

func ReturnDetailInsertUsageRecordApplyForTravelExpenses(w http.ResponseWriter, r *http.Request) {

	var init_insertDetailUTP InsertDetailTransportationApplication
	var _response initialize.Response
	json.NewDecoder(r.Body).Decode(&init_insertDetailUTP)
	db := db.Connect()

	_model := models_init{DB: db}
	resultData, err := _model.Model_InsertDetailUsageRecordApplyForTravelExpenses(&init_insertDetailUTP)

	if err == "Success Response" {
		_response.Status = http.StatusOK
		_response.Message = err
		_response.Data = resultData
		_Response.ResponseJson(w, _response.Status, _response)
	} else {
		_response.Status = http.StatusBadRequest
		_response.Message = err
		_response.Data = ""
		_Response.ResponseJson(w, _response.Status, _response)
	}

}
