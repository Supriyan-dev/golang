package data_master_controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strconv"

	"../../db"
	"../../initialize"
	"../../model1"
	"../../response"
	"github.com/gorilla/mux"
)

func ReturnAllStoreInformation(w http.ResponseWriter, r *http.Request) {
	var storeInformation initialize.StoreInformation
	var arrStoreInformation []initialize.StoreInformation
	var response initialize.Response

	db := db.Connect()

	rows, err := db.Query("SELECT * FROM store_information")
	if err != nil {
		log.Print(err)
	}
	defer db.Close()

	for rows.Next() {
		if err := rows.Scan(&storeInformation.Id_code_store, &storeInformation.Code_store, &storeInformation.Store_name); err != nil {
			log.Fatal(err.Error())

		} else {
			arrStoreInformation = append(arrStoreInformation, storeInformation)
		}
	}
	response.Status = 200
	response.Message = "Success"
	response.Data = arrStoreInformation

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func ReturnAllStoreInformationPagination(w http.ResponseWriter, r *http.Request) {
	var storeInformation initialize.StoreInformation
	var arrStoreInformation []initialize.StoreInformation
	var response initialize.Response

	db := db.Connect()
	defer db.Close()
	code := mux.Vars(r)

	totalDataPerPage, _ := strconv.Atoi(code["perPage"])
	page, _ := strconv.Atoi(code["page"])

	var totalData int
	err := db.QueryRow("SELECT COUNT(*) FROM store_information").Scan(&totalData)

	totalPage := int(math.Ceil(float64(totalData) / float64(totalDataPerPage)))

	if page > totalPage {
		page = totalPage
	}
	if page <= 0 {
		page = 1
	}

	firstIndex := (totalDataPerPage * page) - totalDataPerPage

	query := fmt.Sprintf("select id_code_store,code_store,store_name from store_information limit %d,%d", firstIndex, totalDataPerPage)

	rows, err := db.Query(query)
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&storeInformation.Id_code_store, &storeInformation.Code_store, &storeInformation.Store_name); err != nil {
			log.Fatal(err.Error())
		} else {
			arrStoreInformation = append(arrStoreInformation, storeInformation)
		}
	}

	response.Status = 200
	response.Message = "success"
	response.Data = arrStoreInformation
	response.TotalPage = totalPage
	response.CurrentPage = page

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetStoreInformation(w http.ResponseWriter, r *http.Request) {

	var _response initialize.Response
	db := db.Connect()
	_id := r.URL.Query().Get("id_code_store")

	_con := model1.Models_init{DB: db}
	ExcuteData, err := _con.GetIdStoreInformation(_id)
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
		_response.Message = "data ketemu"
		_response.Data = "Null"
		response.ResponseJson(w, _response.Status, _response)
	}

}

// func CreateStoreInformation(w http.ResponseWriter, r *http.Request) {
// 	var storeS initialize.StoreInformation
// 	var _response initialize.Response

// 	db := db.Connect()

// 	_model := models_init{DB: db}
// 	json.NewDecoder(r.Body).Decode(&storeS)
// 	result, err := _model.InsertStoreInformation(&storeS)

// 	if errExcuteData != nil {
// 		_response.Status = http.StatusMethodNotAllowed
// 		_response.Message = "Sorry Your Input Method Allowed"
// 		_response.Data = "Null"
// 		api.ResponeWithJson(response, _response.Status, _response)
// 	}
// 	if ExcuteData == nil {
// 		_response.Status = http.StatusBadRequest
// 		_response.Message = "Sorry Your Input Missing Body Bad Request"
// 		_response.Data = "Null"
// 		api.ResponeWithJson(response, _response.Status, _response)
// 	} else {
// 		_response.Status = http.StatusOK
// 		_response.Message = "Success"
// 		_response.Data = arrStoreInformation
// 		api.ResponeWithJson(response, _response.Status, _response)
// 	}

// }

func UpdateStoreInformation(w http.ResponseWriter, r *http.Request) {
	var response initialize.Response

	db := db.Connect()

	stmt, err := db.Prepare("UPDATE store_information SET code_store = ?, store_name = ? WHERE id_code_store = ?")
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	idCode := keyVal["id_code_store"]
	newCode := keyVal["code_store"]
	newName := keyVal["store_name"]

	id, err := strconv.Atoi(idCode)

	result, err := stmt.Exec(newCode, newName, id)
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

func DeleteStoreInformation(w http.ResponseWriter, r *http.Request) {

	db := db.Connect()
	params := mux.Vars(r)
	stmt, err := db.Prepare("DELETE FROM store_information WHERE id_code_store = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(params["id_code_store"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Data Sudah Terhapus Dengan ID = ")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(params["id_code_store"])

}
