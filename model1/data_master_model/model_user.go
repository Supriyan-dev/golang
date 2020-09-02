package model1

import (
	"log"

	"../../db"
	"../../initialize"
	"../../models"
	"github.com/mervick/aes-everywhere/go/aes256"
)

type ModelUser_init models.DB_init

func (model1 ModelUser_init) ReturnAllDataUser() (data string, err error) {
	var all initialize.UsersEncrypt
	var allNull initialize.NullString
	rows, err := model1.DB.Query(`SELECT * FROM user`)
	if err != nil {
		log.Print(err)
	}
	rows.Next()
	errScan := rows.Scan(&all.Id_user, &all.First_name, &all.Last_name, &all.Employee_number, &all.Id_code_store, &all.Password, &all.Id_role, &allNull.Email, &allNull.Recovery_pin, &allNull.Photo_url, &allNull.Photo_name)

	if errScan != nil {
		log.Println(errScan)
	}
	tampung := all.Id_user + all.First_name + all.Last_name + all.Employee_number + all.Id_code_store + all.Password + all.Id_role + allNull.Email + allNull.Recovery_pin + allNull.Photo_url + allNull.Photo_name

	key := "P@ssw0rdL0g1n"

	encrypted := aes256.Encrypt(tampung, key)

	decrypted := aes256.Decrypt(encrypted, key)

	log.Println(encrypted)
	log.Println(decrypted)

	return tampungData, nil
}

func (model1 ModelUser_init) GetDataUser(Id_user string) (arrGet []initialize.Users, err error) {
	var get initialize.Users
	db := db.Connect()

	result, err := db.Query("SELECT id_user, first_name, last_name, employee_number, id_code_store, password, id_role, email, recovery_pin, photo_url, photo_name FROM user WHERE id_user = ?", Id_user)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {

		err := result.Scan(&get.Id_user, &get.First_name, &get.Last_name, &get.Employee_number, &get.Id_code_store, &get.Password, &get.Id_role, &get.Email, &get.Recovery_pin, &get.Photo_url, &get.Photo_name)
		if err != nil {
			panic(err.Error())
		} else {
			arrGet = append(arrGet, get)
		}

	}
	return arrGet, nil

}

func (model1 ModelUser_init) InsertDataUser(insert *initialize.Users) (arrInsert []initialize.Users, err error) {

	db := db.Connect()
	stmt, err := db.Prepare("INSERT INTO user (first_name, last_name, employee_number, id_code_store, password, id_role, email, recovery_pin, photo_url, photo_name) VALUES(?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	stmt.Exec(insert.First_name, insert.Last_name, insert.Employee_number, insert.Id_code_store, insert.Password, insert.Id_role, insert.Email, insert.Recovery_pin, insert.Photo_url, insert.Photo_name)
	if err != nil {
		panic(err.Error())
	}

	Excute := initialize.Users{
		First_name:      insert.First_name,
		Last_name:       insert.Last_name,
		Employee_number: insert.Employee_number,
		Id_code_store:   insert.Id_code_store,
		Password:        insert.Password,
		Id_role:         insert.Id_role,
		Email:           insert.Email,
		Recovery_pin:    insert.Recovery_pin,
		Photo_url:       insert.Photo_url,
		Photo_name:      insert.Photo_name,
	}

	arrInsert = append(arrInsert, Excute)

	return arrInsert, nil
}

func (model1 ModelUser_init) UpdateDataUser(update *initialize.Users) (arrUpdate []initialize.Users, err error) {
	db := db.Connect()

	stmt, err := db.Prepare("UPDATE user SET first_name = ?, last_name = ?, employee_number = ?, id_code_store = ?, password = ?, id_role = ?, email = ?, recovery_pin = ?, photo_url = ?, photo_name = ? WHERE id_user = ?")
	if err != nil {
		panic(err.Error())
	}

	stmt.Exec(update.First_name, update.Last_name, update.Employee_number, update.Id_code_store, update.Password, update.Id_role, update.Email, update.Recovery_pin, update.Photo_url, update.Photo_name, update.Id_user)
	if err != nil {
		panic(err.Error())
	}

	Excute := initialize.Users{
		Id_user:         update.Id_user,
		First_name:      update.First_name,
		Last_name:       update.Last_name,
		Employee_number: update.Employee_number,
		Id_code_store:   update.Id_code_store,
		Password:        update.Password,
		Id_role:         update.Id_role,
		Email:           update.Email,
		Recovery_pin:    update.Recovery_pin,
		Photo_url:       update.Photo_url,
		Photo_name:      update.Photo_name,
	}

	arrUpdate = append(arrUpdate, Excute)

	return arrUpdate, nil

}

func (model1 ModelUser_init) DeleteDataUser(delete *initialize.Users) (arrDelete []initialize.Users, err error) {
	db := db.Connect()
	stmt, err := db.Prepare("DELETE FROM user WHERE id_user = ?")
	if err != nil {
		panic(err.Error())
	}

	stmt.Exec(delete.Id_user)
	if err != nil {
		panic(err.Error())
	}

	Excute := initialize.Users{
		Id_user: delete.Id_user,
	}

	arrDelete = append(arrDelete, Excute)

	return arrDelete, nil
}
