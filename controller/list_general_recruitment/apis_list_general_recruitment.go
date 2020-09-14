package list_general_recruitment

import (
	"../../db"
	"../../initialize"
	models_list_GR "../../models/list_general_recruitment"
	_Response "../../response"
	"net/http"
	"strconv"
)

func ReturnGetAllDataByStatus(w http.ResponseWriter, r *http.Request) {
	var _response initialize.ResponseWithPagination

	_status := r.FormValue("status")
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
		_model := models_list_GR.Models_init_listGeneralRecruitment{DB: db}
		ResultData, err, CountData := _model.GetListGeneralRecruitment(_status, page, filter, showData, searching)
		defer db.Close()
		if err != nil {
			_response.Status = http.StatusBadRequest
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

func ReturnGetAllDataDetailByStatus(w http.ResponseWriter, r *http.Request) {
	var _response initialize.ResponseMaster
	IdBasicInformation := r.FormValue("id_basic_information")
	EmployeeType := r.FormValue("employee_type")

	db := db.Connect()
	if r.Method != "POST" {
		_response.Status = http.StatusMethodNotAllowed
		_response.Message = "Status Method Not Allowed"
		_response.Data = nil
		_Response.ResponseJson(w, _response.Status, _response)
	} else {
		_model := models_list_GR.Models_init_listGeneralRecruitment{DB: db}

		toInt, errToInt := strconv.Atoi(IdBasicInformation)

		if errToInt != nil {
			_response.Status = http.StatusBadRequest
			_response.Message = errToInt.Error()
			_response.Data = nil
			_Response.ResponseJson(w, _response.Status, _response)
		}

		ResultData, Condition := _model.GetDetailListGeneralRecruitment(toInt ,EmployeeType)
		defer db.Close()
		if Condition == "Success Response" {
			_response.Status = http.StatusOK
			_response.Message = "Success Response"
			_response.Data = ResultData
			_Response.ResponseJson(w, _response.Status, _response)
		}else {
			_response.Status = http.StatusBadRequest
			_response.Message = Condition
			_response.Data = nil
			_Response.ResponseJson(w, _response.Status, _response)
		}
	}
}

