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

func ReturnAllPrefect(w http.ResponseWriter, r *http.Request) {
	var _response initialize.Response

	db := db.Connect()
	_con := model1.ModelPref_init{DB: db}
	ExcuteData, err := _con.ReturnAllDataPrefecture()
	if err != nil {
		panic(err.Error())
	}

	log.Println(ExcuteData)

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

func ReturnAllPrefectPagination(w http.ResponseWriter, r *http.Request) {
	var prefect initialize.Prefect
	var arrPrefect []initialize.Prefect
	var response initialize.Response

	db := db.Connect()
	defer db.Close()
	code := mux.Vars(r)

	totalDataPerPage, _ := strconv.Atoi(code["perPage"])
	page, _ := strconv.Atoi(code["page"])

	var totalData int
	err := db.QueryRow("SELECT COUNT(*) FROM prefecture").Scan(&totalData)

	totalPage := int(math.Ceil(float64(totalData) / float64(totalDataPerPage)))

	if page > totalPage {
		page = totalPage
	}
	if page <= 0 {
		page = 1
	}

	firstIndex := (totalDataPerPage * page) - totalDataPerPage

	query := fmt.Sprintf("select id_prefecture,ISO,prefecture_name from prefecture limit %d,%d", firstIndex, totalDataPerPage)

	rows, err := db.Query(query)
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&prefect.Id_prefecture, &prefect.ISO, &prefect.Prefecture_name); err != nil {
			log.Fatal(err.Error())

		} else {
			arrPrefect = append(arrPrefect, prefect)
		}
	}
	response.Status = 200
	response.Message = "Success"
	response.Data = arrPrefect
	response.TotalPage = totalPage
	response.CurrentPage = page

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func GetPrefect(w http.ResponseWriter, r *http.Request) {
	var _response initialize.Response
	db := db.Connect()

	_id := r.URL.Query().Get("id_prefecture")

	_con := model1.ModelPref_init{DB: db}
	ExcuteData, err := _con.GetDataPrefecture(_id)
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

func CreatePrefect(w http.ResponseWriter, r *http.Request) {
	var init_insert initialize.Prefect
	var _response initialize.Response
	json.NewDecoder(r.Body).Decode(&init_insert)
	db := db.Connect()

	_con := model1.ModelPref_init{DB: db}
	ExcuteData, _ := _con.InsertDataPrefecture(&init_insert)

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

func UpdatePrefect(w http.ResponseWriter, r *http.Request) {
	var _response initialize.Response
	var init_insert initialize.Prefect
	json.NewDecoder(r.Body).Decode(&init_insert)
	db := db.Connect()

	_con := model1.ModelPref_init{DB: db}
	ExcuteData, _ := _con.UpdateDataprefecture(&init_insert)

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

func DeletePrefect(w http.ResponseWriter, r *http.Request) {
	var _response initialize.Response
	var test initialize.Prefect
	json.NewDecoder(r.Body).Decode(&test)
	db := db.Connect()
	_con := model1.ModelPref_init{DB: db}
	ExcuteData, err := _con.DeleteDataPrefecture(&test)
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
			_response.Data = test.Id_prefecture
			response.ResponseJson(w, _response.Status, _response)
		}
	} else {
		_response.Status = http.StatusMethodNotAllowed
		_response.Message = "Sorry Your Method Missing Not Allowed"
		_response.Data = "Null"
		response.ResponseJson(w, _response.Status, _response)
	}

}
