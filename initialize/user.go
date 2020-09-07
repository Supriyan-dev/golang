package initialize

import "database/sql"

type Users struct {
	Id_user         int    `json:"id_user"`
	First_name      string `json:"first_name"`
	Last_name       string `json:"last_name"`
	Employee_number string `json:"employee_number"`
	Id_code_store   int    `json:"id_code_store"`
	Password        string `json:"password"`
	Id_role         int    `json:"id_role"`
	Email           string `json:"email"`
	Recovery_pin    string `json:"recovery_pin"`
	Photo_url       string `json:"photo_url"`
	Photo_name      string `json:"photo_name"`
}

type Login struct {
	Id_user         int    `json:"id_user"`
	Employee_number string `json:"employee_number"`
	Password        string `json:"password"`
}

type User struct {
	Id            int    `json:"id"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	Last_username string `json:"last_username"`
	Last_password string `json:"last_password"`
}

type UsersEncrypt struct {
	Id_user         string `json:"id_user"`
	First_name      string `json:"first_name"`
	Last_name       string `json:"last_name"`
	Employee_number string `json:"employee_number"`
	Id_code_store   string `json:"id_code_store"`
	Password        string `json:"password"`
	Id_role         string `json:"id_role"`
	Email           string `json:"email"`
	Recovery_pin    string `json:"recovery_pin"`
	Photo_url       string `json:"photo_url"`
	Photo_name      string `json:"photo_name"`
}

type NullString struct {
	Email        sql.NullString `json:"email"`
	Recovery_pin sql.NullString `json:"recovery_pin"`
	Photo_url    sql.NullString `json:"photo_url"`
	Photo_name   sql.NullString `json:"photo_name"`
}
