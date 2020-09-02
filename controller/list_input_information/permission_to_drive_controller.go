package list_input_information

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strconv"

	_Response "../../response"
	"github.com/gorilla/mux"

	"../../db"
	"../../initialize"
)

func PermissionToDrive(w http.ResponseWriter, r *http.Request) {
	var _response initialize.Response
	db := db.Connect()

	_model := models_init{DB: db}

	result, err := _model.ModelPermissionToDrive()

	if err == "Success Response" {
		_response.Status = http.StatusOK
		_response.Message = err
		_response.Data = result
		_Response.ResponseJson(w, _response.Status, _response)

	} else {
		_response.Status = http.StatusBadRequest
		_response.Message = err
		_response.Data = ""
		_Response.ResponseJson(w, _response.Status, _response)
	}

}

func PermissionToDrivePagination(w http.ResponseWriter, r *http.Request) {
	var join initialize.Join
	var arrJoin []initialize.Join
	var response initialize.Response

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
	query := fmt.Sprintf("SELECT store_information.code_store, basic_information.employee_code, basic_information.first_name, basic_information.last_name, commuting_basic_information.driver_license_expiry_date, commuting_basic_information.car_insurance_document_expiry_date  FROM general_information INNER JOIN store_information ON general_information.id_store_code = store_information.id_code_store INNER JOIN basic_information ON general_information.id_basic_information = basic_information.id_basic_information INNER JOIN commuting_basic_information ON commuting_basic_information.id_general_information = general_information.id_general_information LIMIT %d,%d", firstIndex, totalDataPerPage)

	result, err := db.Query(query)
	if err != nil {
		log.Print(err)
	}

	for result.Next() {
		if err := result.Scan(&join.Id_store_code, &join.Employee_code, &join.First_name, &join.Last_name, &join.Driver_license_expiry_date, &join.Car_insurance_document_expiry_date); err != nil {
			log.Fatal(err.Error())
		} else {
			arrJoin = append(arrJoin, join)
		}
	}

	response.Status = 200
	response.Message = "Success"
	response.Data = arrJoin
	w.Header().Set("Content-Type", "Aplication/json")
	json.NewEncoder(w).Encode(response)
}

func PermissionToDriveUpdate(w http.ResponseWriter, r *http.Request) {
	var response initialize.Response

	db := db.Connect()

	stmt, err := db.Prepare("UPDATE commuting_basic_information SET permitted_to_drive = ?, status_approve = ? WHERE id_commuting_basic_information = ?")
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	idCommutingBasic := keyVal["id_commuting_basic_information"]
	newPermitedToDrive := keyVal["permitted_to_drive"]
	newStatusApprove := keyVal["status_approve"]

	id, err := strconv.Atoi(idCommutingBasic)

	result, err := stmt.Exec(newPermitedToDrive, newStatusApprove, id)
	if err != nil {
		panic(err.Error())
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}

	response.Status = 200
	response.Message = "Success"
	response.Data = map[string]int64{
		"Data Yang Behasil Di Update": rowsAffected,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
