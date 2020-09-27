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

func WorkFlowLogin(w http.ResponseWriter, r *http.Request) {
	var _response initialize.Response
	type Login struct {
		Employee_number string
		Password        string
	}
	type Baris struct {
		Data string `json:"data"`
	}
	var msg Baris
	json.NewDecoder(r.Body).Decode(&msg)
	key := "P@ssw0rdL0g1n"
	hasil := msg.Data
	decrypted := aes256.Decrypt(hasil, key)
	log.Println(decrypted)
	jsonData := []byte(decrypted)

	var data Login

	err1 := json.Unmarshal(jsonData, &data)
	if err1 != nil {
		log.Println(err1)
	}

	employee_number := data.Employee_number
	password := data.Password
	db := db.Connect()
	_con := model1.ModelUser_init{DB: db}
	ExcuteData, err := _con.ReadDataUserLogin(employee_number, password)
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
