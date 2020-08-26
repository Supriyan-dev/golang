package transportation_application

import (
	"../../../models"
	"log"
	"../../../initialize/enter_the_information"
	utils_enter_the_information "../../../utils/enter_the_information"
	)

type Models_init_basic_information models.DB_init

func (model Models_init_basic_information) Model_GetByIdCommutingBasicInformation(store_number string, employee_number string) (sh []enter_the_information.ShowBasicInformation, err error) {

	rows, err := model.DB.Query(`select cbi.id_commuting_basic_information, cbi.id_general_information,
 									   cbi.insurance_company, cbi.driver_license_expiry_date, cbi.personal_injury,
 									   cbi.property_damage, cbi.car_insurance_document_expiry_date
 									   from commuting_basic_information cbi, basic_information bi,
 									   store_information si , general_information gi 
 									   where cbi.id_general_information = gi.id_general_information and 
 									   gi.id_basic_information = bi.id_basic_information and 
 									   gi.id_store_code = si.id_code_store and si.code_store =? and 
 									   bi.employee_code=?`, store_number, employee_number)

	var init_container enter_the_information.ShowBasicInformation
	if err != nil {
		log.Println(err.Error())
	}

	for rows.Next() {
		err := rows.Scan(&init_container.IdCommutingBasicInformation, &init_container.IdGeneralInformation, &init_container.InsuranceCompany, &init_container.DriverLicenseExpiryDate, &init_container.PersonalInjury, &init_container.PropertyDamage, &init_container.CarInsuranceDocumentExpiryDate)
		if err != nil {
			panic(err.Error())
		}

		sh = append(sh, init_container)

	}

	return sh, nil
}

func (model Models_init_basic_information) Model_InsertBasicInformation(insertD *enter_the_information.InsertBasicInformation) (it []enter_the_information.InsertBasicInformation, condition string) {

	checkdata := utils_enter_the_information.CheckDataById(`select count(*) from commuting_basic_information where id_general_information = ? `, insertD.IdGeneralInformation)
	log.Println(checkdata)
	if checkdata > 1 {
		rows, err := model.DB.Prepare(`update commuting_basic_information set insurance_company = ?, driver_license_expiry_date =?,
 									personal_injury = ?, property_damage = ?, car_insurance_document_expiry_date = ?
 									where id_general_information = ?  `)

		if err != nil {
			log.Println(err.Error())
		}
		execute, err1 := rows.Exec(insertD.InsuranceCompany, insertD.DriverLicenseExpiryDate, insertD.PersonalInjury, insertD.PropertyDamage, insertD.CarInsuranceDocumentExpiryDate, insertD.IdGeneralInformation)

		if err1 != nil && execute == nil {
			log.Println(err1)
			return nil, "Missing required field in body request"
		}

		datainsert := enter_the_information.InsertBasicInformation{
			InsuranceCompany:               insertD.InsuranceCompany,
			DriverLicenseExpiryDate:        insertD.DriverLicenseExpiryDate,
			PersonalInjury:                 insertD.PersonalInjury,
			PropertyDamage:                 insertD.PropertyDamage,
			CarInsuranceDocumentExpiryDate: insertD.CarInsuranceDocumentExpiryDate,
			IdGeneralInformation:           insertD.IdGeneralInformation,
		}

		it = append(it, datainsert)

		return it, "Success Response"

	} else {

		rows, err := model.DB.Prepare(`INSERT INTO commuting_basic_information (insurance_company, driver_license_expiry_date,
 									personal_injury, property_damage, car_insurance_document_expiry_date,id_general_information)
  									VALUES(?,?,?,?,?,?)`)

		if err != nil {
			log.Println(err.Error())
		}

		defer model.DB.Close()

		valid, message := utils_enter_the_information.ValidatorInsertBasicInformation(insertD)

		if valid == false {
			return nil, message
		}

		execute, err1 := rows.Exec(insertD.InsuranceCompany, insertD.DriverLicenseExpiryDate, insertD.PersonalInjury, insertD.PropertyDamage, insertD.CarInsuranceDocumentExpiryDate, insertD.IdGeneralInformation)

		if err1 != nil && execute == nil {
			log.Println(err1)
			return nil, "Missing required field in body request"
		}

		datainsert := enter_the_information.InsertBasicInformation{
			InsuranceCompany:               insertD.InsuranceCompany,
			DriverLicenseExpiryDate:        insertD.DriverLicenseExpiryDate,
			PersonalInjury:                 insertD.PersonalInjury,
			PropertyDamage:                 insertD.PropertyDamage,
			CarInsuranceDocumentExpiryDate: insertD.CarInsuranceDocumentExpiryDate,
			IdGeneralInformation:           insertD.IdGeneralInformation,
		}

		it = append(it, datainsert)

		return it, "Success Response"
	}
}

