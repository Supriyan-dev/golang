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


func ReturnGetByCommutingUsageRecord(w http.ResponseWriter, r *http.Request) {

	var _response initialize.ResponseMaster

	storeNumber := r.FormValue("store_number")
	employeeNumber := r.FormValue("employee_number")
	db := db.Connect()

	_model := models_enter_the_information.Models_init_Usage_Record{DB: db}
	ResultData, err := _model.Model_GetByIdUsageRecord(storeNumber, employeeNumber)

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

func ReturnInsertUsageRecordApplyForTravelExpenses(w http.ResponseWriter, r *http.Request) {

	var init_insertUTP enter_the_information.InsertTransportationApplication
	var _response initialize.ResponseMaster
	json.NewDecoder(r.Body).Decode(&init_insertUTP)
	db := db.Connect()

	_model := models_enter_the_information.Models_init_Usage_Record{DB: db}
	resultData, err := _model.Model_InsertUsageRecordApplyForTravelExpenses(&init_insertUTP)

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
