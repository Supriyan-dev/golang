package initialize

import "database/sql"

type Users struct {
	Id_user         int            `json:"id_user"`
	First_name      string         `json:"first_name"`
	Last_name       string         `json:"last_name"`
	Employee_number string         `json:"employee_number"`
	Id_code_store   int            `json:"id_code_store"`
	Password        string         `json:"password"`
	Id_role         int            `json:"id_role"`
	Email           sql.NullString `json:"email"`
	Recovery_pin    sql.NullString `json:"recovery_pin"`
	Photo_url       sql.NullString `json:"photo_url"`
	Photo_name      sql.NullString `json:"photo_name"`
}
