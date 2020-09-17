package login_controller

import (
	"encoding/json"
	"log"
	"net/http"

	// controller "../controller/data_master_controller"

	"../db"
	"../helpers"
	"../initialize"
	model1 "../model1/login"
	"../response"
	"github.com/gorilla/mux"
	"github.com/mervick/aes-everywhere/go/aes256"
)

func GenerateHashPasswordWorkFlow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	hash, _ := helpers.HashPassword(vars["password"])

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(hash)

}

func CheckLogin(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var _response initialize.Response
		type Login struct {
			Employee_number string `json:"employee_number"`
			Password        string `json:"password"`
		}
		type Baris struct {
			Data string `json:"data"`
		}
		var msg Baris
		json.NewDecoder(r.Body).Decode(&msg)
		key := "P@ssw0rdL0g1n"
		// json := []byte(b)
		
		// Unmarshal
		// var msg Baris
		//  json.Marshal(b, &msg)
		hasil := msg.Data
		
		// log.Println(hasil)
		// inputan := r.FormValue("data")
		decrypted := aes256.Decrypt(hasil, key)
		// log.Println(decrypted)
		jsonData := []byte(decrypted)

		var data Login

		err1 := json.Unmarshal(jsonData, &data)
		if err1 != nil {
			log.Println(err1)
		}

		// log.Println(decrypted)


		employee_number := data.Employee_number
		password := data.Password
		res, err := model1.CheckLoginUserWorkFlow(employee_number, password)

		if err != nil {
			log.Println(http.StatusInternalServerError, map[string]string{
				"messages": err.Error(),
			})
		}
		if !res {
			_response.Status = http.StatusBadRequest
			_response.Message = "Sorry Your Input Missing Body Bad Request"
			_response.Data = "Null"
			response.ResponseJson(w, _response.Status, _response)

		}
		if res {
			employee_number := data.Employee_number
		// password := data.Password
		// log.Println(employee_number,password)

			db := db.Connect()
			_con := model1.ModelLogin_init{DB: db}
			res, err := _con.ReadDataUserLogin(employee_number)
			if err != nil {
				panic(err.Error())
			}
			if r.Method == "POST" {
				_response.Status = http.StatusOK
				_response.Message = "Success"
				_response.Data = res
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

		// token := jwt.New(jwt.SigningMethodHS256)

		// claims := token.Claims.(jwt.MapClaims)
		// claims["username"] = username
		// claims["level"] = "aplications"
		// claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		// t, err := token.SignedString([]byte("secret"))
		// if err != nil {
		// 	w.Header().Set("Content-Type", "application/json")
		// 	json.NewEncoder(w).Encode("response not Found")
		// }

		// log.Println(http.StatusOK, map[string]string{
		// 	"token": t,
		// })

	}
}
