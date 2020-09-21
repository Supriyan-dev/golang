package model1

import (
	"log"

	"../../db"
	"../../initialize/driver_management"
	"../../models"
)

type ModelUnit_init models.DB_init

func (model1 ModelUnit_init) ReturnAllDataUnitInformation() (arrAll []initialize.UnitInformation, err error) {
	var all initialize.UnitInformation
	db := db.Connect()

	rows, err := db.Query("SELECT id_unit, unit_code, unit_name FROM unit_information")

	if err != nil {
		log.Print(err)
	}
	defer db.Close()

	for rows.Next() {
		if err := rows.Scan(&all.Id_unit, &all.Unit_code, &all.Unit_name); err != nil {

			log.Fatal(err.Error())

		} else {
			arrAll = append(arrAll, all)
		}
	}
	return arrAll, nil
}

func (model1 ModelUnit_init) GetDataUnitInformation(Id_unit string) (arrGet []initialize.UnitInformation, err error) {
	var get initialize.UnitInformation

	db := db.Connect()
	result, err := db.Query("SELECT id_unit, unit_code, unit_name FROM unit_information WHERE id_unit = ?", Id_unit)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {

		err := result.Scan(&get.Id_unit, &get.Unit_code, &get.Unit_name)
		if err != nil {
			panic(err.Error())
		} else {
			arrGet = append(arrGet, get)
		}
	}

	return arrGet, nil
}

func (model1 ModelUnit_init) InsertDataUnitInformation(insert *initialize.UnitInformation) (arrInsert []initialize.UnitInformation, err error) {

	db := db.Connect()
	stmt, err := db.Prepare("INSERT INTO unit_information (unit_code,unit_name) VALUES (?,?)")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	stmt.Exec(insert.Unit_code, insert.Unit_name)
	if err != nil {
		panic(err.Error())
	}

	Excute := initialize.UnitInformation{
		Unit_code: insert.Unit_code,
		Unit_name: insert.Unit_name,
	}

	arrInsert = append(arrInsert, Excute)

	return arrInsert, nil
}

func (model1 ModelUnit_init) UpdateDataUnitInformation(update *initialize.UnitInformation) (arrUpdate []initialize.UnitInformation, err error) {
	db := db.Connect()

	stmt, err := db.Prepare("UPDATE unit_information SET unit_code = ?, unit_name = ? WHERE id_unit = ?")
	if err != nil {
		panic(err.Error())
	}
	stmt.Exec(update.Unit_code, update.Unit_name, update.Id_unit)
	if err != nil {
		panic(err.Error())
	}

	Excute := initialize.UnitInformation{
		Id_unit:   update.Id_unit,
		Unit_code: update.Unit_code,
		Unit_name: update.Unit_name,
	}

	arrUpdate = append(arrUpdate, Excute)

	return arrUpdate, nil
}

func (Model1 ModelUnit_init) DeleteDataUnitInformation(delete *initialize.UnitInformation) (arrDelete []initialize.UnitInformation, err error) {
	db := db.Connect()
	stmt, err := db.Prepare("DELETE FROM unit_information WHERE id_unit = ?")
	if err != nil {
		panic(err.Error())
	}
	stmt.Exec(delete.Id_unit)
	if err != nil {
		panic(err.Error())
	}

	Excute := initialize.UnitInformation{
		Id_unit: delete.Id_unit,
	}

	arrDelete = append(arrDelete, Excute)

	return arrDelete, nil
}
