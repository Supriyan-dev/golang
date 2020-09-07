package general_recrutment_model

import (
	"context"
	"log"

	"../../db"
	initialize "../../initialize/general_recrutment"
	"../../models"
)

type ModelGeneral_init models.DB_init

func (model1 ModelGeneral_init) InsertDataGeneralRecrutment() (allArr []initialize.GeneralRecrutmentStatusApproval, err error) {
	// var all initialize.GeneralRecrutmentStatusApproval
	ctx := context.Background()

	db := db.Connect()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	_, err = tx.ExecContext(ctx, "INSERT INTO users (username, password) VALUES ('jeffri', 'asdf'), ('admin', '12345')")
	if err != nil {
		tx.Rollback()
		return
	}

	// Run a query to get a count of all cats
	row := tx.QueryRow("SELECT count(*) FROM users WHERE username='jeffri'")
	var catCount int
	// Store the count in the `catCount` variable
	err = row.Scan(&catCount)
	if err != nil {
		tx.Rollback()
		return
	}

	// Now update the food table, increasing the quantity of cat food by 10x the number of cats
	_, err = tx.ExecContext(ctx, "UPDATE last_user SET last_name=last_name+$1 WHERE last_password='admin'", 10*catCount)
	if err != nil {
		tx.Rollback()
		return
	}

	// Commit the change if all queries ran successfully
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	// {
	// result, err := tx.Exec("INSERT INTO basic_information (employee_code, first_name, last_name, gender, birthdate, add_postal_code, id_prefecture, addres, addres_kana, addres_detail, addres_detail_kana, add_phone_number, marital_status, dormitory_status) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	// if err != nil {
	// 	tx.Rollback()
	// 	return nil, err
	// 	log.Println(err)
	// }

	// if _, err := result.Exec(all.Employee_code, first_name, last_name, gender, birthdate, add_postal_code, id_prefecture, addres, addres_kana, addres_detail, addres_detail_kana, add_postal_code, marital_status, dormitory_status); err != nil {
	// 	tx.Rollback()
	// 	return nil, err
	// 	log.Println(err)
	// }
	// }
	// {
	// result1, err := tx.Exec("INSERT INTO general_information (id_basic_information, id_store_code, id_department, id_code_store, id_unit, join_date, id_bank, account_type, account_number, account_name, first_smester_in_other_company, distance_trip, resume_document, written_oath_document, employee_agreement_document, certificate_of_residence_card_document, application_form_of_commuting_method_document, compliance_agreement_document, with_holding_slip_document, dependent_deduction_form_document, pension_book_document, health_check_report_document, office_code) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	// if err != nil {
	// 	tx.Rollback()
	// 	return nil, err
	// 	log.Println(err)
	// }

	// if _, err := result.Exce(general_informaeExecon, id_store_code, id_department, id_code_store, id_unit, join_date, id_bank, account_type, account_number, account_name, first_smester_in_other_company, distance_trip, resume_document, written_oath_document, employee_agreement_document, certificate_of_residence_card_document, application_form_of_commuting_method_document, compliance_agreement_document, with_holding_slip_document, dependent_deduction_form_document, pension_book_document, health_check_report_document, office_code); err != nil {
	// 	tx.Rollback()
	// 	return nil, err
	// 	log.Println(err)
	// }

	// }
	// {
	// result2, err := tx.Exec("INSERT INTO general_recruitment_status_approval (id_basic_information, status, reason, date_time, date_time_approve, date_check, date_time_date_check, message_for_edit, by_employee_code, flag, real_data) VALUES (?,?,?,?,?,?,?,?,?,?,?,?)")
	// if err != nil {
	// 	tx.Rollback()
	// 	return nil, err
	// 	log.Println(err)
	// }

	// if _, err := result.Prepare(id_basic_information, status, reason, date_time, date_time_approve, date_check, date_time_date_check, message_for_edit, by_employee_code, flag, real_data); err != nil {
	// 	tx.Rollback()
	// 	return nil, err
	// 	log.Println(err)
	// }
	// }

}
