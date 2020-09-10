package main

import (
	"database/sql"
	"fmt"

	"github.com/mervick/aes-everywhere/go/aes256"
)

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

func main() {

	// ** Hampton / BMG
	// ** 29 Agu 2020

	// Key_hash (salt) untuk digunakan pada proses enkripsi atau dekripsi json (salt jangan diubah)
	// key := "P@ssw0rdL0g1n"
	key := "P@ssw0rdL0g1n"
	// ********************* Contoh JSON ****************************************
	data := `{"id_user" : "id_user", "first_name":"first_name", "last_name":"last_name", "employee_number":"employee_number", "id_code_store":"id_code_store", "password":"password", "id_role":"id_role", "email":"email", "recovery_pin":"recovery_pin", "photo_url":"photo_url", "photo_name":"photo_name"}`

	// ***********************  Contoh hasil enkripsi ****************************
	// strings := "U2FsdGVkX1+A6fHQtv9uULjYK3tAlJMpYGwOgfobshi1ava/vEI+nkyAcLKFfS8lv5VJtOzb1rByx9CnOM+iVw=="

	// **************** Proses Encryption **********************
	encrypted := aes256.Encrypt(data, key)

	// proses decryption
	decrypted := aes256.Decrypt(encrypted, key)

	// **************** Data original (JSON) *****************
	// fmt.Println(data)

	// **************** OUTPUT Enkripsi **********************
	fmt.Println(encrypted)

	// **************** OUTPUT Dekripsi **********************
	fmt.Println(decrypted)

}
