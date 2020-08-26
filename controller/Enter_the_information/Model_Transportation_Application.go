package enter_the_information

import (
	"../../models"
	"log"
)

type models_init models.DB_init

func (model models_init) Model_GetByIdCommutingBasicInformation(store_number string, employee_number string) (sh []ShowBasicInformation, err error) {

	rows, err := model.DB.Query(`select cbi.id_commuting_basic_information, cbi.id_general_information,
 									   cbi.insurance_company, cbi.driver_license_expiry_date, cbi.personal_injury,
 									   cbi.property_damage, cbi.car_insurance_document_expiry_date
 									   from commuting_basic_information cbi, basic_information bi,
 									   store_information si , general_information gi 
 									   where cbi.id_general_information = gi.id_general_information and 
 									   gi.id_basic_information = bi.id_basic_information and 
 									   gi.id_store_code = si.id_code_store and si.code_store =? and 
 									   bi.employee_code=?`, store_number, employee_number)

	var init_container ShowBasicInformation
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

func (model models_init) Model_GetByIdUsageRecord(store_number string, employee_number string) (sh []ShowUsageRecord, err error) {

	rows, err := model.DB.Query(`select b.id_detail_commuting_trip, b.id_commuting_trip, b.type_of_transport, b.purpose, b.detail_from, b.detail_to,
										b.distance, b.cost, b.point_trip, b.transit_point, b.commute_distance, b.go_out_distance
										from basic_information bi, commuting_trip ct, detail_commuting_trip b, store_information si , general_information gi
										where ct.id_commuting_trip = b.id_commuting_trip and gi.id_basic_information = bi.id_basic_information 
										and gi.id_store_code = si.id_code_store and ct.id_general_information = gi.id_general_information and si.code_store =? and bi.employee_code=?
										group by b.id_commuting_trip`, store_number, employee_number)

	var init_container ShowUsageRecord
	if err != nil {
		log.Println(err.Error())
	}

	for rows.Next() {
		err := rows.Scan(&init_container.IdDetailCommutingTrip, &init_container.IdCommutingTrip, &init_container.TypeOfTransport, &init_container.Purpose, &init_container.DetailFrom, &init_container.DetailTo, &init_container.Distance, &init_container.Cost, &init_container.PointTrip, &init_container.TransitPoint, &init_container.CommuteDistance, &init_container.GoOutDistance)
		if err != nil {
			panic(err.Error())
		}

		sh = append(sh, init_container)

	}

	return sh, nil
}

func (model models_init) Model_InsertBasicInformation(insertD *InsertBasicInformation) (it []InsertBasicInformation, condition string) {

	checkdata := CheckDataById(`select count(*) from commuting_basic_information where id_general_information = ? `, insertD.IdGeneralInformation)
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

		datainsert := InsertBasicInformation{
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

		valid, message := ValidatorInsertBasicInformation(insertD)

		if valid == false {
			return nil, message
		}

		execute, err1 := rows.Exec(insertD.InsuranceCompany, insertD.DriverLicenseExpiryDate, insertD.PersonalInjury, insertD.PropertyDamage, insertD.CarInsuranceDocumentExpiryDate, insertD.IdGeneralInformation)

		if err1 != nil && execute == nil {
			log.Println(err1)
			return nil, "Missing required field in body request"
		}

		datainsert := InsertBasicInformation{
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

func (model models_init) Model_InsertUsageRecordApplyForTravelExpenses(insertD *InsertTransportationApplication) (it []InsertTransportationApplication, condition string) {
	vals := []interface{}{}

	rows, err := model.DB.Prepare(`insert into commuting_trip(id_general_information,route_profile_name,date,attendance_code,code_commuting,created_date,created_time)
 		VALUES(?,?,?,?,?,DATE_FORMAT(CONVERT_TZ(NOW(), @@session.time_zone, '+09:00'),'%Y-%m-%d'),TIME_FORMAT(CONVERT_TZ(NOW(), @@session.time_zone, '+09:00'),'%H:%i:%s'))`)

	sqlDetail := `insert into detail_commuting_trip(id_commuting_trip,
									type_of_Transport ,purpose ,detail_from ,detail_to,distance,
									cost,point_trip,transit_point,commute_distance,go_out_distance)
									VALUES`

	for _, insertDD := range insertD.DataDetail {
		sqlDetail += "(?,?,?,?,?,?,?,?,?,?,?),"
		vals = append(vals, insertDD.IdCommutingTrip, insertDD.TypeOfTransport, insertDD.Purpose, insertDD.DetailFrom, insertDD.DetailTo, insertDD.Distance, insertDD.Cost, insertDD.PointTrip, insertDD.TransitPoint, insertDD.CommuteDistance, insertDD.GoOutDistance)
	}
	sqlDetail = sqlDetail[0 : len(sqlDetail)-1]
	stmtDetail, _ := model.DB.Prepare(sqlDetail)

	if err != nil {
		panic(err.Error())
	}

	defer model.DB.Close()

	valid, message := ValidatorInsertUsageRecordApplyForTravelExpenses(insertD)
	if valid == false {
		return nil, message
	}

	execute, err1 := rows.Exec(insertD.IdGeneralInformation, insertD.RouteProfileName, insertD.Date, insertD.Attendance, insertD.CodeCommuting)
	res, _ := stmtDetail.Exec(vals...)
	if res == nil {
		log.Println("gagal")
	}
	if err1 != nil && execute == nil {
		log.Println(err1)
		return nil, "Missing required field in body request"
	}

	datainsert := InsertTransportationApplication{
		RouteProfileName:     insertD.RouteProfileName,
		Date:                 insertD.Date,
		Attendance:           insertD.Attendance,
		CodeCommuting:        insertD.CodeCommuting,
		IdGeneralInformation: insertD.IdGeneralInformation,
		DataDetail:           insertD.DataDetail,
	}

	it = append(it, datainsert)

	return it, "Success Response"

}

//func (model models_init) Model_InsertDetailUsageRecordApplyForTravelExpenses(insertDD *InsertDetailTransportationApplication) (itd []InsertDetailTransportationApplication, condition string) {
//
//	rows, err := model.DB.Prepare(`insert into detail_commuting_trip(id_commuting_trip,
//									type_of_Transport ,purpose ,detail_from ,detail_to,distance,
//									cost,point_trip,transit_point,commute_distance,go_out_distance)
//									VALUES(?,?,?,?,?,?,?,?,?,?,?)`)
//
//	if err != nil {
//		panic(err.Error())
//	}
//	defer model.DB.Close()
//
//	valid, message := ValidatorDetailInsertUsageRecordApplyForTravelExpenses(insertDD)
//
//	if valid == false {
//		return nil, message
//	}
//
//	execute, err1 := rows.Exec(insertDD.IdCommutingTrip, insertDD.TypeOfTransport, insertDD.Purpose, insertDD.DetailFrom, insertDD.DetailTo, insertDD.Distance, insertDD.Cost, insertDD.PointTrip, insertDD.TransitPoint, insertDD.CommuteDistance, insertDD.GoOutDistance)
//	if err1 != nil && execute == nil {
//		return nil, "Missing required field in body request"
//	}
//	datainsert := InsertDetailTransportationApplication{
//		IdCommutingTrip: insertDD.IdCommutingTrip,
//		TypeOfTransport: insertDD.TypeOfTransport,
//		Purpose:         insertDD.Purpose,
//		DetailFrom:      insertDD.DetailFrom,
//		DetailTo:        insertDD.DetailTo,
//		Distance:        insertDD.Distance,
//		Cost:            insertDD.Cost,
//		PointTrip:       insertDD.PointTrip,
//		TransitPoint:    insertDD.TransitPoint,
//		CommuteDistance: insertDD.CommuteDistance,
//		GoOutDistance:   insertDD.GoOutDistance,
//	}
//
//	itd = append(itd, datainsert)
//	return itd, "Success Response"
//
//}
