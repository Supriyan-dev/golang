// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"

// 	"../db"

// 	"github.com/mervick/aes-everywhere/go/aes256"
// )

// // type Person struct {
// // 	First string `json:"fname"`
// // 	Last  string
// // }

// type Users struct {
// 	Id_user         int            `json:"id_user"`
// 	First_name      string         `json:"first_name"`
// 	Last_name       string         `json:"last_name"`
// 	Employee_number string         `json:"employee_number"`
// 	Id_code_store   int            `json:"id_code_store"`
// 	Password        string         `json:"password"`
// 	Id_role         int            `json:"id_role"`
// 	Email           sql.NullString `json:"email"`
// 	Recovery_pin    sql.NullString `json:"recovery_pin"`
// 	Photo_url       sql.NullString `json:"photo_url"`
// 	Photo_name      sql.NullString `json:"photo_name"`
// }

// func main() {
// 	// j := `[{"fname":"jeffri","last":"yanto"},{"fname":"Miss","last":"Monesy"}]`
// 	// fmt.Println("json:", j)

// 	// xp := []Person{}

// 	// err := json.Unmarshal([]byte(j), &xp)

// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// }

// 	// fmt.Printf("go data: %+v\n", xp)
// 	// for i, v := range xp {
// 	// 	fmt.Println(i, v)
// 	// 	fmt.Println("\t", v.First)
// 	// }

// 	var all Users
// 	arrAll := []Users{}

// 	db := db.Connect()

// 	rows, err := db.Query("SELECT * FROM user")
// 	if err != nil {
// 		log.Print(err)
// 	}
// 	defer db.Close()

// 	for rows.Next() {
// 		if err := rows.Scan(&all.Id_user, &all.First_name, &all.Last_name, &all.Employee_number, &all.Id_code_store, &all.Password, &all.Id_role, &all.Email, &all.Recovery_pin, &all.Photo_url, &all.Photo_name); err != nil {
// 			log.Fatal(err.Error())

// 		} else {
// 			arrAll = append(arrAll, all)
// 		}
// 	}

// 	var Id_user string
// 	var First_name string
// 	var Last_name string
// 	var Employee_number string
// 	var Id_code_store string
// 	var Password string
// 	var Email string
// 	var Recovery_pin string
// 	var Photo_url string
// 	var Photo_name string
// 	var tampungData string

// 	tampungData := rows.Scan(all.Id_user, all.First_name, all.Last_name, all.Employee_number, all.Id_code_store, all.Password, all.Email, all.Recovery_pin, all.Photo_url, all.Photo_name)
// 	// 	log.Println(err)
// 	// } else {
// 	key := "P@ssw0rdL0g1n"

// 	// tampungData := Id_user + First_name + Last_name + Employee_number + Id_code_store + Password + Email + Recovery_pin + Photo_url + Photo_name

// 	// json.Unmarshal([]byte(tampungData), &arrAll)

// 	encrypted := aes256.Encrypt(tampungData, key)

// 	decrypted := aes256.Decrypt(encrypted, key)

// 	// **************** Data original (JSON) *****************
// 	// fmt.Println(data)

// 	// **************** OUTPUT Enkripsi **********************
// 	fmt.Println(encrypted)

// 	// **************** OUTPUT Dekripsi **********************
// 	fmt.Println(decrypted)
// }
