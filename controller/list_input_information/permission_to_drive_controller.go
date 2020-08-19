package initialize

import (
	"net/http"

	"github.com/jeffri/golang-test/Go_DX_Services/db"
	"github.com/jeffri/golang-test/Go_DX_Services/initialize"
)

func permissionToDrive(w http.ResponseWriter, r *http.Request) {
	var permission initialize.GeneralInformation
	var arrPermission initialize.GeneralInformation
	var response initialize.Response

	db, err := db.Connect()
	result, err := db.Query("SELECT store_information.id_store_code, store_information.code_store FROM store_information, general_information WHERE store_information.id_store_code = general_information.id_code_store")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	for result.Next() {
		if err := result.Scan(&permission.Id_code_store, &permission.Code_store){
			if err != nil {
				panic(err.Error())
			} else {
				arrPermission = append(arrPermission, GeneralInformation)
			}
		}
	}

	response.Status = 200
	response.Message = "Success"
	response.Data = arrPermission
	w.Header().Set("Content-Type", "Aplication/json", "*")
	json.NewEncoder(w).Encode(response)

}
