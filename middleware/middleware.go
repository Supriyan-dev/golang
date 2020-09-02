package middleware

import (
	"encoding/json"
	"log"
	"net/http"

	model1 "../model1/login"
)

func CheckLogin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := r.FormValue("username")
		password := r.FormValue("password")

		res, err := model1.CheckLoginUser(username, password)

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
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode("berhasil login")
		}
	}
}
