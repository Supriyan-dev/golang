package forgot_password

import (
	"../../db"
	"../../initialize"
	_models_forgot_password "../../models/users"
	_Response "../../response"
	"net/http"
)

func ReturnForgotPassword(w http.ResponseWriter, r *http.Request) {

	var _response initialize.ResponseMaster
	//json.NewDecoder(r.Body).Decode(&init_insert)
	_Email := r.FormValue("email")
	_Pin := r.FormValue("pin")
	db := db.Connect()
	if r.Method != "POST" {
		_response.Status = http.StatusMethodNotAllowed
		_response.Message = "Status Method Not Allowed"
		_response.Data = nil
		_Response.ResponseJson(w, _response.Status, _response)
	} else {
		_model := _models_forgot_password.Models_init_Users{DB: db}
		resultDataEmail, resultDataPin, err := _model.ForgotPassword(_Email, _Pin)
		defer db.Close()
		if err == nil {
			if resultDataPin == 0 && resultDataEmail == 0 {
				_response.Status = http.StatusBadRequest
				_response.Message = `data pin not found and data email not found`
				_response.Data = nil
				_Response.ResponseJson(w, _response.Status, _response)
			} else if resultDataEmail == 0 {
				_response.Status = http.StatusBadRequest
				_response.Message = `data email not found`
				_response.Data = nil
				_Response.ResponseJson(w, _response.Status, _response)
			} else if resultDataEmail == 1 {
				if resultDataPin == 0 && _Pin != "" {
					_response.Status = http.StatusBadRequest
					_response.Message = `data pin not found`
					_response.Data = nil
					_Response.ResponseJson(w, _response.Status, _response)
				} else {
					_response.Status = http.StatusOK
					_response.Message = "Success Response"
					_response.Data = map[string]string{
						"data_email": _Email,
					}
					_Response.ResponseJson(w, _response.Status, _response)
				}
			} else {
				_response.Status = http.StatusOK
				_response.Message = "Success Response"
				_response.Data = map[string]string{
					"data_email": _Email,
					"data_pin":   _Pin,
				}
				_Response.ResponseJson(w, _response.Status, _response)
			}
		}
	}
}

func ReturnForgotPasswordAction(w http.ResponseWriter, r *http.Request) {

	var _response initialize.ResponseMaster
	//json.NewDecoder(r.Body).Decode(&init_insert)
	_Email := r.FormValue("email")
	_Token := r.FormValue("token")
	_NewPassword := r.FormValue("new_password")
	db := db.Connect()
	if r.Method != "POST" {
		_response.Status = http.StatusMethodNotAllowed
		_response.Message = "Status Method Not Allowed"
		_response.Data = nil
		_Response.ResponseJson(w, _response.Status, _response)
	} else {
		_model := _models_forgot_password.Models_init_Users{DB: db}
		resultData, _ := _model.ForgotPasswordAction(_Email, _Token, _NewPassword)
		defer db.Close()
		if resultData == "Success Response" {
			_response.Status = http.StatusOK
			_response.Message = "Success Response"
			_response.Data = map[string]string{
				"data_email": _Email,
			}
			_Response.ResponseJson(w, _response.Status, _response)
		} else {
			_response.Status = http.StatusBadRequest
			_response.Message = resultData
			_response.Data = nil
			_Response.ResponseJson(w, _response.Status, _response)
		}
	}
}
