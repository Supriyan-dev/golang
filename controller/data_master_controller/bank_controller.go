package data_master_controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strconv"

	"../../db"
	"../../initialize"
	"github.com/gorilla/mux"
)

func ReturnAllBank(w http.ResponseWriter, r *http.Request) {
	var bank initialize.Bank
	var arrBank []initialize.Bank
	var response initialize.Response

	db := db.Connect()

	rows, err := db.Query("SELECT * FROM bank")

	if err != nil {
		log.Print(err)
	}
	defer db.Close()

	for rows.Next() {
		if err := rows.Scan(&bank.Id_bank, &bank.Bank_code, &bank.Bank_name, &bank.Branch_code, &bank.Branch_name, &bank.Special); err != nil {

			log.Fatal(err.Error())

		} else {
			arrBank = append(arrBank, bank)
		}
	}
	response.Status = 200
	response.Message = "Success"
	response.Data = arrBank

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func ReturnAllBankPagination(w http.ResponseWriter, r *http.Request) {
	var bank initialize.Bank
	var arrBank []initialize.Bank
	var response initialize.Response

	db := db.Connect()
	defer db.Close()
	code := mux.Vars(r)

	totalDataPerPage, _ := strconv.Atoi(code["perPage"])
	page, _ := strconv.Atoi(code["page"])

	var totalData int
	err := db.QueryRow("SELECT COUNT(*) FROM bank").Scan(&totalData)

	totalPage := int(math.Ceil(float64(totalData) / float64(totalDataPerPage)))

	if page > totalPage {
		page = totalPage
	}
	if page <= 0 {
		page = 1
	}

	firstIndex := (totalDataPerPage * page) - totalDataPerPage

	query := fmt.Sprintf("select id_bank,bank_code,bank_name,branch_code,branch_name,special FROM bank limit %d,%d", firstIndex, totalDataPerPage)

	rows, err := db.Query(query)
	if err != nil {
		log.Print(err)
	}
	for rows.Next() {
		if err := rows.Scan(&bank.Id_bank, &bank.Bank_code, &bank.Bank_name, &bank.Branch_code, &bank.Branch_name, &bank.Special); err != nil {

			log.Fatal(err.Error())

		} else {
			arrBank = append(arrBank, bank)
		}
	}
	response.Status = 200
	response.Message = "Success"
	response.Data = arrBank
	response.TotalPage = totalPage
	response.CurrentPage = page

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetBank(w http.ResponseWriter, r *http.Request) {
	var bank initialize.Bank
	var arrBank []initialize.Bank
	var response initialize.Response

	db := db.Connect()
	code := mux.Vars(r)

	result, err := db.Query("SELECT id_bank, bank_code, bank_name, branch_code, branch_name, special FROM bank WHERE id_bank = ?", code["id_bank"])
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {

		err := result.Scan(&bank.Id_bank, &bank.Bank_code, &bank.Bank_name, &bank.Branch_code, &bank.Branch_name, &bank.Special)
		if err != nil {
			panic(err.Error())
		} else {
			arrBank = append(arrBank, bank)
		}
	}
	response.Status = 200
	response.Message = "Success"
	response.Data = arrBank

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func CreateBank(w http.ResponseWriter, r *http.Request) {
	var err error
	var response initialize.Response

	db := db.Connect()
	stmt, err := db.Prepare("INSERT INTO bank (bank_code, bank_name, branch_code,branch_name,special) VALUES(?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	BankCode := keyVal["bank_code"]
	BankName := keyVal["bank_name"]
	BranchCode := keyVal["branch_code"]
	BranchName := keyVal["branch_name"]
	Sprecial := keyVal["special"]

	result, err := stmt.Exec(BankCode, BankName, BranchCode, BranchName, Sprecial)
	if err != nil {
		panic(err.Error())
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}

	response.Status = 200
	response.Message = "Success"
	response.Data = map[string]int64{
		"Data baru telah dibuat": rowsAffected,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func UpdateBank(w http.ResponseWriter, r *http.Request) {
	var response initialize.Response

	db := db.Connect()

	stmt, err := db.Prepare("UPDATE bank SET bank_code = ?, bank_name = ?, branch_code = ?, branch_name = ? , special = ? WHERE id_bank = ?")
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	idBank := keyVal["id_bank"]
	NewBankCode := keyVal["bank_code"]
	NewBankName := keyVal["bank_name"]
	NewsBranchCode := keyVal["branch_code"]
	NewsBranchName := keyVal["branch_name"]
	NewSpecial := keyVal["special"]

	id, err := strconv.Atoi(idBank)

	result, err := stmt.Exec(NewBankCode, NewBankName, NewsBranchCode, NewsBranchName, NewSpecial, id)
	if err != nil {
		panic(err.Error())
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}

	response.Status = 200
	response.Message = "Success"
	response.Data = map[string]int64{
		"Data Yang Behasil Di Update": rowsAffected,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func DeleteBank(w http.ResponseWriter, r *http.Request) {

	db := db.Connect()
	params := mux.Vars(r)
	stmt, err := db.Prepare("DELETE FROM bank WHERE id_bank = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(params["id_bank"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Data Sudah Terhapus Dengan ID = ")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(params["id_bank"])

}
