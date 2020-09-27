package model1

import (
	"log"

	"../../db"
	"../../initialize"
	"../../models"
)

type ModelFull_init models.DB_init

func (model1 ModelFull_init) ReturnAllFulltime() (arrGet []initialize.FullTimeSalary, err error) {
	var all initialize.FullTimeSalary
	db := db.Connect()

	rows, err := db.Query("SELECT id_full_time_salary, id_code_store, salary, fish_section_salary FROM full_time_salary")

	if err != nil {
		log.Print(err)
	}
	defer db.Close()

	for rows.Next() {
		if err := rows.Scan(&all.Id_full_time_salary, &all.Id_code_store, &all.Salary, &all.Fish_section_salary); err != nil {

			log.Println(err.Error())

		} else {
			arrGet = append(arrGet, all)
		}
	}

	return arrGet, nil
}

func (model1 ModelFull_init) SearchFullTimeSalaryModels(Keyword string) (arrJoin []initialize.FullTimeSalary, err error) {
	var all initialize.FullTimeSalary
	db := db.Connect()
	result, err := db.Query(`SELECT id_full_time_salary, id_code_store, salary, fish_section_salary WHERE CONCAT_WS('',id_code_store, salary, fish_section_salary) LIKE ?`, `%` + Keyword + `%`)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(result)
	defer db.Close()
	for result.Next() {
		if err := result.Scan(&all.Id_full_time_salary, &all.Id_code_store, &all.Salary, &all.Fish_section_salary); err != nil {
			log.Println(err.Error())
		} else {
			arrJoin = append(arrJoin, all)
		}
	}

	return arrJoin, nil
}

func (model1 ModelFull_init) GetDataFullTime(Id_full_time_salary string) (arrGet []initialize.FullTimeSalary, err error) {
	var get initialize.FullTimeSalary
	db := db.Connect()

	result, err := db.Query("SELECT id_full_time_salary, id_code_store, salary, fish_section_salary FROM full_time_salary WHERE id_full_time_salary = ?", Id_full_time_salary)
	if err != nil {
		log.Println(err.Error())
	}
	defer result.Close()
	for result.Next() {

		err := result.Scan(&get.Id_full_time_salary, &get.Id_code_store, &get.Salary, &get.Fish_section_salary)
		if err != nil {
			log.Println(err.Error())
		} else {
			arrGet = append(arrGet, get)
		}
	}

	return arrGet, nil
}

func (model1 ModelFull_init) InsertDataFullTime(insert *initialize.FullTimeSalary) (arrInsert []initialize.FullTimeSalary, err error) {
	db := db.Connect()
	stmt, err := db.Prepare("INSERT INTO full_time_salary (id_code_store, salary, fish_section_salary) VALUES(?,?,?)")
	if err != nil {
		log.Println(err.Error())
	}
	defer db.Close()

	stmt.Exec(insert.Id_code_store, insert.Salary, insert.Fish_section_salary)
	if err != nil {
		log.Println(err.Error())
	}

	Excute := initialize.FullTimeSalary{
		Id_code_store:       insert.Id_code_store,
		Salary:              insert.Salary,
		Fish_section_salary: insert.Fish_section_salary,
	}

	arrInsert = append(arrInsert, Excute)

	return arrInsert, nil

}

func (model1 ModelFull_init) UpdateDataFullTime(update *initialize.FullTimeSalary) (arrUpdate []initialize.FullTimeSalary, err error) {
	db := db.Connect()

	stmt, err := db.Prepare("UPDATE full_time_salary SET id_code_store = ?, salary = ?, fish_section_salary = ? WHERE id_full_time_salary = ?")
	if err != nil {
		log.Println(err.Error())
	}

	stmt.Exec(update.Id_code_store, update.Salary, update.Fish_section_salary, update.Id_full_time_salary)
	if err != nil {
		log.Println(err.Error())
	}

	Excute := initialize.FullTimeSalary{
		Id_full_time_salary: update.Id_full_time_salary,
		Id_code_store:       update.Id_code_store,
		Salary:              update.Salary,
		Fish_section_salary: update.Fish_section_salary,
	}

	arrUpdate = append(arrUpdate, Excute)

	return arrUpdate, nil

}

func (model1 ModelFull_init) DeleteDataFullTime(delete *initialize.FullTimeSalary) (arrDelete []initialize.FullTimeSalary, err error) {
	db := db.Connect()
	stmt, err := db.Prepare("DELETE FROM full_time_salary WHERE id_full_time_salary = ?")
	if err != nil {
		log.Println(err.Error())
	}

	stmt.Exec(delete.Id_full_time_salary)
	if err != nil {
		log.Println(err.Error())
	}

	Excute := initialize.FullTimeSalary{
		Id_full_time_salary: delete.Id_full_time_salary,
	}

	arrDelete = append(arrDelete, Excute)

	return arrDelete, nil
}
