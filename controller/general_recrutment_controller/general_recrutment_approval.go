package general_recrutment_controller

import (
	"encoding/json"
	"net/http"

	"../../db"
	initialize "../../initialize/general_recrutment"
	model1 "../../model1/general_recrutment_model"
	response "../../response"
)

func DataGeneralRecrutment(w http.ResponseWriter, r *http.Request) {
	var init_insert initialize.GeneralRecrutmentJoin
	var _response initialize.Response
	json.NewDecoder(r.Body).Decode(&init_insert)
	db := db.Connect()

	_con := model1.ModelGeneral_init{DB: db}
	ExcuteData, _ := _con.InsertDataGeneralRecrutment(&init_insert)

	if r.Method == "POST" {
		if ExcuteData == nil {
			_response.Status = http.StatusBadRequest
			_response.Message = "Sorry Your Input Missing Body Bad Request"
			_response.Data = "Null"
			response.ResponseJson(w, _response.Status, _response)
		} else {
			_response.Status = http.StatusOK
			_response.Message = "Success"
			_response.Data = init_insert
			response.ResponseJson(w, _response.Status, _response)
		}
	} else {
		_response.Status = http.StatusMethodNotAllowed
		_response.Message = "Sorry Your Method Missing Not Allowed"
		_response.Data = "Null"
		response.ResponseJson(w, _response.Status, _response)
	}

}
