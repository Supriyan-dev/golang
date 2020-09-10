package login

import (
	"database/sql"
	"fmt"

	"../../db"
	"../../helpers"
	"../../initialize"
	"../../models"
)

type ModelLogin_init models.DB_init

func CheckLoginUser(employee_number, password string) (bool, error) {
	var login initialize.Users
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

// func (model1 ModelLogin_init) ReadDataUserLogin(Employee_number, Password string) (all initialize.Users, err error) {
// 	rows, err := model1.DB.Query(`SELECT id_user, employee_number, password FROM user WHERE employee_number = ? AND password = ?`, Employee_number, Password)
// 	if err != nil {
// 		log.Print(err)
// 	}
// 	errScan := rows.Scan(&all.Employee_number, &all.Password)
// 	if errScan != nil {
// 		log.Println(errScan)
// 	}

// 	return all, nil
// }
