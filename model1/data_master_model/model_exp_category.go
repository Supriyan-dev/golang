package model1

import (
	"log"
	"strconv"
	"../../db"
	"../../initialize"
	"../../models"
)

type ModelExp_init models.DB_init

func (model1 ModelExp_init) ReturnAllExp() (arrAll []initialize.ExpCategory, err error) {
	var all initialize.ExpCategory
	db := db.Connect()
	rows, err := db.Query("SELECT id_exp, exp_category, created_date, created_time, code_category, content, rule_code FROM exp_category")

	if err != nil {
		log.Print(err)
	}
	defer db.Close()

	for rows.Next() {
		if err := rows.Scan(&all.Id_exp, &all.Exp_category, &all.Created_date, &all.Created_time, &all.Code_category, &all.Content, &all.Rule_code); err != nil {
			log.Println(err.Error())
		} else {
			arrAll = append(arrAll, all)
		}
	}
	return arrAll, nil
}

func (model1 ModelExp_init) SearchExpCategoryModels(Keyword string, page int ,limit int) (arrSearch []initialize.ExpCategory, err error, CheckData int) {
	var Search initialize.ExpCategory
	db := db.Connect()
	querylimit := ``
	if strconv.Itoa(page) == "" && strconv.Itoa(limit) == ""{
		querylimit = ``
	}else {
		pageacheck := strconv.Itoa((page-1)*limit)
		showadata := strconv.Itoa(limit)
		querylimit = ` LIMIT `+pageacheck+`,`+showadata
	}
	db.QueryRow("SELECT count(*) FROM exp_category WHERE CONCAT_WS('',exp_category, created_date, created_time, code_category, content, rule_code ) LIKE ?", "%" + Keyword + "%").Scan(&CheckData)
	queryT := `SELECT id_exp, exp_category, created_date, created_time, code_category, content, rule_code FROM exp_category WHERE CONCAT_WS('',exp_category, created_date, created_time, code_category, content, rule_code ) LIKE ?` +querylimit

	rows, err := db.Query(queryT, "%" + Keyword + "%")

	if err != nil {
		log.Print(err)
	}

	defer db.Close()
	for rows.Next() {
		if err := rows.Scan(&Search.Id_exp, &Search.Exp_category, &Search.Created_date, &Search.Created_time, &Search.Code_category, &Search.Content, &Search.Rule_code); err != nil {
			log.Fatal(err.Error())
		} else {
			arrSearch = append(arrSearch, Search)
		}
	}

	return arrSearch, nil, CheckData
}

func (model1 ModelExp_init) GetDataExp(Id_exp string) (arrGet []initialize.ExpCategory, err error) {
	var get initialize.ExpCategory
	db := db.Connect()

	result, err := db.Query("SELECT id_exp, exp_category, created_date, created_time, code_category, content, rule_code FROM exp_category WHERE id_exp = ?", Id_exp)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {

		err := result.Scan(&get.Id_exp, &get.Exp_category, &get.Created_date, &get.Created_time, &get.Code_category, &get.Content, &get.Rule_code)
		if err != nil {
			panic(err.Error())
		} else {
			arrGet = append(arrGet, get)
		}
	}

	return arrGet, nil
}

func (model1 ModelExp_init) InsertDataExp(insert *initialize.ExpCategory) (arrInsert []initialize.ExpCategory, err error) {
	db := db.Connect()
	stmt, err := db.Prepare("INSERT INTO exp_category (exp_category, created_date, created_time, code_category, content, rule_code) VALUES(?,?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	stmt.Exec(insert.Exp_category, insert.Created_date, insert.Created_time, insert.Code_category, insert.Content, insert.Rule_code)
	if err != nil {
		panic(err.Error())
	}

	Execute := initialize.ExpCategory{
		Exp_category:  insert.Exp_category,
		Created_date:  insert.Created_date,
		Created_time:  insert.Created_time,
		Code_category: insert.Code_category,
		Content:       insert.Content,
		Rule_code:     insert.Rule_code,
	}

	arrInsert = append(arrInsert, Execute)

	return arrInsert, nil

}

func (model1 ModelExp_init) UpdateDataExp(update *initialize.ExpCategory) (arrUpdate []initialize.ExpCategory, err error) {
	db := db.Connect()

	stmt, err := db.Prepare("UPDATE exp_category SET exp_category = ?, created_date = ?, created_time = ?, code_category = ? , content = ?, rule_code = ? WHERE id_exp = ?")
	if err != nil {
		panic(err.Error())
	}
	stmt.Exec(update.Exp_category, update.Created_date, update.Created_time, update.Code_category, update.Content, update.Rule_code, update.Id_exp)
	if err != nil {
		panic(err.Error())
	}

	Execute := initialize.ExpCategory{
		Id_exp:        update.Id_exp,
		Exp_category:  update.Exp_category,
		Created_date:  update.Created_date,
		Created_time:  update.Created_time,
		Code_category: update.Code_category,
		Content:       update.Content,
		Rule_code:     update.Rule_code,
	}

	arrUpdate = append(arrUpdate, Execute)

	return arrUpdate, nil
}

func (model1 ModelExp_init) DeleteDataExp(delete *initialize.ExpCategory) (arrDelete []initialize.ExpCategory, err error) {
	db := db.Connect()
	stmt, err := db.Prepare("DELETE FROM exp_category WHERE id_exp = ?")
	if err != nil {
		panic(err.Error())
	}

	stmt.Exec(delete.Id_exp)
	if err != nil {
		panic(err.Error())
	}

	Execute := initialize.ExpCategory{
		Id_exp: delete.Id_exp,
	}

	arrDelete = append(arrDelete, Execute)

	return arrDelete, nil

}
