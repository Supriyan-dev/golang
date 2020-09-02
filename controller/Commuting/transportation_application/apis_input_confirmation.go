package transportation_application

import (
	"../../../db"
	"../../../initialize"
	models_enter_the_information "../../../models/enter_the_information/transportation_application"
	_Response "../../../response"
	"github.com/gorilla/mux"
	"net/http"
)

func ReturnGetByCommutingInputConfirmation(w http.ResponseWriter, r *http.Request) {

	var _response initialize.ResponseMaster

	storeNumber := r.FormValue("store_number")
	employeeNumber := r.FormValue("employee_number")
	db := db.Connect()

	_model := models_enter_the_information.Models_init_input_confirmation{DB: db}
	ResultData, err := _model.GetDataInputConfimation(storeNumber, employeeNumber)

	if err != nil {
		_response.Status = http.StatusInternalServerError
		_response.Message = err.Error()
		_response.Data = nil
		_Response.ResponseJson(w, _response.Status, _response)
	} else if r.Method != "POST" {
		_response.Status = http.StatusMethodNotAllowed
		_response.Message = "Status Method Not Allowed"
		_response.Data = nil
		_Response.ResponseJson(w, _response.Status, _response)
	} else {
		_response.Status = http.StatusOK
		_response.Message = "Success Response"
		_response.Data = ResultData
		_Response.ResponseJson(w, _response.Status, _response)

	}

}

func ReturnSubmitInputConfirmation(w http.ResponseWriter, r *http.Request) {

	var _response initialize.ResponseMaster
	db := db.Connect()
	vars := mux.Vars(r)
	_id := vars["id_commuting_trip"]
	if r.Method != "PATCH" {
		_response.Status = http.StatusMethodNotAllowed
		_response.Message = "Status Method Not Allowed"
		_response.Data = nil
		_Response.ResponseJson(w, _response.Status, _response)
	} else {
		_model := models_enter_the_information.Models_init_input_confirmation{DB: db}
		resultData, err := _model.Model_SubmitInputConfirmation(_id)

		if err == "Success Response" && resultData > 0 {
			_response.Status = http.StatusOK
			_response.Message = err
			_response.Data = map[string]string{
				"id_commuting_trip": _id,
			}
			_Response.ResponseJson(w, _response.Status, _response)
		} else {
			_response.Status = http.StatusBadRequest
			_response.Message = "Please Check Your ID"
			_response.Data = nil
			_Response.ResponseJson(w, _response.Status, _response)
		}
	}
}