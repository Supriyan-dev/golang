package data_master_controller

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	_Response "../../response"
	"../../db"
	"../../initialize"
	model1 "../../model1/data_master_model"
	"../../response"
	"github.com/gorilla/mux"
)

func ReturnAllDepartementInformation(w http.ResponseWriter, r *http.Request) {
	var _response initialize.Response

	db := db.Connect()
	_con := model1.ModelDept_init{DB: db}
	ExcuteData, err := _con.ReadDataDepartmentInformation()
	if err != nil {
		panic(err.Error())
	}

	if r.Method == "GET" {
		if ExcuteData == nil {
			_response.Status = http.StatusBadRequest
			_response.Message = "Sorry Your Input Missing Body Bad Request"
			_response.Data = "Null"
			response.ResponseJson(w, _response.Status, _response)
		} else {
			_response.Status = http.StatusOK
			_response.Message = "Success"
			_response.Data = ExcuteData
			response.ResponseJson(w, _response.Status, _response)
		}
	} else {
		_response.Status = http.StatusMethodNotAllowed
		_response.Message = "Sorry Your Method Missing Not Allowed"
		_response.Data = "Null"
		response.ResponseJson(w, _response.Status, _response)
	}

}

func SearchDataDepartmentInformation(w http.ResponseWriter, r *http.Request) {
	var _response initialize.Response
	db := db.Connect()
	type Name struct {
		Keyword string `json:"keyword"`
	}
	var Keyword Name
	json.NewDecoder(r.Body).Decode(&Keyword)
	_con := model1.ModelDept_init{DB: db}
	result, err := _con.SearchDepartmentInformationModels(Keyword.Keyword)
	if err != nil {
		panic(err.Error())
	}

	if r.Method == "POST" {
		if result == nil {
			_response.Status = http.StatusBadRequest
			_response.Message = "Sorry Your Input Missing Body Bad Request"
			_response.Data = "Null"
			_Response.ResponseJson(w, _response.Status, _response)
		} else {
			_response.Status = http.StatusOK
			_response.Message = "Success"
			_response.Data = result
			_Response.ResponseJson(w, _response.Status, _response)
		}
	} else {
		_response.Status = http.StatusMethodNotAllowed
		_response.Message = "Sorry Your Method Missing Not Allowed"
		_response.Data = "Null"
		_Response.ResponseJson(w, _response.Status, _response)
	}
}

func ReturnAllDepartementInformationPagination(w http.ResponseWriter, r *http.Request) {
	var DeptInfo initialize.DepartementInformation
	var arrDepartementInformation []initialize.DepartementInformation
	var _response initialize.Response
	db := db.Connect()
	defer db.Close()
	code := mux.Vars(r)

	totalDataPerPage, _ := strconv.Atoi(code["perPage"])
	page, _ := strconv.Atoi(code["page"])

	var totalData int
	err := db.QueryRow("SELECT COUNT(*) FROM department_information").Scan(&totalData)

	totalPage := int(math.Ceil(float64(totalData) / float64(totalDataPerPage)))

	if page > totalPage {
		page = totalPage
	}
	if page <= 0 {
		page = 1
	}
	firstIndex := (totalDataPerPage * page) - totalDataPerPage

	query := fmt.Sprintf("select id_department,department_code,department_name,id_code_store from department_information limit %d,%d", firstIndex, totalDataPerPage)

	rows, err := db.Query(query)
	if err != nil {
		log.Print(err)
	}
	for rows.Next() {
		if err := rows.Scan(&DeptInfo.Id_department, &DeptInfo.Department_code, &DeptInfo.Department_name, &DeptInfo.Id_code_store); err != nil {

			log.Fatal(err.Error())

		} else {
			arrDepartementInformation = append(arrDepartementInformation, DeptInfo)
		}
	}
	if r.Method == "GET" {
		if arrDepartementInformation != nil {
			_response.Status = http.StatusOK
			_response.Message = "Success"
			_response.TotalPage = totalPage
			_response.CurrentPage = page
			_response.Data = arrDepartementInformation
			response.ResponseJson(w, _response.Status, _response)
		} else if page > totalPage {
			_response.Status = http.StatusBadRequest
			_response.Message = "Sorry Your Input Missing Body Bad Request"
			_response.TotalPage = totalPage
			_response.CurrentPage = page
			_response.Data = "Null"
			response.ResponseJson(w, _response.Status, _response)
		}
	} else {
		_response.Status = http.StatusMethodNotAllowed
		_response.Message = "Sorry Your Method Missing Not Allowed"
		_response.TotalPage = totalPage
		_response.CurrentPage = page
		_response.Data = "Null"
		response.ResponseJson(w, _response.Status, _response)
	}
}

func GetDepartementInformation(w http.ResponseWriter, r *http.Request) {
	var _response initialize.Response
	db := db.Connect()

	_id := r.URL.Query().Get("id_department")

	_con := model1.ModelDept_init{DB: db}
	Execute, _ := _con.GetDataDepartmentInformation(_id)

	if r.Method == "GET" {
		if Execute == nil {
			_response.Status = http.StatusBadRequest
			_response.Message = "You entered the wrong body"
			_response.Data = "Null"
			response.ResponseJson(w, _response.Status, _response)
		} else {
			_response.Status = http.StatusOK
			_response.Message = "Success"
			_response.Data = Execute
			response.ResponseJson(w, _response.Status, _response)
		}
	} else {
		_response.Status = http.StatusMethodNotAllowed
		_response.Message = "You entered the wrong method"
		_response.Data = "Null"
		response.ResponseJson(w, _response.Status, _response)
	}

}

func CreateDepartementInformation(w http.ResponseWriter, r *http.Request) {
	var DepartI initialize.DepartementInformation
	var _response initialize.Response
	db := db.Connect()

	json.NewDecoder(r.Body).Decode(&DepartI)

	_con := model1.ModelDept_init{DB: db}
	Execute, _ := _con.InsertDataDepartmentInformation(&DepartI)

	if r.Method == "POST" {
		if Execute == nil {
			_response.Status = http.StatusBadRequest
			_response.Message = "You entered the wrong body"
			_response.Data = "Null"
			response.ResponseJson(w, _response.Status, _response)
		} else {
			_response.Status = http.StatusOK
			_response.Message = "Success"
			_response.Data = Execute
			response.ResponseJson(w, _response.Status, _response)
		}
	} else {
		_response.Status = http.StatusMethodNotAllowed
		_response.Message = "You entered the wrong method"
		_response.Data = "Null"
		response.ResponseJson(w, _response.Status, _response)
	}
}

func UpdateDepartementInformation(w http.ResponseWriter, r *http.Request) {
	var DepartI initialize.DepartementInformation
	var _response initialize.Response
	db := db.Connect()

	json.NewDecoder(r.Body).Decode(&DepartI)

	_con := model1.ModelDept_init{DB: db}
	Execute, _ := _con.UpdateDataDepartmentInformation(&DepartI)

	if r.Method == "PUT" {
		if Execute == nil {
			_response.Status = http.StatusBadRequest
			_response.Message = "You entered the wrong body"
			_response.Data = "Null"
			response.ResponseJson(w, _response.Status, _response)
		} else {
			_response.Status = http.StatusOK
			_response.Message = "Success"
			_response.Data = Execute
			response.ResponseJson(w, _response.Status, _response)
		}
	} else {
		_response.Status = http.StatusMethodNotAllowed
		_response.Message = "You entered the wrong method"
		_response.Data = "Null"
		response.ResponseJson(w, _response.Status, _response)
	}
}

func DeleteDepartementInformation(w http.ResponseWriter, r *http.Request) {
	var _response initialize.Response
	db := db.Connect()
	params := mux.Vars(r)
	delete := params["id_department"]
	stmt, err := db.Exec("DELETE FROM department_information WHERE id_department = ?", delete)
	if err != nil {
		panic(err.Error())
	}

	statment, err := stmt.RowsAffected()

	if r.Method == "DELETE" {
		if statment != 1 {
			_response.Status = http.StatusBadRequest
			_response.Message = "Sorry Your Input Missing Body Bad Request"
			_response.Data = nil
			response.ResponseJson(w, _response.Status, _response)
		} else {
			_response.Status = http.StatusOK
			_response.Message = "Success Data has been Deleted with ID"
			_response.Data = delete
			response.ResponseJson(w, _response.Status, _response)
		}
	} else {
		_response.Status = http.StatusMethodNotAllowed
		_response.Message = "Sorry Your Method Missing Not Allowed"
		_response.Data = nil
		response.ResponseJson(w, _response.Status, _response)
	}
}
