package transportation_application

import (
	"../../../db"
	"../../../initialize"
	"../../../initialize/enter_the_information"
	models_enter_the_information "../../../models/enter_the_information/transportation_application"
	_Response "../../../response"
	models3 "../../../utils/enter_the_information"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
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

func ReturnGetByCommutingUsageRecordForEdit(w http.ResponseWriter, r *http.Request) {

	var _response initialize.ResponseMaster

	storeNumber := r.FormValue("store_number")
	employeeNumber := r.FormValue("employee_number")
	id_commuting_trip := r.FormValue("id_commuting_trip")
	db := db.Connect()

	_model := models_enter_the_information.Models_init_Usage_Record{DB: db}
	ResultData, err := _model.Model_GetByIdUsageRecordForEdit(storeNumber, employeeNumber, id_commuting_trip)

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

func ReturnGetByCommutingUsageRecordUseMyRoute(w http.ResponseWriter, r *http.Request) {
	var _response initialize.ResponseMaster

	storeNumber := r.FormValue("store_number")
	employeeNumber := r.FormValue("employee_number")
	db := db.Connect()
	_model := models_enter_the_information.Models_init_Usage_Record{DB: db}
	ResultData, err := _model.Model_GetByIdUsageRecordUseMyRoute(storeNumber, employeeNumber)

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

func ReturnGetByCommutingUsageRecordHistory(w http.ResponseWriter, r *http.Request) {
	var _response initialize.Response

	storeNumber := r.FormValue("store_number")
	employeeNumber := r.FormValue("employee_number")
	page := r.FormValue("page")
	filter := r.FormValue("filter")
	showData := r.FormValue("show_data")
	searching := r.FormValue("searching")
	Mpage, _ := strconv.Atoi(page)
	db := db.Connect()
	_model := models_enter_the_information.Models_init_Usage_Record{DB: db}
	ResultData, err := _model.Model_GetByIdUsageRecordHistory(storeNumber, employeeNumber, page, filter, showData, searching)

	if err != nil {
		_response.Status = http.StatusInternalServerError
		_response.Message = err.Error()
		_response.CurrentPage = 0
		_response.TotalPage = 0
		_response.Data = nil
		_Response.ResponseJson(w, _response.Status, _response)
	} else if r.Method != "POST" {
		_response.Status = http.StatusMethodNotAllowed
		_response.CurrentPage = 0
		_response.TotalPage = 0
		_response.Message = "Status Method Not Allowed"
		_response.Data = nil
		_Response.ResponseJson(w, _response.Status, _response)
	} else {
		CountData := models3.CheckDataByStoreAndEmployee(`select count(*) from (select COUNT(*)
										from commuting_trip comtrip, code_commuting cc,
										detail_commuting_trip detcomtrip, general_information geninfo, basic_information bainfo, store_information storeinfo
										where comtrip.id_commuting_trip = detcomtrip.id_commuting_trip and geninfo.id_general_information = comtrip.id_general_information AND
										geninfo.id_basic_information = bainfo.id_basic_information and geninfo.id_store_code = storeinfo.id_code_store  and storeinfo.code_store =? and cc.code_random = comtrip.code_commuting
										and bainfo.employee_code =? and comtrip.save_trip ='N' and comtrip.submit = 'Y'
										group by detcomtrip.id_commuting_trip order by comtrip.date asc) t`, storeNumber, employeeNumber)
		_response.Status = http.StatusOK
		_response.Message = "Success Response"
		_response.CurrentPage = Mpage
		_response.TotalPage = CountData
		_response.Data = ResultData
		_Response.ResponseJson(w, _response.Status, _response)

	}

}

func ReturnInsertUsageRecordApplyForTravelExpenses(w http.ResponseWriter, r *http.Request) {

	var initializeData enter_the_information.InsertUsageRecordApplyForTravelExpenses
	var _response initialize.ResponseMaster
	json.NewDecoder(r.Body).Decode(&initializeData)
	param := mux.Vars(r)
	con := param["condition"]
	employee_id := param["employee_id"]
	store_id := param["store_id"]

	db := db.Connect()
	if r.Method != "POST" {
		_response.Status = http.StatusMethodNotAllowed
		_response.Message = "Status Method Not Allowed"
		_response.Data = nil
		_Response.ResponseJson(w, _response.Status, _response)
	} else {
		_model := models_enter_the_information.Models_init_Usage_Record{DB: db}
		resultData, err := _model.Model_InsertUsageRecordApplyForTravelExpenses(con, store_id, employee_id, &initializeData)

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
}

func ReturnUpdateUsageRecordApplyForTravelExpenses(w http.ResponseWriter, r *http.Request) {

	var initializeData enter_the_information.UpdateUsageRecordApplyForTravelExpenses
	var _response initialize.ResponseMaster
	json.NewDecoder(r.Body).Decode(&initializeData)
	db := db.Connect()
	if r.Method != "PATCH" {
		_response.Status = http.StatusMethodNotAllowed
		_response.Message = "Status Method Not Allowed"
		_response.Data = nil
		_Response.ResponseJson(w, _response.Status, _response)
	} else {
		_model := models_enter_the_information.Models_init_Usage_Record{DB: db}
		resultData, err := _model.Model_UpdateUsageRecordApplyForTravelExpenses(&initializeData)

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
}

func ReturnDeleteUsageRecord(w http.ResponseWriter, r *http.Request) {

	var _response initialize.ResponseMaster
	db := db.Connect()
	vars := mux.Vars(r)
	_id := vars["id_commuting_trip"]
	if r.Method != "DELETE" {
		_response.Status = http.StatusMethodNotAllowed
		_response.Message = "Status Method Not Allowed"
		_response.Data = nil
		_Response.ResponseJson(w, _response.Status, _response)
	} else {
		_model := models_enter_the_information.Models_init_Usage_Record{DB: db}
		resultData, err := _model.Model_DeleteUsageRecordApplyForTravelExpenses(_id)

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

func ReturnUpdateUsageRecordDraft(w http.ResponseWriter, r *http.Request) {

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
		_model := models_enter_the_information.Models_init_Usage_Record{DB: db}
		resultData, err := _model.Model_UpdateUsageRecordDraft(_id)
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

func ReturnUseUsageRecord(w http.ResponseWriter, r *http.Request) {

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
		_model := models_enter_the_information.Models_init_Usage_Record{DB: db}
		resultData, err := _model.Model_UseUsageRecord(_id)
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
