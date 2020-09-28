package login

import (
	"database/sql"
	"fmt"
"log"
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

func (model1 ModelLogin_init) ReadDataUserLogin(Employee_number string) (all initialize.Users, err error) {
	var allNull initialize.NullString
	rows, err := model1.DB.Query(`SELECT id_user, first_name, last_name, employee_number, id_code_store, password, 
	id_role, email, recovery_pin, photo_url, photo_name FROM user WHERE employee_number = ?`, Employee_number)
	if err != nil {
		log.Print(err)
	}
	for rows.Next() {
		errScan := rows.Scan(&all.Id_user, &all.First_name, &all.Last_name, &all.Employee_number, &all.Id_code_store, &all.Password, &all.Id_role, &allNull.Email, &allNull.Recovery_pin, &allNull.Photo_url, &allNull.Photo_name)
		if errScan != nil {
			log.Println(errScan)
		}
	}
	return all, nil
}
