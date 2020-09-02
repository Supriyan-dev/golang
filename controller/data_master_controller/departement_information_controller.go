package data_master_controller

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"

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

func ReturnAllDepartementInformationPagination(w http.ResponseWriter, r *http.Request) {
	var DeptInfo initialize.DepartementInformation
	var arrDepartementInformation []initialize.DepartementInformation
	var response initialize.Response
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
	response.Status = 200
	response.Message = "Success"
	response.Data = arrDepartementInformation
	response.TotalPage = totalPage
	response.CurrentPage = page

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

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
	var delete initialize.DepartementInformation
	json.NewDecoder(r.Body).Decode(&delete)
	db := db.Connect()
	_con := model1.ModelDept_init{DB: db}
	ExcuteData, err := _con.DeleteDataDepartmentInformation(&delete)
	if err != nil {
		panic(err.Error())
	}
	if r.Method == "DELETE" {
		if ExcuteData == nil {
			_response.Status = http.StatusBadRequest
			_response.Message = "Sorry Your Input Missing Body Bad Request"
			_response.Data = "Null"
			response.ResponseJson(w, _response.Status, _response)
		} else {
			_response.Status = http.StatusOK
			_response.Message = "Success Data has been Deleted with ID"
			_response.Data = delete.Id_department
			response.ResponseJson(w, _response.Status, _response)
		}
	} else {
		_response.Status = http.StatusMethodNotAllowed
		_response.Message = "Sorry Your Method Missing Not Allowed"
		_response.Data = "Null"
		response.ResponseJson(w, _response.Status, _response)
	}
}
