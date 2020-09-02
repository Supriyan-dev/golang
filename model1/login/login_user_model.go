package login

import (
	"database/sql"
	"fmt"
	"log"

	"../../db"
	"../../helpers"
	"../../initialize"
	"../../models"
	"github.com/mervick/aes-everywhere/go/aes256"
)

type ModelLogin_init models.DB_init

func CheckLoginUser(employee_number, password string) (bool, error) {
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

func (model1 ModelLogin_init) ReturnAllDataUser() (arrAll []initialize.Users, err error) {
	var all initialize.Users

	rows, err := model1.DB.Query("SELECT * FROM user")
	if err != nil {
		log.Print(err)
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&all.Id_user, &all.First_name, &all.Last_name, &all.Employee_number, &all.Id_code_store, &all.Password, &all.Id_role, &all.Email, &all.Recovery_pin, &all.Photo_url, &all.Photo_name); err != nil {
			log.Fatal(err.Error())

		} else {
			arrAll = append(arrAll, all)
		}
	}

	return arrAll, nil

}

func (model1 ModelLogin_init) ReturnAllDataUserEncript() {
	// ** Hampton / BMG
	// ** 29 Agu 2020

	// Key_hash (salt) untuk digunakan pada proses enkripsi atau dekripsi json (salt jangan diubah)
	key := "P@ssw0rdL0g1n"

	// ********************* Contoh JSON ****************************************
	data := `{"id_user" : "id_user", "first_name":"first_name", "last_name":"last_name", "employee_number":"employee_number", "id_code_store":"id_code_store", "password":"password", "id_role":"id_role", "email":"email", "recovery_pin":"recovery_pin", "photo_url":"photo_url", "photo_name":"photo_name"}`

	// ***********************  Contoh hasil enkripsi ****************************
	// strings := "U2FsdGVkX1+A6fHQtv9uULjYK3tAlJMpYGwOgfobshi1ava/vEI+nkyAcLKFfS8lv5VJtOzb1rByx9CnOM+iVw=="
	// var encrypted string
	// **************** Proses Encryption **********************
	encrypted := aes256.Encrypt(data, key)

	// proses decryption
	// decrypted := aes256.Decrypt(strings, key)

	// **************** Data original (JSON) *****************
	// fmt.Println(data)

	// **************** OUTPUT Enkripsi **********************
	log.Println(encrypted)
	// return encrypted

	// **************** OUTPUT Dekripsi **********************
	// fmt.Println(decrypted)

}
