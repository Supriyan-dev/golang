package transportation_application

import (
	"../../../initialize/Commuting"
	"../../../models"
	utils_enter_the_information "../../../utils/enter_the_information"
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
 									   	bi.employee_code=?`, store_number, employee_number)

	GetCommutingBasicInformation, errGetCommutingBasicInformation := model.DB.Query(`select cbi.id_commuting_basic_information, cbi.id_general_information,
 									   cbi.insurance_company, cbi.driver_license_expiry_date, cbi.personal_injury,
 									   cbi.property_damage, cbi.car_insurance_document_expiry_date,cbi.status_approve, cbi.date_approve, cbi.time_approve, cbi.date_submit
 									   from commuting_basic_information cbi, basic_information bi,
 									   store_information si , general_information gi 
 									   where cbi.id_general_information = gi.id_general_information and 
 									   gi.id_basic_information = bi.id_basic_information and 
 									   gi.id_store_code = si.id_code_store and si.code_store =? and 
 									   bi.employee_code=?`, store_number, employee_number)
	var init_BasicInformation Commuting.ShowBasicInformation1
	var Arr_BasicInformation []Commuting.ShowBasicInformation1
	var init_DataApprove Commuting.ShowBasicInformation2
	var Arr_DataApprove []Commuting.ShowBasicInformation2
	var init_CommutingBasicInformation Commuting.ShowBasicInformation3
	var Arr_init_CommutingBasicInformation []Commuting.ShowBasicInformation3
	if errGetBasicInformation != nil && errGetCommutingBasicInformation != nil {

		log.Println(errGetBasicInformation.Error())
		log.Println(errGetCommutingBasicInformation.Error())
	}

	GetData1 := GetBasicInformation.Next()
	ScanGetBasicInformation := GetBasicInformation.Scan(&init_DataApprove.IdGeneralBasicInformation, &init_BasicInformation.IdBasicInformation, &init_BasicInformation.FirstName, &init_BasicInformation.LastName, &init_BasicInformation.Address, &init_BasicInformation.AddressKana, &init_BasicInformation.AddressDetail, &init_BasicInformation.AddressDetailKana, &init_BasicInformation.AddPhoneNumber)

	if ScanGetBasicInformation != nil {
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
		}
		Arr_BasicInformation = append(Arr_BasicInformation, showData)
	} else {
		Arr_BasicInformation = nil
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
		Arr_DataApprove = append(Arr_DataApprove, showData2)
		showData3 := Commuting.ShowBasicInformation3{
			IdCommutingBasicInformation:    init_CommutingBasicInformation.IdCommutingBasicInformation,
			IdGeneralInformation:           init_CommutingBasicInformation.IdGeneralInformation,
			InsuranceCompany:               init_CommutingBasicInformation.InsuranceCompany,
			DriverLicenseExpiryDate:        init_CommutingBasicInformation.DriverLicenseExpiryDate,
			PersonalInjury:                 init_CommutingBasicInformation.PersonalInjury,
			PropertyDamage:                 init_CommutingBasicInformation.PropertyDamage,
			CarInsuranceDocumentExpiryDate: init_CommutingBasicInformation.CarInsuranceDocumentExpiryDate,
		}
		Arr_init_CommutingBasicInformation = append(Arr_init_CommutingBasicInformation, showData3)
	} else {
		Arr_DataApprove = nil
		Arr_init_CommutingBasicInformation = nil
	}

	FinallyData := Commuting.FormatShowBasicInformation{
		//DataBasicInformation: Arr_BasicInformation,
		//DataApprove:          Arr_DataApprove,
		//DataApply:            Arr_init_CommutingBasicInformation,
		DataBasicInformation: init_BasicInformation,
		DataApprove:          init_DataApprove,
		DataApply:            init_DataApprove,
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

func (model Models_init_basic_information) Model_InsertBasicInformation(insertD *Commuting.InsertBasicInformation) (it []Commuting.InsertBasicInformation, condition string) {

	checkdata := utils_enter_the_information.CheckDataById(`select count(*) from commuting_basic_information where id_general_information = ? `, insertD.IdGeneralInformation)
	log.Println(checkdata)
	if checkdata > 1 {
		rows, err := model.DB.Prepare(`update commuting_basic_information set insurance_company = ?, driver_license_expiry_date =?,
 									personal_injury = ?, property_damage = ?, car_insurance_document_expiry_date = ?
 									where id_commuting_basic_information = ?  `)

		if err != nil {
			log.Println(err.Error())
		}
		execute, err1 := rows.Exec(insertD.InsuranceCompany, insertD.DriverLicenseExpiryDate, insertD.PersonalInjury, insertD.PropertyDamage, insertD.CarInsuranceDocumentExpiryDate, insertD.IdCommutingBasicInformation)

		if err1 != nil && execute == nil {
			log.Println(err1)
			return nil, "Missing required field in body request"
		}

		datainsert := Commuting.InsertBasicInformation{
			IdCommutingBasicInformation:    insertD.IdCommutingBasicInformation,
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

		datainsert := Commuting.InsertBasicInformation{
			IdCommutingBasicInformation:    "Auto",
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
