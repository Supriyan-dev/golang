package initialize
import (
	"../models"
)

type StoreInformation struct {
	Latitude string `json:"latitude"`
	Longitude string `json:"longitude"`
	Id_code_store int    `json:"id_code_store"`
	Code_store    models.NullString  `json:"code_store"`
	Store_name    string `json:"store_name"`
}

type Filter struct {
	Id_code_store int    `json:"id_code_store"`
	First_name string    `json:"first_name"`
	Last_name string    `json:"fast_name"`
	Employee_number    string `json:"employee_number"`
}
