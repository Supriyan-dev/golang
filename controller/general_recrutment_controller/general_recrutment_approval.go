package general_recrutment_controller

import (
	"log"
	"net/http"

	"../../db"
)

func DataGeneralRecrutment(w http.ResponseWriter, r *http.Request) {
	// var test initialize.User
	// var test1 []initialize.User
	// var _response initialize.Response
	// ctx := context.Background()
	db := db.Connect()
	var err error
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	// username := r.FormValue("username")
	// password := r.FormValue("password")

	{
		stmt, err := tx.Prepare(`SELECT employee_code FROM basic_information WHERE id_basic_information = ?
		INSERT INTO basic_information (employee_code, first_name, last_name, gender, birthdate, add_postal_code, id_prefecture, addres, addres_kana, addres_detail, addres_detail_kana, add_phone_number, marital_status, dormitory_status)
		VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?)
		INNER JOIN general_information AS employee_code ON basic_information.id_basic_information = general_information.id_basic_information`)
		if err != nil {
			tx.Rollback()
			log.Fatal(err)
		}
		defer stmt.Close()
		if _, err := stmt.Exec(employee_code, first_name, last_name, gender, birthdate, add_postal_code, id_prefecture, addres, addres_kana, addres_detail, addres_detail_kana, add_phone_number, marital_status, dormitory_status); err != nil {
			tx.Rollback() // log.Fatal(err) an error too, we may want to wrap them
			log.Fatal(err)
		}
	}

	{
		stmt, err := tx.Prepare(`SELECT * FROM general_information WHERE id_general_information = ?
		SELECT id_basic_information, id_store_code, id_department, id_code_store, id_unit, join_date, id_bank, account_type, account_number, account_name, first_smester_in_other_company, distance_trip, resume_document, written_oath_document, employee_agreement_document, certificate_of_residence_card_document, application_form_of_commuting_method_document, compliance_agreement_document, with_holding_slip_document, dependent_deduction_form_document, pension_book_document, health_check_report_document, office_code FORM general_information INNER JOIN employment_type ON general_information.id_basic_information = employment_type.id_basic_information
		INSERT INTO employment_type IF(employment_status = 34) SELECT employment_status FROM employment_type INNER JOIN full_time_employee ON employment_type.id_employment_type = full_time_employee.id_employment_type
		INSERT INTO employment_type IF(employment_status = 41 OR 43) SELECT employment_status FROM employment_type INNER JOIN part_time_above_18_employee ON employment_type.id_employment_type = part_time_above_18_employee.id_employment_type
		INSERT INTO employment_type IF(employment_status = 42) SELECT employment_status FROM employment_type INNER JOIN part_time_under_18_employee ON employment_type.id_employment_type = part_time_under_18_employee.id_employment_type`)
		if err != nil {
			tx.Rollback()
			log.Fatal(err)
		}
		defer stmt.Close()
		if _, err := stmt.Exec(id_basic_information, id_store_code, id_department, id_code_store, id_unit, join_date, id_bank, account_type, account_number, account_name, first_smester_in_other_company, distance_trip, resume_document, written_oath_document, employee_agreement_document, certificate_of_residence_card_document, application_form_of_commuting_method_document, compliance_agreement_document, with_holding_slip_document, dependent_deduction_form_document, pension_book_document, health_check_report_document, office_code); err != nil {
			tx.Rollback() // log.Fatal(err) an error too, we may want to wrap them
			log.Fatal(err)
		}
	}

	// {
	// 	stmt, err := tx.Prepare(` SELECT employee_code, first_name, last_name, gender, birthdate, add_postal_code, id_prefecture, addres, addres_kana, addres_detail, addres_detail_kana, add_phone_number, marital_status, dormitory_status FORM basic_information WHERE employee_number = employee_number
	// 	INSERT INTO basic_information (employee_code, first_name, last_name, gender, birthdate, add_postal_code, id_prefecture, addres, addres_kana, addres_detail, addres_detail_kana, add_phone_number, marital_status, dormitory_status)
	// 	INNER JOIN general_information AS employee_code ON basic_information.id_basic_information = general_information.id_basic_information`)
	// 	if err != nil {
	// 		tx.Rollback()
	// 		log.Fatal(err)
	// 	}
	// 	defer stmt.Close()
	// 	if _, err := stmt.Exec(employee_code, first_name, last_name, gender, birthdate, add_postal_code, id_prefecture, addres, addres_kana, addres_detail, addres_detail_kana, add_phone_number, marital_status, dormitory_status); err != nil {
	// 		tx.Rollback() // log.Fatal(err) an error too, we may want to wrap them
	// 		log.Fatal(err)
	// 	}
	// }

	// {
	// 	stmt, err := tx.Prepare(`INSERT INTO general_information (id_basic_information, id_store_code, id_department, id_code_store, id_unit, join_date, id_bank, account_type, account_number, account_name, first_smester_in_other_company, distance_trip, resume_document, written_oath_document, employee_agreement_document, certificate_of_residence_card_document, application_form_of_commuting_method_document, compliance_agreement_document, with_holding_slip_document, dependent_deduction_form_document, pension_book_document, health_check_report_document, office_code)
	// 	SELECT id_basic_information, id_store_code, id_department, id_code_store, id_unit, join_date, id_bank, account_type, account_number, account_name, first_smester_in_other_company, distance_trip, resume_document, written_oath_document, employee_agreement_document, certificate_of_residence_card_document, application_form_of_commuting_method_document, compliance_agreement_document, with_holding_slip_document, dependent_deduction_form_document, pension_book_document, health_check_report_document, office_code FORM general_information INNER JOIN employment_type ON general_information.id_basic_information = employment_type.id_basic_information
	// 	INSERT INTO employment_type IF(employment_status = 34) SELECT employment_status FROM employment_type INNER JOIN full_time_employee ON employment_type.id_employment_type = full_time_employee.id_employment_type
	// 	INSERT INTO employment_type IF(employment_status = 41 OR 43) SELECT employment_status FROM employment_type INNER JOIN part_time_above_18_employee ON employment_type.id_employment_type = part_time_above_18_employee.id_employment_type
	// 	INSERT INTO employment_type IF(employment_status = 42) SELECT employment_status FROM employment_type INNER JOIN part_time_under_18_employee ON employment_type.id_employment_type = part_time_under_18_employee.id_employment_type`)
	// 	if err != nil {
	// 		tx.Rollback()
	// 		log.Fatal(err)
	// 	}
	// 	defer stmt.Close()
	// 	if _, err := stmt.Exec(id_basic_information, id_store_code, id_department, id_code_store, id_unit, join_date, id_bank, account_type, account_number, account_name, first_smester_in_other_company, distance_trip, resume_document, written_oath_document, employee_agreement_document, certificate_of_residence_card_document, application_form_of_commuting_method_document, compliance_agreement_document, with_holding_slip_document, dependent_deduction_form_document, pension_book_document, health_check_report_document, office_code); err != nil {
	// 		tx.Rollback() // log.Fatal(err) an error too, we may want to wrap them
	// 		log.Fatal(err)
	// 	}
	// }

	commitTx := tx.Commit()

	if commitTx != nil {
		log.Fatal(commitTx)
	}

	// _, errB := tx.ExecContext(ctx, "INSERT INTO users (username, password) VALUES (?, ?)", username, password)
	// // if err != nil {
	// // 	tx.Rollback()
	// // 	log.Fatal(err)

	// // 	return
	// // }

	// _, errC := tx.ExecContext(ctx, "INSERT INTO las_user (username1, password11) VALUES (?, ?)", username, password)
	// if errC != nil && errB != nil {
	// 	tx.Rollback()
	// 	log.Fatal(err)
	// 	return
	// }

	// if err != nil {
	// 	log.Println("gagal masuk data")
	// } else {
	// 	log.Println("berhasil masuk data")
	// }

	// // Finally, if no errors are recieved from the queries, commit the transaction
	// // this applies the above changes to our database
	// err = tx.Commit()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// _, err = tx.ExecContext(ctx, "INSERT INTO users (username, password) VALUES ('jeffri', 'asdf'), ('admin', '12345')")
	// if err != nil {
	// 	tx.Rollback()
	// 	log.Println(err)

	// 	return
	// }
	// // for rows.Next() {
	// // 	rows.Scan(&test.Username, &test.Password)

	// // }

	// // Run a query to get a count of all cats
	// row := tx.QueryRow("SELECT count(*) FROM users WHERE username='jeffri'")
	// var catCount int
	// // Store the count in the `catCount` variable
	// err = row.Scan(&catCount)
	// if err != nil {
	// 	tx.Rollback()
	// 	log.Println(err)
	// 	return
	// }

	// // Now update the food table, increasing the quantity of cat food by 10x the number of cats
	// _, err = tx.ExecContext(ctx, "UPDATE las_user SET last_user=last_user+$1 WHERE last_password=''", 10*catCount)
	// if err != nil {
	// 	tx.Rollback()
	// 	log.Println(err)

	// 	return
	// }

	// // Commit the change if all queries ran successfully
	// err = tx.Commit()
	// if err != nil {
	// 	log.Println(err)

	// 	log.Fatal(err)
	// }

	// if r.Method == "POST" {
	// 	if ExcuteData == nil {
	// 		_response.Status = http.StatusBadRequest
	// 		_response.Message = "Sorry Your Input Missing Body Bad Request"
	// 		_response.Data = "Null"
	// 		response.ResponseJson(w, _response.Status, _response)
	// 	} else {
	// _response.Status = http.StatusOK
	// _response.Message = "Success"
	// _response.Data = tx.Commit()
	// response.ResponseJson(w, _response.Status, _response)
	// 	}
	// } else {
	// 	_response.Status = http.StatusMethodNotAllowed
	// 	_response.Message = "Sorry Your Method Missing Not Allowed"
	// 	_response.Data = "Null"
	// 	response.ResponseJson(w, _response.Status, _response)
	// }

}
