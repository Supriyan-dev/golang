package list_input_information

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	_Response "../../response"
	"github.com/gorilla/mux"
	"../../db"
	initialize "../../initialize/permission_to_drive"
	model1 "../../model1/permission_to_drive"
)

func PermissionToDrive(w http.ResponseWriter, r *http.Request) {
	var _response initialize.Response
	db := db.Connect()
	_con := model1.ModelsPermission_init{DB: db}
	result, err := _con.ModelPermissionToDrive()
	if err != nil {
		log.Println(err.Error())
	}

	if r.Method == "GET" {
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

func PermissionToDriveSearch(w http.ResponseWriter, r *http.Request) {
	var _response initialize.Response
	db := db.Connect()
	type Name struct {
		Keyword string `json:"keyword"`
	}
	var Keyword Name
	json.NewDecoder(r.Body).Decode(&Keyword)
	_con := model1.ModelsPermission_init{DB: db}
	result, err := _con.ModelPermissionToDriveSearch(Keyword.Keyword)
	if err != nil {
		log.Println(err.Error())
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

func PermissionToDrivePagination(w http.ResponseWriter, r *http.Request) {
	var join initialize.Join
	var arrJoin []initialize.Join
	var _response initialize.Response

	db := db.Connect()
	defer db.Close()
	code := mux.Vars(r)

	totalDataPerPage, _ := strconv.Atoi(code["perPage"])
	page, _ := strconv.Atoi(code["page"])

	var totalData int
	err := db.QueryRow("SELECT COUNT(*) FROM general_information").Scan(&totalData)

	totalPage := int(math.Ceil(float64(totalData) / float64(totalDataPerPage)))
	if page > totalPage {
		page = totalPage
	}
	if page <= 0 {
		page = 1
	}

	firstIndex := (totalDataPerPage * page) - totalDataPerPage
	query := fmt.Sprintf(`SELECT store_information.id_code_store, store_information.code_store, basic_information.employee_code, basic_information.first_name, 
	basic_information.last_name, commuting_basic_information.driver_license_expiry_date, commuting_basic_information.car_insurance_document_expiry_date, 
	commuting_basic_information.insurance_company, commuting_basic_information.personal_injury, commuting_basic_information.property_damage, commuting_basic_information.status_approve
	FROM general_information INNER JOIN store_information ON general_information.id_store_code = store_information.id_code_store 
	INNER JOIN basic_information ON general_information.id_basic_information = basic_information.id_basic_information 
	INNER JOIN commuting_basic_information ON commuting_basic_information.id_general_information = general_information.id_general_information LIMIT %d, %d`, firstIndex, totalDataPerPage)

	result, err := db.Query(query)
	if err != nil {
		log.Print(err)
	}

	for result.Next() {
		if err := result.Scan(&join.Id_code_store, &join.Code_store, &join.Employee_code,
			&join.First_name, &join.Last_name, &join.Driver_license_expiry_date, &join.Car_insurance_document_expiry_date, &join.Insurance_company,
			 &join.Personal_injury, &join.Property_damage, &join.Status_approve); err != nil {
			log.Fatal(err.Error())
		} else {
			arrJoin = append(arrJoin, join)
		}
	}

	if r.Method == "GET" {
		if arrJoin != nil {
			_response.Status = http.StatusOK
			_response.Message = "Success"
			_response.TotalPage = totalPage
			_response.CurrentPage = page
			_response.Data = arrJoin
			_Response.ResponseJson(w, _response.Status, _response)
		} else {
			_response.Status = http.StatusBadRequest
			_response.Message = "Sorry Your Input Missing Body Bad Request"
			_response.TotalPage = totalPage
			_response.CurrentPage = page
			log.Println(totalDataPerPage)
			_response.Data = "Null"
			_Response.ResponseJson(w, _response.Status, _response)
		}
	} else {
		_response.Status = http.StatusMethodNotAllowed
		_response.Message = "Sorry Your Method Missing Not Allowed"
		_response.TotalPage = totalPage
		_response.CurrentPage = page
		_response.Data = "Null"
		_Response.ResponseJson(w, _response.Status, _response)
	}
}

func PermissionToDriveUpdate(w http.ResponseWriter, r *http.Request) {
	var _response initialize.Response
	var init_insert initialize.UpdatePermissionToDrive
	db := db.Connect()
	json.NewDecoder(r.Body).Decode(&init_insert)
	_con := model1.ModelsPermission_init{DB: db}
	result, err := _con.UpdateDataPermissionToDrive(&init_insert)
	if err != nil {
		log.Println(err)
	}
	if r.Method == "PUT" {
		if result != nil {
			_response.Status = http.StatusOK
			_response.Message = "Success"
			_response.Data = result
			_Response.ResponseJson(w, _response.Status, _response)
		} else {
			_response.Status = http.StatusBadRequest
			_response.Message = "Sorry Your Input Missing Body Bad Request"
			_response.Data = "Null"
			_Response.ResponseJson(w, _response.Status, _response)
		}
	} else {
		_response.Status = http.StatusMethodNotAllowed
		_response.Message = "Sorry Your Method Missing Not Allowed"
		_response.Data = "Null"
		_Response.ResponseJson(w, _response.Status, _response)
	}
}
