package Approve

import (
	"../../../db"
	"../../../initialize"
	"../../../initialize/Commuting"
	model_Approve "../../../models/Commuting/Approve"
	_Response "../../../response"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func ReturnGetDataApproveCommutingSumByAllEmployeeCode(w http.ResponseWriter, r *http.Request) {
	var _response initialize.ResponseWithPagination

	page := r.FormValue("page")
	filter := r.FormValue("filter")
	showData := r.FormValue("show_data")
	searching := r.FormValue("searching")
	condition := r.FormValue("condition")
	store_code := r.FormValue("store_code")
	department_code := r.FormValue("department_code")
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
		_model := model_Approve.Init_DB_CommutingApprove{DB: db}
		ResultData, err, CountData := _model.GetDataApproveCommutingSumByAllEmployeeCode(page, filter, showData, searching, condition, store_code,department_code)
		defer _model.DB.Close()
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

func ReturnGetDataApproveByCommutingEmployeeCode(w http.ResponseWriter, r *http.Request) {
	var _response initialize.ResponseWithPagination

	page := r.FormValue("page")
	employee_number := r.FormValue("employee_number")
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
		_model := model_Approve.Init_DB_CommutingApprove{DB: db}
		ResultData, err,CountData := _model.GetDataApproveByCommutingEmployeeCode(page, showData, searching, employee_number)
		defer db.Close()
		if err != nil {
			_response.Status = http.StatusInternalServerError
			_response.Message = err.Error()
			_response.CurrentPage = 0
			_response.TotalPage = 0
			_response.Data = nil
			_Response.ResponseJson(w, _response.Status, _response)
		} else {
			//CountData := models3.CheckDataByStoreAndEmployee(`select count(*) from (select COUNT(*)
			//								from commuting_trip comtrip, code_commuting cc,
			//								detail_commuting_trip detcomtrip, general_information geninfo, basic_information bainfo, store_information storeinfo
			//								where comtrip.id_commuting_trip = detcomtrip.id_commuting_trip and geninfo.id_general_information = comtrip.id_general_information AND
			//								geninfo.id_basic_information = bainfo.id_basic_information and geninfo.id_store_code = storeinfo.id_code_store  and storeinfo.code_store =? and cc.code_random = comtrip.code_commuting
			//								and bainfo.employee_code =? and comtrip.save_trip ='N' and comtrip.submit = 'Y'
			//								group by detcomtrip.id_commuting_trip order by comtrip.date asc) t`, storeNumber, employeeNumber)
			_response.Status = http.StatusOK
			_response.Message = "Success Response"
			_response.CurrentPage = Mpage
			_response.CountData = CountData
			_response.TotalPage = (CountData/ showDataint) +1
			_response.Data = ResultData
			_Response.ResponseJson(w, _response.Status, _response)

		}
	}
}

func ReturnDetailCommutingByEmployeeCode(w http.ResponseWriter, r *http.Request) {

	//var initializeData Commuting.Init_DetailCommutingByEmployeeCodeApprove
	var _response initialize.ResponseMaster
	//json.NewDecoder(r.Body).Decode(&initializeData)
	//change morning
	employee_number := r.FormValue("employee_number")
	id_basic_information := r.FormValue("id_basic_information")
	code_commuting := r.FormValue("code_commuting")
	db := db.Connect()
	if r.Method != "POST" {
		_response.Status = http.StatusMethodNotAllowed
		_response.Message = "Status Method Not Allowed"
		_response.Data = nil
		_Response.ResponseJson(w, _response.Status, _response)
	} else {
		_model := model_Approve.Init_DB_CommutingApprove{DB: db}
		//change morning
		resultData, err := _model.DetailCommutingByEmployeeCode(employee_number, id_basic_information, code_commuting)
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

func ReturnCommutingApproveOrReject(w http.ResponseWriter, r *http.Request) {

	//var initializeData Commuting.Init_InputDataApprove
	var initializeData []Commuting.Init_InputDataApprove
	var _response initialize.ResponseMaster
	json.NewDecoder(r.Body).Decode(&initializeData)
	param := mux.Vars(r)
	employee_number := param["employee_number"]
	id_basic_information := param["id_basic_information"]
	code_commuting := param["code_commuting"]
	db := db.Connect()
	if r.Method != "POST" {
		_response.Status = http.StatusMethodNotAllowed
		_response.Message = "Status Method Not Allowed"
		_response.Data = nil
		_Response.ResponseJson(w, _response.Status, _response)
	} else {
		_model := model_Approve.Init_DB_CommutingApprove{DB: db}
		//change morning
		resultData, err := _model.CommutingApproveOrReject(initializeData,employee_number,id_basic_information,code_commuting)
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
