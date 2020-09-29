package login

import (
	"log"
	"../../db"
	"../../initialize"
	"../../models"
	"crypto/md5"
	"encoding/hex"
)

type ModelLogin_init models.DB_init

func CheckLoginUser(employee_number, password string) (bool, error) {
	var login initialize.Login
	var pwd string
	hasher := md5.New()
	hasher.Write([]byte(password))
	NewPasswordString := hex.EncodeToString(hasher.Sum(nil))
	log.Println(NewPasswordString)
	db := db.Connect()
	sqlStatement := "SELECT id_user, employee_number, password FROM user WHERE employee_number = ? and password = ?"

	err := db.QueryRow(sqlStatement, employee_number,NewPasswordString).Scan(
		&login.Id_user, &login.Employee_number, &pwd,
	)
	var datacheck int
	db.QueryRow(`SELECT count(*) FROM user WHERE employee_number = ? and password = ?`,employee_number,NewPasswordString).Scan(&datacheck)
	if datacheck <1 {
		return  false , err
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