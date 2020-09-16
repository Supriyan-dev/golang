package list_input_information

import (
	"log"

	"../../db"
	initialize "../../initialize/permission_to_drive"
	"../../models"
)

type ModelsPermission_init models.DB_init

func (model ModelsPermission_init) ModelPermissionToDrive() (arrJoin []initialize.Join, err error) {
	var join initialize.Join
	db := db.Connect()
	result, err := db.Query(`SELECT store_information.code_store, basic_information.employee_code, basic_information.first_name, 
	basic_information.last_name, commuting_basic_information.driver_license_expiry_date, commuting_basic_information.car_insurance_document_expiry_date, 
	commuting_basic_information.insurance_company, commuting_basic_information.personal_injury, commuting_basic_information.property_damage
	FROM general_information INNER JOIN store_information ON general_information.id_store_code = store_information.id_code_store 
	INNER JOIN basic_information ON general_information.id_basic_information = basic_information.id_basic_information 
	INNER JOIN commuting_basic_information ON commuting_basic_information.id_general_information = general_information.id_general_information`)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	for result.Next() {
		if err := result.Scan(&join.Id_store_code, &join.Employee_code,
			 &join.First_name, &join.Last_name, &join.Driver_license_expiry_date, &join.Car_insurance_document_expiry_date, &join.Insurance_company,
			  &join.Personal_injury, &join.Property_damage); err != nil {
			log.Fatal(err.Error())
		} else {
			arrJoin = append(arrJoin, join)
		}
	}

	return arrJoin, nil
}

func (Model1 ModelsPermission_init) UpdateDataPermissionToDrive(update *initialize.UpdatePermissionToDrive) (arrUpdate []initialize.UpdatePermissionToDrive, err error) {
	db := db.Connect()
	stmt, err := db.Prepare("UPDATE commuting_basic_information SET permitted_to_drive = ?, status_approve = ? WHERE id_commuting_basic_information = ?")
	if err != nil {
		panic(err.Error())
	}
	stmt.Exec(update.Permitted_to_drive, update.Status_approve, update.Id_commuting_basic_information)
	if err != nil {
		log.Println(err)
	}

	Excute := initialize.UpdatePermissionToDrive{
		Id_commuting_basic_information: update.Id_commuting_basic_information,
		Permitted_to_drive:             update.Permitted_to_drive,
		Status_approve:                 update.Status_approve,
	}

	arrUpdate = append(arrUpdate, Excute)

	return arrUpdate, nil
}
