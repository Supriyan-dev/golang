package initialize

import (
	"database/sql"
)

type BasicInformation struct {
	Id_basic_information int            `json :"id_basic_information"`
	Employee_code        int            `json :"employee_code"`
	First_name           string         `json :"first_name"`
	Last_name            string         `json :"last_name"`
	Gender               string         `json :"gender"`
	Birthdate            string         `json :"birthdate"`
	Add_postal_code      sql.NullString `json :"add_postal_code"`
	Id_prefecture        int            `json :"id_prefecture"`
	Adress               sql.NullString `json :"addres"`
	Adress_kana          sql.NullString `json :"adress_kana"`
	Adress_detail        sql.NullString `json :"adress_detail"`
	Adress_detail_kana   sql.NullString `json :"adress_detail_kana"`
	Add_phone_number     sql.NullString `json :"add_phone_number"`
	Marital_status       string         `json :"marital_status"`
	Dormitory_status     string         `json :"dormitory_status"`
}
