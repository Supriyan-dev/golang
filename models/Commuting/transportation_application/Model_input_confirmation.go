package transportation_application

import (
	"../../../initialize/Commuting"
	"../../../models"
	utils_enter_the_information "../../../utils/enter_the_information"
	"log"
)

type Models_init_input_confirmation models.DB_init

func (model Models_init_input_confirmation) GetDataInputConfimation(store_number string, employee_number string) (sh []Commuting.IC_Format, err error) {

	var init_CommutingBasicInformation Commuting.ShowBasicInformation3

	GetBasicInformation, errGetBasicInformation := model.DB.Query(`select bi.id_basic_information,bi.first_name, bi.last_name, bi.adress, bi.adress_kana,
										bi.adress_detail,bi.adress_detail_kana, bi.add_phone_number
										from basic_information bi,store_information si , general_information gi where
										gi.id_basic_information = bi.id_basic_information and 
 									   	gi.id_store_code = si.id_code_store and si.code_store =? and 
 									   	bi.employee_code=?`, store_number, employee_number)
	if errGetBasicInformation != nil {
		log.Println(errGetBasicInformation)
	}
	defer GetBasicInformation.Close()

	GetCommutingBasicInformation, errGetCommutingBasicInformation := model.DB.Query(`select cbi.id_commuting_basic_information, cbi.id_general_information,
 									   cbi.insurance_company, cbi.driver_license_expiry_date, cbi.personal_injury,
 									   cbi.property_damage, cbi.car_insurance_document_expiry_date
 									   from commuting_basic_information cbi, basic_information bi,
 									   store_information si , general_information gi 
 									   where cbi.id_general_information = gi.id_general_information and 
 									   gi.id_basic_information = bi.id_basic_information and 
 									   gi.id_store_code = si.id_code_store and si.code_store =? and 
 									   bi.employee_code=?`, store_number, employee_number)
	if errGetCommutingBasicInformation != nil {
		log.Println(errGetCommutingBasicInformation)
	}
	defer GetCommutingBasicInformation.Close()
	rows, err := model.DB.Query(`select  ct.date,ct.route_profile_name,MIN(b.id_commuting_trip),COALESCE(SUM(b.distance),0)
 										as distance,COALESCE(SUM(commute_distance),0) as commute_distance, COALESCE(SUM(b.cost),0) as cost , 
 										MIN(ct.draft),MIN(b.purpose)
 										 from basic_information bi, commuting_trip ct, detail_commuting_trip b, store_information si , general_information gi, 
										master_transportation trans
										where ct.id_commuting_trip = b.id_commuting_trip and gi.id_basic_information = bi.id_basic_information
										and b.type_of_transport =  trans.code_transportation
										and gi.id_store_code = si.id_code_store and ct.id_general_information = gi.id_general_information and si.code_store =? and bi.employee_code=?
										and ct.submit ='N' and ct.save_trip ='N'
										group by b.id_commuting_trip`, store_number, employee_number)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	var init_biC interface{}
	var init_bi Commuting.ShowBasicInformation1
	//var Arr_bi []Commuting.ShowBasicInformation1
	var init_ur Commuting.ShowUsageRecord2
	var Arr_ur []Commuting.ShowUsageRecord2

	GetCommutingBasicInformation.Next()

	ScanCommutingBasicInformation := GetCommutingBasicInformation.Scan(&init_CommutingBasicInformation.IdCommutingBasicInformation, &init_CommutingBasicInformation.IdGeneralInformation, &init_CommutingBasicInformation.InsuranceCompany, &init_CommutingBasicInformation.DriverLicenseExpiryDate, &init_CommutingBasicInformation.PersonalInjury, &init_CommutingBasicInformation.PropertyDamage, &init_CommutingBasicInformation.CarInsuranceDocumentExpiryDate)

	if ScanCommutingBasicInformation != nil {
		log.Println(ScanCommutingBasicInformation)
	}

	GetBasicInformation.Next()
	errScanBasicInformation := GetBasicInformation.Scan(&init_bi.IdBasicInformation, &init_bi.FirstName, &init_bi.LastName, &init_bi.Address, &init_bi.AddressKana, &init_bi.AddressDetail, &init_bi.AddressDetailKana, &init_bi.AddPhoneNumber)
	var KodeBasicInformation models.NullInt
	GetKodeBasicInformation := model.DB.QueryRow(`SELECT CONCAT(RIGHT(store_information.code_store, 4),
	LPAD(RIGHT(department_information.department_code, 2), 2 , '0'),
	LPAD(RIGHT(store_section_information.store_section_code, 2), 2 , '0'),
	LPAD(RIGHT(unit_information.unit_code, 2), 2 , '0')) AS 'division_code'
	FROM general_information LEFT OUTER JOIN store_information ON general_information.id_store_code = store_information.id_code_store
	LEFT OUTER JOIN department_information ON general_information.id_department = department_information.id_department LEFT OUTER JOIN 
	unit_information ON general_information.id_unit = unit_information.id_unit LEFT OUTER JOIN store_section_information ON
	general_information.id_store_section = store_section_information.id_store_section LEFT OUTER JOIN basic_information ON
	basic_information.id_basic_information = general_information.id_basic_information WHERE basic_information.id_basic_information = ?`, init_bi.IdBasicInformation).Scan(&KodeBasicInformation)

	if GetKodeBasicInformation != nil {
		log.Println(GetKodeBasicInformation)
	}

	if errScanBasicInformation != nil {
		init_biC = nil
	} else {
		init_biC = init_bi
	}
	StatusTemporari := ""
	StatusDriversLicense := ""
	StatusCarInsurance := ""

	CheckStatusDriverLicense := utils_enter_the_information.CheckDataByIdNullString(`select COUNT(*) from commuting_basic_information a where a.driver_license_expiry_date = ? and a.driver_license_expiry_date = DATE_FORMAT(CONVERT_TZ(NOW(), @@session.time_zone, '+09:00'),'%Y-%m-%d')`, init_CommutingBasicInformation.DriverLicenseExpiryDate)

	if CheckStatusDriverLicense > 1 {
		//no
		StatusDriversLicense = `いいえ`
	} else {
		//yes
		StatusDriversLicense = `はい`
	}

	CheckStatusCarInsurancce := utils_enter_the_information.CheckDataByIdNullString(`select COUNT(*) from commuting_basic_information a where a.car_insurance_document_expiry_date = ? and a.car_insurance_document_expiry_date = DATE_FORMAT(CONVERT_TZ(NOW(), @@session.time_zone, '+09:00'),'%Y-%m-%d')`, init_CommutingBasicInformation.CarInsuranceDocumentExpiryDate)

	if CheckStatusCarInsurancce > 1 {
		//no
		StatusCarInsurance = `いいえ`
	} else {
		//yes
		StatusCarInsurance = `はい`
	}
	for rows.Next() {
		err := rows.Scan(&init_ur.Date,&init_ur.RouteProfileName,&init_ur.IdCommutingTrip, &init_ur.Distance, &init_ur.CommuteDistance, &init_ur.Cost, &StatusTemporari, &init_ur.Purpose)
		//err := rows.Scan(&init_ur.IdDetailCommutingTrip, &init_ur.IdCommutingTrip, &init_ur.TypeOfTransport, &init_ur.Purpose, &init_ur.DetailFrom, &init_ur.DetailTo, &init_ur.Distance, &init_ur.Cost, &init_ur.PointTrip, &init_ur.TransitPoint, &init_ur.CommuteDistance, &init_ur.GoOutDistance)
		if err != nil {
			log.Println(err.Error())
			Arr_ur = nil
		} else {
			if StatusTemporari == "Y" {
				//yes
				StatusTemporari = "はい"
			} else {
				//no
				StatusTemporari = "いいえ"
			}
			DatatypeOfTransportation, DataPurpose, DataRoute := utils_enter_the_information.GetAdditionalUsageRecord(store_number, employee_number, init_ur.IdCommutingTrip, `usageRecord-CheckData`)

			dataCommutingTrip := Commuting.ShowUsageRecord2{
				IdCommutingTrip:  init_ur.IdCommutingTrip,
				RouteProfileName: init_ur.RouteProfileName,
				Date:             init_ur.Date,
				TypeOfTransport:  DatatypeOfTransportation,
				Purpose:          DataPurpose,
				Route:            DataRoute,
				Distance:         init_ur.Distance,
				CommuteDistance:  init_ur.CommuteDistance,
				Cost:             init_ur.Cost,
				StatusTemporary:  StatusTemporari,
			}
			Arr_ur = append(Arr_ur, dataCommutingTrip)
		}
	}

	if init_biC != nil && Arr_ur != nil {
		FinallyData := Commuting.IC_Format{
			StatusDriversLicense: StatusDriversLicense,
			StatusCarInsurance:   StatusCarInsurance,
			KodeBasicInformation: KodeBasicInformation,
			DataBasic:            init_biC,
			DataDetail:           Arr_ur,
			DataCommutingBasic:   init_CommutingBasicInformation,
		}
		sh = append(sh, FinallyData)
		return sh, nil
	}
	return nil, nil
}

func (model Models_init_input_confirmation) Model_SubmitInputConfirmation(id string) (response int64, condition string) {

	sqlUpdate := `update commuting_trip set submit = 'Y', 
	date_submit= DATE_FORMAT(CONVERT_TZ(NOW(), @@session.time_zone, '+09:00'),'%Y-%m-%d'),
	time_submit= TIME_FORMAT(CONVERT_TZ(NOW(), @@session.time_zone, '+09:00'),'%H:%i:%s') 
	where id_commuting_trip IN(` + id + `)`

	stmtUpdate, errStmtUpdate := model.DB.Query(sqlUpdate)

	if errStmtUpdate != nil {
		return 0, "Please Check Your ID"
	}
	defer stmtUpdate.Close()

	return 1, "Success Response"
}
