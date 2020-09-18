package model1

import (
	"log"

	"../../db"
	"../../initialize"
	"../../models"
)

type ModelAbove_init models.DB_init

func (model1 ModelAbove_init) ReturnAllDataAbove() (arrAll []initialize.PartTimeAbove18Salary, err error) {
	var all initialize.PartTimeAbove18Salary
	db := db.Connect()

	rows, err := db.Query("SELECT id_part_time_above_18_salary, id_code_store, Day_salary, night_salary, morning_salary, peek_time_salary FROM part_time_above_18_salary")

	if err != nil {
		log.Print(err)
	}
	defer db.Close()

	for rows.Next() {
		if err := rows.Scan(&all.Id_part_time_above_18_salary, &all.Id_code_store, &all.Day_salary, &all.Night_salary, &all.Morning_salary, &all.Peek_time_salary); err != nil {

			log.Fatal(err.Error())

		} else {
			arrAll = append(arrAll, all)
		}
	}

	return arrAll, nil
}

func (model1 ModelAbove_init) GetDataAbove(Id_part_time_above_18_salary string) (arrGet []initialize.PartTimeAbove18Salary, err error) {
	var get initialize.PartTimeAbove18Salary
	db := db.Connect()

	result, err := db.Query("SELECT id_part_time_above_18_salary, id_code_store, day_salary,night_salary, morning_salary, peek_time_salary FROM part_time_above_18_salary WHERE id_part_time_above_18_salary = ?", Id_part_time_above_18_salary)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {

		err := result.Scan(&get.Id_part_time_above_18_salary, &get.Id_code_store, &get.Day_salary, &get.Night_salary, &get.Morning_salary, &get.Peek_time_salary)
		if err != nil {
			panic(err.Error())
		} else {
			arrGet = append(arrGet, get)
		}
	}

	return arrGet, nil
}

func (model1 ModelAbove_init) InsertDataPartTimeAbove(insert *initialize.PartTimeAbove18Salary) (arrInsert []initialize.PartTimeAbove18Salary, err error) {
	db := db.Connect()
	stmt, err := db.Prepare("INSERT INTO part_time_above_18_salary (id_code_store, day_salary, night_salary, morning_salary, peek_time_salary) VALUES(?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	stmt.Exec(insert.Id_code_store, insert.Day_salary, insert.Night_salary, insert.Morning_salary, insert.Peek_time_salary)
	if err != nil {
		panic(err.Error())
	}

	Excute := initialize.PartTimeAbove18Salary{
		Id_code_store:    insert.Id_code_store,
		Day_salary:       insert.Day_salary,
		Night_salary:     insert.Night_salary,
		Morning_salary:   insert.Morning_salary,
		Peek_time_salary: insert.Peek_time_salary,
	}

	arrInsert = append(arrInsert, Excute)

	return arrInsert, nil

}

func (model1 ModelAbove_init) UpdateDataPartTimeAbove(update *initialize.PartTimeAbove18Salary) (arrUpdate []initialize.PartTimeAbove18Salary, err error) {
	db := db.Connect()

	stmt, err := db.Prepare("UPDATE part_time_above_18_salary SET id_code_store = ?, day_salary = ?, night_salary = ?, morning_salary = ?, peek_time_salary = ? WHERE id_part_time_above_18_salary = ?")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	stmt.Exec(update.Id_code_store, update.Day_salary, update.Night_salary, update.Morning_salary, update.Peek_time_salary, update.Id_part_time_above_18_salary)
	if err != nil {
		panic(err.Error())
	}

	Excute := initialize.PartTimeAbove18Salary{
		Id_part_time_above_18_salary: update.Id_part_time_above_18_salary,
		Id_code_store:                update.Id_code_store,
		Day_salary:                   update.Day_salary,
		Night_salary:                 update.Night_salary,
		Morning_salary:               update.Morning_salary,
		Peek_time_salary:             update.Peek_time_salary,
	}

	arrUpdate = append(arrUpdate, Excute)

	return arrUpdate, nil

}

func (model1 ModelAbove_init) DeleteDataPartTimeAbove(delete *initialize.PartTimeAbove18Salary) (arrDetete []initialize.PartTimeAbove18Salary, err error) {
	db := db.Connect()
	stmt, err := db.Prepare("DELETE FROM part_time_above_18_salary WHERE id_part_time_above_18_salary = ?")
	if err != nil {
		panic(err.Error())
	}

	stmt.Exec(delete.Id_part_time_above_18_salary)
	if err != nil {
		panic(err.Error())
	}

	Excute := initialize.PartTimeAbove18Salary{
		Id_part_time_above_18_salary: delete.Id_part_time_above_18_salary,
	}

	arrDetete = append(arrDetete, Excute)

	return arrDetete, nil
}
