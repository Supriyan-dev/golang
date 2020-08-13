package data_master_controller

import (
	"encoding/json"
	"fmt"
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

	query := fmt.Sprintf("select id_exp,exp_category,created_date,created_time,code_category,content,rule_code from exp_category limit %d,%d", firstIndex, totalDataPerPage)

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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
