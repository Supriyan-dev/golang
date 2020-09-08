package login

import (
	"database/sql"
	"fmt"

	"../../db"
	"../../helpers"
	"../../initialize"
	"../../models"
)

type ModelLoginWorkFlow_init models.DB_init

func CheckLoginUserWorkFlow(employee_number, password string) (bool, error) {
	var login initialize.Login
	var pwd string

	db := db.Connect()
	sqlStatement := "SELECT id_user, employee_number, password FROM user WHERE employee_number = ?"

	err := db.QueryRow(sqlStatement, employee_number).Scan(
		&login.Id_user, &login.Employee_number, &pwd,
	)

	if err == sql.ErrNoRows {
		fmt.Println("employee_number not found")
		return false, err
	}

	if err != nil {
		fmt.Println("Query Error")
		return false, err
	}

	match, err := helpers.CheckPasswordHash(password, pwd)
	if !match {
		fmt.Println("hash and password doesn't match")
		return false, err
	}

	return true, nil
}
