package transportation_application

import (
	models_enter_the_information "../../../models/enter_the_information/transportation_application"
	"../../../db"
	"../../../initialize"
	_Response "../../../response"
	"encoding/json"
	"net/http"
	"../../../initialize/enter_the_information"
)



func ReturnCreateCommutingBasicInformation(w http.ResponseWriter, r *http.Request) {

	var init_insert enter_the_information.InsertBasicInformation
	var _response initialize.ResponseMaster
	json.NewDecoder(r.Body).Decode(&init_insert)
	db := db.Connect()
	_model := models_enter_the_information.Models_init_basic_information{DB: db}
	resultData, err := _model.Model_InsertBasicInformation(&init_insert)

	if err == "Success Response" {
		_response.Status = http.StatusOK
		_response.Message = err
		_response.Data = resultData
		_Response.ResponseJson(w, _response.Status, _response)

	} else {
		_response.Status = http.StatusBadRequest
		_response.Message = err
		_response.Data = nil
		_Response.ResponseJson(w, _response.Status, _response)
	}

}

func ReturnGetByCommutingBasicInformation(w http.ResponseWriter, r *http.Request) {

	var _response initialize.ResponseMaster

	storeNumber := r.FormValue("store_number")
	employeeNumber := r.FormValue("employee_number")
	db := db.Connect()

	_model := models_enter_the_information.Models_init_basic_information{DB: db}
	ResultData, err := _model.Model_GetByIdCommutingBasicInformation(storeNumber, employeeNumber)

	if err != nil {
		_response.Status = http.StatusInternalServerError
		_response.Message = err.Error()
		_response.Data = nil
		_Response.ResponseJson(w, _response.Status, _response)
	} else {
		_response.Status = http.StatusOK
		_response.Message = "Success Response"
		_response.Data = ResultData
		_Response.ResponseJson(w, _response.Status, _response)

	}

}
