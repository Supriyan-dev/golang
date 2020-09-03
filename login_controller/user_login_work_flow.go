package login_controller

import (
	"encoding/json"
	"log"
	"net/http"

	// controller "../controller/data_master_controller"

	"../helpers"
	"../initialize"
	model1 "../model1/login"
	"../response"
	"github.com/gorilla/mux"
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

		employee_number := r.FormValue("employee_number")
		password := r.FormValue("password")
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
			handler(w, r)
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
