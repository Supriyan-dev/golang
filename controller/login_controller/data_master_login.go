package login_controller

import (
	"encoding/json"
	"log"
	"net/http"

	"../../db"
	"../../initialize"
	model1 "../../model1/data_master_model"
	"../../response"
	"github.com/mervick/aes-everywhere/go/aes256"
)

func DataMasterLogin(w http.ResponseWriter, r *http.Request) {
	var _response initialize.Response
	type Login struct {
		Employee_number string
		Password        string
	}

	key := "P@ssw0rdL0g1n"

	inputan := r.FormValue("data")
	decrypted := aes256.Decrypt(inputan, key)

	jsonData := []byte(decrypted)

	var data Login

	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		log.Println(err)
	}
	employee_number := data.Employee_number
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
		// data := []byte(aes256.Encrypt(ExcuteData, key))
		// w.Write(data)
	} else {
		_response.Status = http.StatusMethodNotAllowed
		_response.Message = "Sorry Your Method Missing Not Allowed"
		_response.Data = "Null"
		response.ResponseJson(w, _response.Status, _response)
	}

}
