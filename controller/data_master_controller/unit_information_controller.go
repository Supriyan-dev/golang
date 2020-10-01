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

func ReturnAllUnitInformation(w http.ResponseWriter, r *http.Request) {
	var _response initialize.Response

	db := db.Connect()
	_con := model1.ModelUnit_init{DB: db}
	ExcuteData, err := _con.ReturnAllDataUnitInformation()
	if err != nil {
		log.Println(err.Error())
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

func SearchDataUnitInformation(w http.ResponseWriter, r *http.Request) {
	var _response initialize.ResponseSearch
	db := db.Connect()
	type Name struct {
		Keyword  string `json:"keyword"`
		Page     int    `json:"page"`
		Show_data int    `json:"show_data"`
	}
	var Keyword Name
	json.NewDecoder(r.Body).Decode(&Keyword)
	_con := model1.ModelUnit_init{DB: db}
	result, err, totalDataSearch := _con.SearchUnitInformationModels(Keyword.Keyword, Keyword.Page, Keyword.Show_data)
	if err != nil {
		log.Println(err.Error())
	}

	if r.Method == "POST" {
		if result == nil {
			var _response initialize.ResponseDataNull
			_response.Status = http.StatusBadRequest
			_response.Message = "Unit Information Data doesn't exists."
			response.ResponseJson(w, _response.Status, _response)
		} else {
			_response.Status = http.StatusOK
			_response.Message = "Success"
			_response.TotalData = totalDataSearch
			_response.TotalPage = (totalDataSearch / Keyword.Show_data) + 1
			_response.Data = result
			response.ResponseJson(w, _response.Status, _response)
		}
	} else {
		var _response initialize.ResponseDataNull
		_response.Status = http.StatusMethodNotAllowed
		_response.Message = "Sorry Your Method Missing Not Allowed"
		response.ResponseJson(w, _response.Status, _response)
	}
}


func ReturnAllUnitInformationPagination(w http.ResponseWriter, r *http.Request) {
	var unit initialize.UnitInformation
	var arrUnitInformation []initialize.UnitInformation
	var _response initialize.Response

	db := db.Connect()
	defer db.Close()
	code := mux.Vars(r)

	totalDataPerPage, _ := strconv.Atoi(code["perPage"])
	page, _ := strconv.Atoi(code["page"])

	var totalData int
	err := db.QueryRow("SELECT COUNT(*) FROM unit_information").Scan(&totalData)

	totalPage := int(math.Ceil(float64(totalData) / float64(totalDataPerPage)))

	if page <= 0 {
		page = 1
	}

	firstIndex := (totalDataPerPage * page) - totalDataPerPage

	query := fmt.Sprintf("select id_unit,unit_code,unit_name from unit_information limit %d,%d", firstIndex, totalDataPerPage)

	rows, err := db.Query(query)
	if err != nil {
		log.Print(err)
	}
	for rows.Next() {
		if err := rows.Scan(&unit.Id_unit, &unit.Unit_code, &unit.Unit_name); err != nil {

			log.Fatal(err.Error())

		} else {
			arrUnitInformation = append(arrUnitInformation, unit)
		}
	}
	if r.Method == "GET" {
		if arrUnitInformation != nil {
			_response.Status = http.StatusOK
			_response.Message = "Success"
			_response.TotalPage = totalPage
			_response.CurrentPage = page
			_response.Data = arrUnitInformation
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

func GetUnitInformation(w http.ResponseWriter, r *http.Request) {
	var _response initialize.Response
	db := db.Connect()

	_id := r.URL.Query().Get("id_unit")

	_con := model1.ModelUnit_init{DB: db}
	ExcuteData, err := _con.GetDataUnitInformation(_id)
	if err != nil {
		log.Println(err.Error())
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

func CreateUnitInformation(w http.ResponseWriter, r *http.Request) {
	var init_insert initialize.UnitInformation
	var _response initialize.Response
	json.NewDecoder(r.Body).Decode(&init_insert)
	db := db.Connect()

	_con := model1.ModelUnit_init{DB: db}
	ExcuteData, _ := _con.InsertDataUnitInformation(&init_insert)

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

func UpdateUnitInformation(w http.ResponseWriter, r *http.Request) {
	var _response initialize.Response
	var init_insert initialize.UnitInformation
	json.NewDecoder(r.Body).Decode(&init_insert)
	db := db.Connect()

	_con := model1.ModelUnit_init{DB: db}
	ExcuteData, _ := _con.UpdateDataUnitInformation(&init_insert)

	if r.Method == "PUT" {
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

func DeleteUnitInformation(w http.ResponseWriter, r *http.Request) {
	var _response initialize.Response
	db := db.Connect()
	params := mux.Vars(r)
	delete := params["id_unit"]
	stmt, err := db.Exec("DELETE FROM unit_information WHERE id_unit = ?", delete)
	if err != nil {
		log.Println(err.Error())
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
