package enter_the_information

import (
	"Go_DX_Services/models"
	"log"
)

type models_init models.DB_init

func (model models_init) Model_GetIdByCodeCommutingGet() (sh []ShowDetailTransportationApplicationGet, err error) {

	//rows, err := model.DB.Query(`select b.id_detail_commuting_trip, b.id_commuting_trip, b.type_of_transport, b.purpose, b.detail_from, b.detail_to,
	//									b.distance, b.cost, b.point_trip, b.transit_point, b.commute_distance, b.go_out_distance
	//								   from commuting_trip a,detail_commuting_trip b, basic_information c , store_information d,
	//								   general_information gen
	//								   where a.id_commuting_trip = b.id_commuting_trip and a.id_general_information = gen.id_general_information
	//								   and gen.id_store_code = d.id_store_code and gen.id_basic_information = c.id_code_store
	//								   and c.employee_code =? and d.code_store =? group by b.id_commuting_trip`, store_number, employee_number)

	rows, err := model.DB.Query(`select b.id_detail_commuting_trip, b.id_commuting_trip, b.type_of_transport, b.purpose, b.detail_from, b.detail_to,
										b.distance, b.cost, b.point_trip, b.transit_point, b.commute_distance, b.go_out_distance
									   from detail_commuting_trip b LIMIT 10`)


	var init_container ShowDetailTransportationApplicationGet
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


func (model models_init) Model_GetIdByCodeCommuting(store_number string, employee_number string) (sh []ShowDetailTransportationApplication, err error) {

	//rows, err := model.DB.Query(`select b.id_detail_commuting_trip, b.id_commuting_trip, b.type_of_transport, b.purpose, b.detail_from, b.detail_to,
	//									b.distance, b.cost, b.point_trip, b.transit_point, b.commute_distance, b.go_out_distance
	//								   from commuting_trip a,detail_commuting_trip b, basic_information c , store_information d,
	//								   general_information gen
	//								   where a.id_commuting_trip = b.id_commuting_trip and a.id_general_information = gen.id_general_information
	//								   and gen.id_store_code = d.id_store_code and gen.id_basic_information = c.id_code_store
	//								   and c.employee_code =? and d.code_store =? group by b.id_commuting_trip`, store_number, employee_number)

	rows, err := model.DB.Query(`select b.id_detail_commuting_trip, b.id_commuting_trip, b.type_of_transport, b.purpose, b.detail_from, b.detail_to,
										b.distance, b.cost, b.point_trip, b.transit_point, b.commute_distance, b.go_out_distance
									   from detail_commuting_trip b where b.id_detail_commuting_trip = ? and b.id_commuting_trip = ?`, store_number, employee_number)


	var init_container ShowDetailTransportationApplication
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

	rows, err := model.DB.Prepare(`INSERT INTO commuting_basic_information (insurance_company, driver_license_expiry_date,
 									personal_injury, property_damage, car_insurance_document_expiry_date,id_general_information)
  									VALUES(?,?,?,?,?,?)`)

	if err != nil {
		panic(err.Error())
	}

	defer model.DB.Close()

	valid, message := ValidatorInsertBasicInformation(insertD)

	if valid == false {
		return nil, message
	}

	execute, err1 := rows.Exec(insertD.InsuranceCompany, insertD.DriverLicenseExpiryDate, insertD.PersonalInjury, insertD.PropertyDamage, insertD.CarInsuranceDocumentExpiryDate, insertD.IdGeneralInformation)
	log.Println(execute)
	log.Println(err1)

	if err1 != nil {
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

func (model models_init) Model_InsertUsageRecordApplyForTravelExpenses(insertD *InsertTransportationApplication) (it []InsertTransportationApplication, condition string) {

	rows, err := model.DB.Prepare(`insert into commuting_trip(id_general_information,route_profile_name,date,attendance_code,code_commuting,created_date,created_time)
 		VALUES(?,?,?,?,?,DATE_FORMAT(CONVERT_TZ(NOW(), @@session.time_zone, '+09:00'),'%Y-%m-%d'),TIME_FORMAT(CONVERT_TZ(NOW(), @@session.time_zone, '+09:00'),'%H:%i:%s'))`)

	if err != nil {
		panic(err.Error())
	}

	defer model.DB.Close()

	valid, message := ValidatorInsertUsageRecordApplyForTravelExpenses(insertD)

	if valid == false {
		return nil, message
	}

	execute, err1 := rows.Exec(insertD.IdGeneralInformation, insertD.RouteProfileName, insertD.Date, insertD.Attendance, insertD.CodeCommuting)
	log.Println(execute)
	log.Println(err1)

	if err1 != nil {
		log.Println(err1)
		return nil, "Missing required field in body request"
	}

	datainsert := InsertTransportationApplication{
		RouteProfileName:     insertD.RouteProfileName,
		Date:                 insertD.Date,
		Attendance:           insertD.Attendance,
		CodeCommuting:        insertD.CodeCommuting,
		IdGeneralInformation: insertD.IdGeneralInformation,
	}

	it = append(it, datainsert)

	return it, "Success Response"

}

func (model models_init) Model_InsertDetailUsageRecordApplyForTravelExpenses(insertDD *InsertDetailTransportationApplication) (itd []InsertDetailTransportationApplication, condition string) {

	rows, err := model.DB.Prepare(`insert into detail_commuting_trip(id_commuting_trip, 
									type_of_Transport ,purpose ,detail_from ,detail_to,distance,
									cost,point_trip,transit_point,commute_distance,go_out_distance) 
									VALUES(?,?,?,?,?,?,?,?,?,?,?)`)

	if err != nil {
		panic(err.Error())
	}
	defer model.DB.Close()

	valid, message := ValidatorDetailInsertUsageRecordApplyForTravelExpenses(insertDD)

	if valid == false {
		return nil, message
	}

	execute, err1 := rows.Exec(insertDD.IdCommutingTrip, insertDD.TypeOfTransport, insertDD.Purpose, insertDD.DetailFrom, insertDD.DetailTo, insertDD.Distance, insertDD.Cost, insertDD.PointTrip, insertDD.TransitPoint, insertDD.CommuteDistance, insertDD.GoOutDistance)
	log.Println(execute)
	if err1 != nil {
		return nil, "Error Execute Please Check Data"
	}
	datainsert := InsertDetailTransportationApplication{
		IdCommutingTrip: insertDD.IdCommutingTrip,
		TypeOfTransport: insertDD.TypeOfTransport,
		Purpose:         insertDD.Purpose,
		DetailFrom:      insertDD.DetailFrom,
		DetailTo:        insertDD.DetailTo,
		Distance:        insertDD.Distance,
		Cost:            insertDD.Cost,
		PointTrip:       insertDD.PointTrip,
		TransitPoint:    insertDD.TransitPoint,
		CommuteDistance: insertDD.CommuteDistance,
		GoOutDistance:   insertDD.GoOutDistance,
	}

	itd = append(itd, datainsert)
	return itd, "Success Response"

}
