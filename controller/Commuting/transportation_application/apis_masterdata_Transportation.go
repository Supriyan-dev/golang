package transportation_application

import (
	"net/http"
	"../../../initialize"
	"../../../db"
	_Response "../../../response"
	models_MasterData "../../../models/Commuting/Master_data"
)

func ReturnGetDataMasterTransportation(w http.ResponseWriter, r *http.Request) {

	var _response initialize.ResponseMaster

	db := db.Connect()
	if r.Method != "GET" {
		_response.Status = http.StatusMethodNotAllowed
		_response.Message = "Status Method Not Allowed"
		_response.Data = nil
		_Response.ResponseJson(w, _response.Status, _response)
	} else {
		_model := models_MasterData.Init_DB_CommutingMaster_Data{DB: db}
		ResultData, err := _model.GetDataTransportation()
		defer db.Close()
		if err != nil {
			_response.Status = http.StatusBadRequest
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
}