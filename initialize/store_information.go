package initialize
import (
	"../models"
)

type StoreInformation struct {
	Id_code_store int    `json:"id_code_store"`
	Code_store    models.NullString  `json:"code_store"`
	Store_name    string `json:"store_name"`
}
