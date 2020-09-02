package login_controller

import (
	"encoding/json"
	"log"
	"net/http"

	// controller "../controller/data_master_controller"
	"../db"
	"../helpers"
	model1 "../model1/login"
	"github.com/gorilla/mux"
)

func GenerateHashPassword(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	hash, _ := helpers.HashPassword(vars["password"])

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(hash)

}

func CheckLogin(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		employee_number := r.FormValue("employee_number")
		password := r.FormValue("password")
		db := db.Connect()
		res, err := model1.CheckLoginUser(employee_number, password)

		if err != nil {
			log.Println(http.StatusInternalServerError, map[string]string{
				"messages": err.Error(),
			})
		}

		if !res {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode("response not Found")
		}

		if res {
			err2 := model1.ModelLogin_init{DB: db}
			model, errr := err2.ReturnAllDataUser()
			if model != nil {
				handler(w, r)
			} else if errr == nil {
				log.Println("gagal response")
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

// func LoginHandler(w http.ResponseWriter, r *http.Request) {

// 	w.Header().Set("Content-Type", "application/json")
// 	var user initialize.Login
// 	var result initialize.Login
// 	var res initialize.ResponseResult
// 	body, _ := ioutil.ReadAll(r.Body)
// 	err := json.Unmarshal(body, &user)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	db := db.Connect()

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	err = db.FindOne(context.TODO() db.D{{"employee_number", user.Employee_number}}).Decode(&result)
// 	result, err := db.

// 	if err != nil {
// 		res.Error = "Invalid employee_number"
// 		json.NewEncoder(w).Encode(res)
// 		return
// 	}

// 	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(user.Password))

// 	if err != nil {
// 		res.Error = "Invalid password"
// 		json.NewEncoder(w).Encode(res)
// 		return
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"employee_number": result.Employee_number,
// 		"password":        result.Password,
// 	})

// 	tokenString, err := token.SignedString([]byte("secret"))

// 	if err != nil {
// 		res.Error = "Error while generating token,Try again"
// 		json.NewEncoder(w).Encode(res)
// 		return
// 	}

// 	result.Token = tokenString
// 	result.Password = ""

// 	json.NewEncoder(w).Encode(result)

// }
