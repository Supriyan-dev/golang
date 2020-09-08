package transportation_application

import (
	"../../../db"
	"../../../initialize"
	"../../../initialize/Commuting"
	models_enter_the_information "../../../models/Commuting/transportation_application"
	_Response "../../../response"
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
	if r.Method != "POST" {
		_response.Status = http.StatusMethodNotAllowed
		_response.Message = "Status Method Not Allowed"
		_response.Data = nil
		_Response.ResponseJson(w, _response.Status, _response)
	} else {
		_model := models_enter_the_information.Models_init_Usage_Record{DB: db}
		ResultData, err := _model.Model_GetByIdUsageRecord(storeNumber, employeeNumber)
		defer _model.DB.Close()
		defer db.Close()
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
}

func ReturnGetByCommutingUsageRecordForEdit(w http.ResponseWriter, r *http.Request) {

	var _response initialize.ResponseMaster

	storeNumber := r.FormValue("store_number")
	employeeNumber := r.FormValue("employee_number")
	id_commuting_trip := r.FormValue("id_commuting_trip")
	db := db.Connect()
	if r.Method != "POST" {
		_response.Status = http.StatusMethodNotAllowed
		_response.Message = "Status Method Not Allowed"
		_response.Data = nil
		_Response.ResponseJson(w, _response.Status, _response)
	} else {
		_model := models_enter_the_information.Models_init_Usage_Record{DB: db}
		ResultData, err := _model.Model_GetByIdUsageRecordForEdit(storeNumber, employeeNumber, id_commuting_trip)
		defer db.Close()
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
}

func ReturnGetByCommutingUsageRecordUseMyRoute(w http.ResponseWriter, r *http.Request) {
	var _response initialize.ResponseMaster

	storeNumber := r.FormValue("store_number")
	employeeNumber := r.FormValue("employee_number")
	db := db.Connect()
	if r.Method != "POST" {
		_response.Status = http.StatusMethodNotAllowed
		_response.Message = "Status Method Not Allowed"
		_response.Data = nil
		_Response.ResponseJson(w, _response.Status, _response)
	} else {
		_model := models_enter_the_information.Models_init_Usage_Record{DB: db}
		ResultData, err := _model.Model_GetByIdUsageRecordUseMyRoute(storeNumber, employeeNumber)
		defer db.Close()
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
}

func ReturnGetByCommutingUsageRecordHistory(w http.ResponseWriter, r *http.Request) {
	var _response initialize.ResponseWithPagination

	storeNumber := r.FormValue("store_number")
	employeeNumber := r.FormValue("employee_number")
	page := r.FormValue("page")
	filter := r.FormValue("filter")
	showData := r.FormValue("show_data")
	searching := r.FormValue("searching")
	Mpage, _ := strconv.Atoi(page)
	var showDataint int
	if showData != "" {
		showDataint, _ = strconv.Atoi(showData)
	}
	db := db.Connect()
	if r.Method != "POST" {
		_response.Status = http.StatusMethodNotAllowed
		_response.CurrentPage = 0
		_response.TotalPage = 0
		_response.Message = "Status Method Not Allowed"
		_response.Data = nil
		_Response.ResponseJson(w, _response.Status, _response)
	} else {
		_model := models_enter_the_information.Models_init_Usage_Record{DB: db}
		ResultData, err, CountData := _model.Model_GetByIdUsageRecordHistory(storeNumber, employeeNumber, page, filter, showData, searching)
		defer db.Close()
		if err != nil {
			_response.Status = http.StatusInternalServerError
			_response.Message = err.Error()
			_response.CurrentPage = 0
			_response.TotalPage = 0
			_response.Data = nil
			_Response.ResponseJson(w, _response.Status, _response)
		} else {
			_response.Status = http.StatusOK
			_response.Message = "Success Response"
			_response.CountData = CountData
			_response.CurrentPage = Mpage
			_response.TotalPage = (CountData / showDataint) + 1
			_response.Data = ResultData
			_Response.ResponseJson(w, _response.Status, _response)

		}
	}
}

func ReturnInsertUsageRecordApplyForTravelExpenses(w http.ResponseWriter, r *http.Request) {
	var initializeData Commuting.InsertUsageRecordApplyForTravelExpenses
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
		defer db.Close()
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

	var initializeData Commuting.UpdateUsageRecordApplyForTravelExpenses
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
		defer db.Close()
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
		defer db.Close()
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
		defer db.Close()
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
	_date := vars["date"]
	if r.Method != "PATCH" {
		_response.Status = http.StatusMethodNotAllowed
		_response.Message = "Status Method Not Allowed"
		_response.Data = nil
		_Response.ResponseJson(w, _response.Status, _response)
	} else {
		_model := models_enter_the_information.Models_init_Usage_Record{DB: db}
		resultData, err := _model.Model_UseUsageRecord(_id, _date)
		defer db.Close()
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
