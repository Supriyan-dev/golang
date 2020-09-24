package model1

import (
	"log"

	"../../db"
	"../../initialize"
	"../../models"
)

type ModelBank_init models.DB_init

func (model1 ModelBank_init) ReturnAllDatabank() (arrAll []initialize.Bank, err error) {
	var all initialize.Bank

	db := db.Connect()

	rows, err := db.Query("SELECT id_bank, bank_code, bank_name, branch_code, branch_name, special FROM bank")

	if err != nil {
		log.Print(err)
	}
	defer db.Close()

	for rows.Next() {
		if err := rows.Scan(&all.Id_bank, &all.Bank_code, &all.Bank_name, &all.Branch_code, &all.Branch_name, &all.Special); err != nil {

			log.Fatal(err.Error())

		} else {
			arrAll = append(arrAll, all)
		}
	}

	return arrAll, nil
}


func (model1 ModelBank_init) SearchDataBankModels(Keyword string) (arrJoin []initialize.Bank, err error) {
	var join initialize.Bank
	db := db.Connect()
	result, err := db.Query(`SELECT id_bank, bank_code, bank_name, branch_code, branch_name, special FROM bank" WHERE CONCAT_WS('',bank_code, bank_name, branch_code, branch_name, special) LIKE ?`, `%` + Keyword + `%`)
	if err != nil {
		panic(err.Error())
	}
	log.Println(result)
	defer db.Close()
	for result.Next() {
		if err := result.Scan(&join.Id_bank, &join.Bank_code, &join.Bank_name, &join.Branch_code, &join.Branch_name, &join.Special); err != nil {
			log.Fatal(err.Error())
		} else {
			arrJoin = append(arrJoin, join)
		}
	}

	return arrJoin, nil
}

func (model1 ModelBank_init) GetDataBank(Id_bank string) (arrGet []initialize.Bank, err error) {
	var all initialize.Bank

	db := db.Connect()

	result, err := db.Query("SELECT id_bank, bank_code, bank_name, branch_code, branch_name, special FROM bank WHERE id_bank = ?", Id_bank)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {

		err := result.Scan(&all.Id_bank, &all.Bank_code, &all.Bank_name, &all.Branch_code, &all.Branch_name, &all.Special)
		if err != nil {
			panic(err.Error())
		} else {
			arrGet = append(arrGet, all)
		}
	}

	return arrGet, nil
}

func (model1 ModelBank_init) InsertDataBank(insert *initialize.Bank) (arrInsert []initialize.Bank, err error) {
	db := db.Connect()
	stmt, err := db.Prepare("INSERT INTO bank (bank_code, bank_name, branch_code,branch_name,special) VALUES (?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	result, err := stmt.Exec(insert.Bank_code, insert.Bank_name, insert.Branch_code, insert.Branch_name, insert.Special)
	log.Println(result)

	Execute := initialize.Bank{
		Bank_code:   insert.Bank_code,
		Bank_name:   insert.Bank_name,
		Branch_code: insert.Branch_code,
		Branch_name: insert.Branch_name,
		Special:     insert.Special,
	}

	arrInsert = append(arrInsert, Execute)

	return arrInsert, nil

}

func (model1 ModelBank_init) UdpateDatabank(update *initialize.Bank) (arrUpdate []initialize.Bank, err error) {

	db := db.Connect()

	stmt, err := db.Prepare("UPDATE bank SET bank_code = ?, bank_name = ?, branch_code = ?, branch_name = ? , special = ? WHERE id_bank = ?")
	if err != nil {
		panic(err.Error())
	}

	result, err := stmt.Exec(update.Bank_code, update.Bank_name, update.Branch_code, update.Branch_name, update.Special, update.Id_bank)
	log.Println(result)

	Execute := initialize.Bank{
		Id_bank:     update.Id_bank,
		Bank_code:   update.Bank_code,
		Bank_name:   update.Bank_name,
		Branch_code: update.Branch_code,
		Branch_name: update.Branch_name,
		Special:     update.Special,
	}

	arrUpdate = append(arrUpdate, Execute)

	return arrUpdate, nil
}

func (model1 ModelBank_init) DeleteDataBank(delete *initialize.Bank) (arrDelete []initialize.Bank, err error) {
	db := db.Connect()
	stmt, err := db.Prepare("DELETE FROM bank WHERE id_bank = ?")
	if err != nil {
		panic(err.Error())
	}

	stmt.Exec(delete.Id_bank)
	if err != nil {
		panic(err.Error())
	}

	Execute := initialize.Bank{
		Id_bank: delete.Id_bank,
	}

	arrDelete = append(arrDelete, Execute)

	return arrDelete, nil
}
