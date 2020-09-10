package login_controller

import (
	"net/http"

	"../../db"
	"../../initialize"
	model1 "../../model1/data_master_model"
	"../../response"
)

func WorkFlowLogin(w http.ResponseWriter, r *http.Request) {
	var _response initialize.Response
	employee_number := r.FormValue("employee_number")
	// password := r.FormValue("password")
	db := db.Connect()
	_con := model1.ModelUser_init{DB: db}
	ExcuteData, err := _con.ReadDataUserLogin(employee_number)
	if err != nil {
		panic(err.Error())
	}
	// key := "P@ssw0rdL0g1n"

	if r.Method == "POST" {
		_response.Status = http.StatusOK
		_response.Message = "Success"
		_response.Data = ExcuteData
		response.ResponseJson(w, _response.Status, _response)
		// data := []byte(aes256.Decrypt(ExcuteData, key))
		// w.Write(data)
	} else {
		_response.Status = http.StatusMethodNotAllowed
		_response.Message = "Sorry Your Method Missing Not Allowed"
		_response.Data = "Null"
		response.ResponseJson(w, _response.Status, _response)
	}

}
