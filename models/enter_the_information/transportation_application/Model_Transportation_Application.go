package transportation_application

import (
	"../../../models"
	"log"
	"../../../initialize/enter_the_information"
	utils_enter_the_information "../../../utils/enter_the_information"
	)

type Models_init_Usage_Record models.DB_init

func (model Models_init_Usage_Record) Model_GetByIdUsageRecord(store_number string, employee_number string) (sh []enter_the_information.ShowUsageRecord, err error) {

	rows, err := model.DB.Query(`select b.id_detail_commuting_trip, b.id_commuting_trip, b.type_of_transport, b.purpose, b.detail_from, b.detail_to,
										b.distance, b.cost, b.point_trip, b.transit_point, b.commute_distance, b.go_out_distance
										from basic_information bi, commuting_trip ct, detail_commuting_trip b, store_information si , general_information gi
										where ct.id_commuting_trip = b.id_commuting_trip and gi.id_basic_information = bi.id_basic_information 
										and gi.id_store_code = si.id_code_store and ct.id_general_information = gi.id_general_information and si.code_store =? and bi.employee_code=?
										group by b.id_commuting_trip`, store_number, employee_number)

	var init_container enter_the_information.ShowUsageRecord
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

func (model Models_init_Usage_Record) Model_InsertUsageRecordApplyForTravelExpenses(insertD *enter_the_information.InsertTransportationApplication) (it []enter_the_information.InsertTransportationApplication, condition string) {
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

	valid, message := utils_enter_the_information.ValidatorInsertUsageRecordApplyForTravelExpenses(insertD)
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

	datainsert := enter_the_information.InsertTransportationApplication{
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
