package data_master_controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jeffri/golang-test/GO_DX_SERVICES/db"

	"github.com/jeffri/golang-test/GO_DX_SERVICES/initialize"
)

func ReturnAllExpCategory(w http.ResponseWriter, r *http.Request) {
	var exp initialize.ExpCategory
	var arrExpCategory []initialize.ExpCategory
	var response initialize.Response

	db := db.Connect()

	rows, err := db.Query("SELECT * FROM exp_category")

	if err != nil {
		log.Print(err)
	}
	defer db.Close()

	for rows.Next() {
		if err := rows.Scan(&exp.Id_exp, &exp.Exp_category, &exp.Created_date, &exp.Created_time, &exp.Code_category, &exp.Content, &exp.Rule_code); err != nil {

			log.Fatal(err.Error())

		} else {
			arrExpCategory = append(arrExpCategory, exp)
		}
	}
	response.Status = 200
	response.Message = "Success"
	response.Data = arrExpCategory

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func ReturnAllExpCategoryPagination(w http.ResponseWriter, r *http.Request) {
	var exp initialize.ExpCategory
	var arrExpCategory []initialize.ExpCategory
	var response initialize.Response

	db := db.Connect()
	defer db.Close()
	code := mux.Vars(r)

	totalDataPerPage, _ := strconv.Atoi(code["perPage"])
	page, _ := strconv.Atoi(code["page"])

	var totalData int
	err := db.QueryRow("SELECT COUNT(*) FROM exp_category").Scan(&totalData)

	totalPage := int(math.Ceil(float64(totalData) / float64(totalDataPerPage)))

	if page > totalPage {
		page = totalPage
	}
	if page <= 0 {
		page = 1
	}

	firstIndex := (totalDataPerPage * page) - totalDataPerPage

	query := fmt.Sprintf("select id_exp, exp_category, created_date, created_time, code_category, content, rule_code from exp_category limit %d,%d", firstIndex, totalDataPerPage)

	rows, err := db.Query(query)
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&exp.Id_exp, &exp.Exp_category, &exp.Created_date, &exp.Created_time, &exp.Code_category, &exp.Content, &exp.Rule_code); err != nil {

			log.Fatal(err.Error())

		} else {
			arrExpCategory = append(arrExpCategory, exp)
		}
	}
	response.Status = 200
	response.Message = "Success"
	response.Data = arrExpCategory
	response.TotalPage = totalPage
	response.CurrentPage = page

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func GetExpCategory(w http.ResponseWriter, r *http.Request) {
	var expCategory initialize.ExpCategory
	var arrExpCategory []initialize.ExpCategory
	var response initialize.Response

	db := db.Connect()
	code := mux.Vars(r)
	fmt.Fprintf(w, "Category: %v\n", code["id_exp"])

	result, err := db.Query("SELECT id_exp, exp_category, created_date, created_time, code_category, content, rule_code FROM exp_category WHERE id_exp = ?", code["id_exp"])
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {

		err := result.Scan(&expCategory.Id_exp, &expCategory.Exp_category, &expCategory.Created_date, &expCategory.Created_time, &expCategory.Code_category, &expCategory.Content, &expCategory.Rule_code)
		if err != nil {
			panic(err.Error())
		} else {
			arrExpCategory = append(arrExpCategory, expCategory)
		}
	}
	response.Status = 200
	response.Message = "Success"
	response.Data = arrExpCategory

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func CreateExpCategory(w http.ResponseWriter, r *http.Request) {
	var err error
	var response initialize.Response

	db := db.Connect()
	stmt, err := db.Prepare("INSERT INTO exp_category (exp_category, created_date, created_time, code_category, content, rule_code) VALUES(?,?,?,?,?,?)")
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
	ExpCategory := keyVal["exp_category"]
	CreatedDate := keyVal["created_date"]
	CreatedTime := keyVal["created_time"]
	CodeCategory := keyVal["code_category"]
	Content := keyVal["content"]
	RuleCode := keyVal["rule_code"]

	result, err := stmt.Exec(ExpCategory, CreatedDate, CreatedTime, CodeCategory, Content, RuleCode)
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

func UpdateExpCategory(w http.ResponseWriter, r *http.Request) {
	var response initialize.Response

	db := db.Connect()

	stmt, err := db.Prepare("UPDATE exp_category SET exp_category = ?, created_date = ?, created_time = ?, code_category = ? , content = ?, rule_code = ? WHERE id_exp = ?")
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	idExp := keyVal["id_exp"]
	ExpCategory := keyVal["exp_category"]
	CreatedDate := keyVal["created_date"]
	CreatedTime := keyVal["created_time"]
	CodeCategory := keyVal["code_category"]
	Content := keyVal["content"]
	RuleCode := keyVal["rule_code"]

	id, err := strconv.Atoi(idExp)

	result, err := stmt.Exec(ExpCategory, CreatedDate, CreatedTime, CodeCategory, Content, RuleCode, id)
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

func DeleteExpCategory(w http.ResponseWriter, r *http.Request) {

	db := db.Connect()
	params := mux.Vars(r)
	stmt, err := db.Prepare("DELETE FROM exp_category WHERE id_exp = ?")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(params["id_exp"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Data Sudah Terhapus Dengan ID = ")

	w.Header().Set("Content-Type", "Aplication/json")
	json.NewEncoder(w).Encode(params["id_exp"])

}
