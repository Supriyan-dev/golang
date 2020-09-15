package transportation_application

import (
	"../../../initialize/Commuting"
	"../../../models"
	utils_enter_the_information "../../../utils/enter_the_information"
	"errors"
	"log"
)

type Models_init_basic_information models.DB_init

// indonesia
// - nampilin data berdasarkan employee_code dan code_store
// - menggunakan table commuting_basic_information, basic_information, store_inormation dan general_information
// - data di looping dari pencarian di atas

// english
// - displays data based on employee_code and code_store
// - using table commuting_basic_information, basic_information, store_inormation and general_information
// - looped data from the above search

func (model Models_init_basic_information) Model_GetByIdCommutingBasicInformation(store_number string, employee_number string) (sh []Commuting.FormatShowBasicInformation, err error) {

	GetBasicInformation, errGetBasicInformation := model.DB.Query(`select gi.id_general_information
										,bi.id_basic_information,bi.first_name, bi.last_name, bi.adress, bi.adress_kana,
										bi.adress_detail,bi.adress_detail_kana, bi.add_phone_number
										from basic_information bi,store_information si , general_information gi where
										gi.id_basic_information = bi.id_basic_information and 
 									   	gi.id_store_code = si.id_code_store and si.code_store =? and 
 									   	bi.employee_code=? `, store_number, employee_number)

	GetCommutingBasicInformation, errGetCommutingBasicInformation := model.DB.Query(`select cbi.id_commuting_basic_information, cbi.id_general_information,
 									   cbi.insurance_company, cbi.driver_license_expiry_date, cbi.personal_injury,
 									   cbi.property_damage, cbi.car_insurance_document_expiry_date,cbi.status_approve, cbi.date_approve, cbi.time_approve, cbi.date_submit
 									   from commuting_basic_information cbi, basic_information bi,
 									   store_information si , general_information gi 
 									   where cbi.id_general_information = gi.id_general_information and 
 									   gi.id_basic_information = bi.id_basic_information and 
 									   gi.id_store_code = si.id_code_store and si.code_store =? and 
 									   bi.employee_code=? `, store_number, employee_number)
	var init_BasicInformation Commuting.ShowBasicInformation1
	var interface_BasicInformation interface{}
	var init_DataApprove Commuting.ShowBasicInformation2
	var interface_DataApprove interface{}
	var init_CommutingBasicInformation Commuting.ShowBasicInformation3
	var interface_CommutingBasicInformation interface{}
	if errGetBasicInformation != nil && errGetCommutingBasicInformation != nil {
		return nil, errors.New("error basic information and commuting basic information")
		log.Println(errGetBasicInformation.Error())
		log.Println(errGetCommutingBasicInformation.Error())
	}
	defer GetBasicInformation.Close()
	defer GetCommutingBasicInformation.Close()
	GetData1 := GetBasicInformation.Next()
	ScanGetBasicInformation := GetBasicInformation.Scan(&init_DataApprove.IdGeneralBasicInformation, &init_BasicInformation.IdBasicInformation, &init_BasicInformation.FirstName, &init_BasicInformation.LastName, &init_BasicInformation.Address, &init_BasicInformation.AddressKana, &init_BasicInformation.AddressDetail, &init_BasicInformation.AddressDetailKana, &init_BasicInformation.AddPhoneNumber)
	//get divisi code
	var KodeBasicInformation models.NullInt
	GetKodeBasicInformation := model.DB.QueryRow(`SELECT CONCAT(RIGHT(store_information.code_store, 4),
	LPAD(RIGHT(department_information.department_code, 2), 2 , '0'),
	LPAD(RIGHT(store_section_information.store_section_code, 2), 2 , '0'),
	LPAD(RIGHT(unit_information.unit_code, 2), 2 , '0')) AS 'division_code'
	FROM general_information LEFT OUTER JOIN store_information ON general_information.id_store_code = store_information.id_code_store
	LEFT OUTER JOIN department_information ON general_information.id_department = department_information.id_department LEFT OUTER JOIN 
	unit_information ON general_information.id_unit = unit_information.id_unit LEFT OUTER JOIN store_section_information ON
	general_information.id_store_section = store_section_information.id_store_section LEFT OUTER JOIN basic_information ON
	basic_information.id_basic_information = general_information.id_basic_information WHERE basic_information.id_basic_information = ?`, init_BasicInformation.IdBasicInformation).Scan(&KodeBasicInformation)
	if GetKodeBasicInformation != nil {
		log.Println(GetKodeBasicInformation)
		defer GetBasicInformation.Close()
	}
	//end divisi code
	if ScanGetBasicInformation != nil {
		return nil, ScanGetBasicInformation
		log.Println(ScanGetBasicInformation)
	}

	if GetData1 == true {
		showData := Commuting.ShowBasicInformation1{
			IdBasicInformation: init_BasicInformation.IdBasicInformation,
			FirstName:          init_BasicInformation.FirstName,
			LastName:           init_BasicInformation.LastName,
			Address:            init_BasicInformation.Address,
			AddressKana:        init_BasicInformation.AddressKana,
			AddressDetail:      init_BasicInformation.AddressDetail,
			AddressDetailKana:  init_BasicInformation.AddressKana,
			AddPhoneNumber:     init_BasicInformation.AddPhoneNumber,
			DivisionCode:       KodeBasicInformation,
		}
		interface_BasicInformation = showData
	} else {
		interface_BasicInformation = nil
	}

	GetData2 := GetCommutingBasicInformation.Next()
	ScanCommutingBasicInformation := GetCommutingBasicInformation.Scan(&init_CommutingBasicInformation.IdCommutingBasicInformation, &init_CommutingBasicInformation.IdGeneralInformation, &init_CommutingBasicInformation.InsuranceCompany, &init_CommutingBasicInformation.DriverLicenseExpiryDate, &init_CommutingBasicInformation.PersonalInjury, &init_CommutingBasicInformation.PropertyDamage, &init_CommutingBasicInformation.CarInsuranceDocumentExpiryDate, &init_DataApprove.StatusApproved, &init_DataApprove.DateApprove, &init_DataApprove.TimeApprove, &init_DataApprove.DateSubmit)

	if ScanCommutingBasicInformation != nil {
		log.Println(ScanCommutingBasicInformation)
	}

	if GetData2 == true {
		showData2 := Commuting.ShowBasicInformation2{
			IdGeneralBasicInformation: init_DataApprove.IdGeneralBasicInformation,
			StatusApproved:            init_DataApprove.StatusApproved,
			DateApprove:               init_DataApprove.DateApprove,
			TimeApprove:               init_DataApprove.TimeApprove,
			DateSubmit:                init_DataApprove.DateSubmit,
		}
		interface_DataApprove = showData2
		showData3 := Commuting.ShowBasicInformation3{
			IdCommutingBasicInformation:    init_CommutingBasicInformation.IdCommutingBasicInformation,
			IdGeneralInformation:           init_CommutingBasicInformation.IdGeneralInformation,
			InsuranceCompany:               init_CommutingBasicInformation.InsuranceCompany,
			DriverLicenseExpiryDate:        init_CommutingBasicInformation.DriverLicenseExpiryDate,
			PersonalInjury:                 init_CommutingBasicInformation.PersonalInjury,
			PropertyDamage:                 init_CommutingBasicInformation.PropertyDamage,
			CarInsuranceDocumentExpiryDate: init_CommutingBasicInformation.CarInsuranceDocumentExpiryDate,
		}
		interface_CommutingBasicInformation = showData3
	} else {
		interface_DataApprove = nil
		interface_CommutingBasicInformation = nil
	}

	FinallyData := Commuting.FormatShowBasicInformation{
		DataBasicInformation: interface_BasicInformation,
		DataApprove:          interface_DataApprove,
		DataApply:            interface_CommutingBasicInformation,
	}
	sh = append(sh, FinallyData)

	return sh, nil
}

// indonesia
// check jumlah data dari commuting_basic_information berdasarkan id_general_information
// jika ada akan melakukan update ke table commuting_basic_information
// jika tidak ada akan melakukan insert ke table commuting_basic_information

// english
// check the amount of data from commuting_basic_information based on id_general_information
// if there is an update to the commuting_basic_information table by id_general_information
// if not present will insert into the commuting_basic_information table

func (model Models_init_basic_information) Model_InsertBasicInformation(insertD *Commuting.InsertBasicInformation,employee_number string) (it []Commuting.InsertBasicInformation, condition string) {

	checkdata := utils_enter_the_information.CheckDataById(`select count(*) from commuting_basic_information where id_commuting_basic_information = ? `, insertD.IdCommutingBasicInformation)
	var	DataIdGeralInformation string
	GetIdGeneralInformation := model.DB.QueryRow(`select gi.id_general_information from general_information gi, basic_information bi where gi.id_basic_information = bi.id_basic_information and bi.employee_code = ?`,employee_number).Scan(&DataIdGeralInformation)

	if GetIdGeneralInformation != nil {
		return nil, GetIdGeneralInformation.Error()
	}

	if DataIdGeralInformation == ""{
		return nil, "general information not found"
	}

	if checkdata > 0 {
		rows, err := model.DB.Prepare(`update commuting_basic_information set insurance_company = ?, driver_license_expiry_date =?,
 									personal_injury = ?, property_damage = ?, car_insurance_document_expiry_date = ?
 									where id_commuting_basic_information = ?  `)

		if err != nil {
			log.Println(err.Error())
		}
		defer rows.Close()
		execute, err1 := rows.Exec(insertD.InsuranceCompany, insertD.DriverLicenseExpiryDate, insertD.PersonalInjury, insertD.PropertyDamage, insertD.CarInsuranceDocumentExpiryDate, insertD.IdCommutingBasicInformation)

		if err1 != nil && execute == nil {
			log.Println(err1)
			return nil, "Missing required field in body request"
		}

		dataupdate := Commuting.InsertBasicInformation{
			IdCommutingBasicInformation:    insertD.IdCommutingBasicInformation,
			InsuranceCompany:               insertD.InsuranceCompany,
			DriverLicenseExpiryDate:        insertD.DriverLicenseExpiryDate,
			PersonalInjury:                 insertD.PersonalInjury,
			PropertyDamage:                 insertD.PropertyDamage,
			CarInsuranceDocumentExpiryDate: insertD.CarInsuranceDocumentExpiryDate,
		}

		it = append(it, dataupdate)

		return it, "Success Response"

	} else {

		rows, err := model.DB.Prepare(`INSERT INTO commuting_basic_information (insurance_company, driver_license_expiry_date,
 									personal_injury, property_damage, car_insurance_document_expiry_date,id_general_information)
  									VALUES(?,?,?,?,?,?)`)

		if err != nil {
			log.Println(err.Error())
		}

		defer rows.Close()

		valid, message := utils_enter_the_information.ValidatorInsertBasicInformation(insertD)

		if valid == false {
			return nil, message
		}

		execute, err1 := rows.Exec(insertD.InsuranceCompany, insertD.DriverLicenseExpiryDate, insertD.PersonalInjury, insertD.PropertyDamage, insertD.CarInsuranceDocumentExpiryDate, DataIdGeralInformation)

		if err1 != nil && execute == nil {
			log.Println(err1)
			return nil, "Missing required field in body request"
		}

		datainsert := Commuting.InsertBasicInformation{
			IdCommutingBasicInformation:    "Auto",
			InsuranceCompany:               insertD.InsuranceCompany,
			DriverLicenseExpiryDate:        insertD.DriverLicenseExpiryDate,
			PersonalInjury:                 insertD.PersonalInjury,
			PropertyDamage:                 insertD.PropertyDamage,
			CarInsuranceDocumentExpiryDate: insertD.CarInsuranceDocumentExpiryDate,
		}

		it = append(it, datainsert)

		return it, "Success Response"
	}
}
