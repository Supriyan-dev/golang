package transportation_application

import (
	"../../../initialize/enter_the_information"
	"../../../models"
	utils_enter_the_information "../../../utils/enter_the_information"
	"log"
)

type Models_init_Usage_Record models.DB_init

// indonesia
// Menampilkan semua detail_commuting_trip berdasarkan code_store dan employee_store di group by berdasarkan id_commuting_trip
// data di looping

// english
// Show all detail_commuting_trip based on code_store and employee_store in group by based on id_commuting_trip
// data is looped

func (model Models_init_Usage_Record) Model_GetByIdUsageRecord(store_number string, employee_number string) (sh []enter_the_information.ShowUsageRecord, err error) {

	rows, err := model.DB.Query(`select b.id_detail_commuting_trip, b.id_commuting_trip, trans.name_transportation_japanese, b.purpose, b.detail_from, b.detail_to,
										b.distance, b.cost, b.point_trip, b.transit_point, b.commute_distance, b.go_out_distance
										from basic_information bi, commuting_trip ct, detail_commuting_trip b, store_information si , general_information gi, 
										master_transportation trans
										where ct.id_commuting_trip = b.id_commuting_trip and gi.id_basic_information = bi.id_basic_information
										and b.type_of_transport =  trans.code_transportation
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

// indonesia
// Menampilkan data Usage Record untuk di edit berdasarkan id commuting trip ,store number dan employee number

// english
// Show Data Usage Record to edit by id commuting trip, store number dan employee number
func (model Models_init_Usage_Record) Model_GetByIdUsageRecordForEdit(store_number string, employee_number string, id_commuting_trip string) (sh []enter_the_information.ShowUsageRecord, err error) {

	rows, err := model.DB.Query(`select b.id_detail_commuting_trip, b.id_commuting_trip, trans.name_transportation_japanese, b.purpose, b.detail_from, b.detail_to,
										b.distance, b.cost, b.point_trip, b.transit_point, b.commute_distance, b.go_out_distance
										from basic_information bi, commuting_trip ct, detail_commuting_trip b, store_information si , general_information gi,
										master_transportation trans
										where ct.id_commuting_trip = b.id_commuting_trip and gi.id_basic_information = bi.id_basic_information 
										and b.type_of_transport =  trans.code_transportation
										and gi.id_store_code = si.id_code_store and ct.id_general_information = gi.id_general_information 
										and si.code_store =? and bi.employee_code=? and b.id_commuting_trip = ?
										`, store_number, employee_number, id_commuting_trip)

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
// indonesia
// Menampilkan Semua route favorit berdasarkan store number dan employee number

// english
// get all data route favorite by store number and employee number

func (model Models_init_Usage_Record) Model_GetByIdUsageRecordUseMyRoute(store_number string, employee_number string) (sh []enter_the_information.ShowUseMyRoute, err error) {

	rows, err := model.DB.Query(`select comtrip.id_commuting_trip, detcomtrip.id_detail_commuting_trip, comtrip.route_profile_name, detcomtrip.type_of_transport, comtrip.attendance_code,
detcomtrip.purpose, detcomtrip.detail_to, detcomtrip.detail_from, detcomtrip.commute_distance, detcomtrip.go_out_distance, detcomtrip.cost from commuting_trip comtrip, detail_commuting_trip detcomtrip, general_information geninfo, basic_information bainfo, store_information storeinfo
where comtrip.id_commuting_trip = detcomtrip.id_commuting_trip and geninfo.id_general_information = comtrip.id_general_information AND
geninfo.id_basic_information = bainfo.id_basic_information and geninfo.id_store_code = storeinfo.id_code_store  and storeinfo.code_store =? and bainfo.employee_code =? and comtrip.save_trip ='Y'`, store_number, employee_number)

	var init_container enter_the_information.ShowUseMyRoute
	if err != nil {
		log.Println(err.Error())
	}

	for rows.Next() {
		err := rows.Scan(&init_container.IdCommutingTrip, &init_container.IdDetailCommutingTrip, &init_container.RouteProfileName, &init_container.TypeOfTransport, &init_container.AttendanceCode, &init_container.Purpose, &init_container.DetailTo, &init_container.DetailFrom, &init_container.CommuteDistance, &init_container.GoOutDistance, &init_container.Cost)
		if err != nil {
			panic(err.Error())
		}

		sh = append(sh, init_container)

	}

	return sh, nil
}

// indonesia
// insert commuting_trip dan detail_commuting_trip
// data insert dalam row / json
// con cuma bisa di isi Y/N jika Y maka nampil di use my route jika N akan nampil langsung Confirmation of submission contents

// english
// insert commuting_trip and detail_commuting_trip
// body row -> json

func (model Models_init_Usage_Record) Model_InsertUsageRecordApplyForTravelExpenses(con string, store_id string, employee_id string, initializeData *enter_the_information.InsertUsageRecordApplyForTravelExpenses) (it []enter_the_information.InsertUsageRecordApplyForTravelExpenses, condition string) {

	if con != "Y" && con != "N" {
		return nil, "Missing required, Please use /Y or /N"
	}
	checkCountData := utils_enter_the_information.CheckDataByStoreAndEmployee(`select COUNT(*) from commuting_trip comtrip, detail_commuting_trip detcomtrip, general_information geninfo, basic_information bainfo, store_information storeinfo where comtrip.id_commuting_trip = detcomtrip.id_commuting_trip and geninfo.id_general_information = comtrip.id_general_information AND geninfo.id_basic_information = bainfo.id_basic_information and geninfo.id_store_code = storeinfo.id_code_store and storeinfo.code_store =? and bainfo.employee_code =? and comtrip.save_trip ='Y'`, store_id, employee_id)
	log.Println(checkCountData)
	if con == "Y" && checkCountData >= 3 {
		return nil, "You cannot register up to more than 3 routes"
	}
	vals := []interface{}{}

	rows, err := model.DB.Prepare(`insert into commuting_trip(id_general_information,route_profile_name,date,attendance_code,code_commuting,created_date,created_time,save_trip)
 		VALUES(?,?,?,?,?,DATE_FORMAT(CONVERT_TZ(NOW(), @@session.time_zone, '+09:00'),'%Y-%m-%d'),TIME_FORMAT(CONVERT_TZ(NOW(), @@session.time_zone, '+09:00'),'%H:%i:%s'),?)`)

	sqlDetail := `insert into detail_commuting_trip(id_commuting_trip,
									type_of_Transport ,purpose ,detail_from ,detail_to,distance,
									cost,point_trip,transit_point,commute_distance,go_out_distance)
									VALUES`

	for _, initializeDataD := range initializeData.DataDetail {
		sqlDetail += "(?,?,?,?,?,?,?,?,?,?,?),"
		vals = append(vals, initializeDataD.IdCommutingTrip, initializeDataD.TypeOfTransport, initializeDataD.Purpose, initializeDataD.DetailFrom, initializeDataD.DetailTo, initializeDataD.Distance, initializeDataD.Cost, initializeDataD.PointTrip, initializeDataD.TransitPoint, initializeDataD.CommuteDistance, initializeDataD.GoOutDistance)
	}
	sqlDetail = sqlDetail[0 : len(sqlDetail)-1]
	stmtDetail, _ := model.DB.Prepare(sqlDetail)

	if err != nil {
		panic(err.Error())
	}

	defer model.DB.Close()

	valid, message := utils_enter_the_information.ValidatorInsertUsageRecordApplyForTravelExpenses(initializeData)
	if valid == false {
		return nil, message
	}

	execute, err1 := rows.Exec(initializeData.IdGeneralInformation, initializeData.RouteProfileName, initializeData.Date, initializeData.Attendance, initializeData.CodeCommuting, con)
	res, _ := stmtDetail.Exec(vals...)
	if res == nil {
		log.Println("gagal")
	}
	if err1 != nil && execute == nil {
		log.Println(err1)
		return nil, "Missing required field in body request"
	}

	datainsert := enter_the_information.InsertUsageRecordApplyForTravelExpenses{
		RouteProfileName:     initializeData.RouteProfileName,
		Date:                 initializeData.Date,
		Attendance:           initializeData.Attendance,
		CodeCommuting:        initializeData.CodeCommuting,
		IdGeneralInformation: initializeData.IdGeneralInformation,
		DataDetail:           initializeData.DataDetail,
	}

	it = append(it, datainsert)

	return it, "Success Response"

}

// indonesia
// update commuting_trip dan detail_commuting_trip
// body row -> json

// english
// update commuting_trip by id_commuting_trip and detail_commuting_trip by id_commuting_trip_detail
// body row -> json
func (model Models_init_Usage_Record) Model_UpdateUsageRecordApplyForTravelExpenses(initializeData *enter_the_information.UpdateUsageRecordApplyForTravelExpenses) (it []enter_the_information.UpdateUsageRecordApplyForTravelExpenses, condition string) {
	vals := []interface{}{}

	rows, err := model.DB.Prepare(`update commuting_trip set id_general_information = ?,
	route_profile_name = ?,date =?,attendance_code = ?,code_commuting= ? where id_commuting_trip = ?`)

	sqlDetail := `insert into detail_commuting_trip(id_detail_commuting_trip,id_commuting_trip,
									type_of_Transport ,purpose ,detail_from ,detail_to,distance,
									cost,point_trip,transit_point,commute_distance,go_out_distance)
									VALUES`

	for _, initializeDataD := range initializeData.DataDetail {
		sqlDetail += "(?,?,?,?,?,?,?,?,?,?,?,?),"
		vals = append(vals, initializeDataD.IdCommutingTripDetail, initializeDataD.IdCommutingTrip, initializeDataD.TypeOfTransport, initializeDataD.Purpose, initializeDataD.DetailFrom, initializeDataD.DetailTo, initializeDataD.Distance, initializeDataD.Cost, initializeDataD.PointTrip, initializeDataD.TransitPoint, initializeDataD.CommuteDistance, initializeDataD.GoOutDistance)
	}
	sqlDetail = sqlDetail[0 : len(sqlDetail)-1]
	sqlDetail += `ON DUPLICATE KEY UPDATE id_detail_commuting_trip=VALUES(id_detail_commuting_trip),id_commuting_trip=VALUES(id_commuting_trip),type_of_transport=VALUES(type_of_transport), purpose=VALUES(purpose), detail_from=VALUES(detail_from), detail_to=VALUES(detail_to),distance=VALUES(distance),cost=VALUES(cost),point_trip=VALUES(point_trip),transit_point=VALUES(transit_point),commute_distance=VALUES(commute_distance),go_out_distance=VALUES(go_out_distance)`
	stmtDetail, _ := model.DB.Prepare(sqlDetail)

	if err != nil {
		panic(err.Error())
	}

	defer model.DB.Close()

	//valid, message := utils_enter_the_information.ValidatorInsertUsageRecordApplyForTravelExpenses(initializeData)
	//if valid == false {
	//	return nil, message
	//}

	execute, err1 := rows.Exec(initializeData.IdGeneralInformation, initializeData.RouteProfileName, initializeData.Date, initializeData.Attendance, initializeData.CodeCommuting, initializeData.IdCommutingTrip)
	res, _ := stmtDetail.Exec(vals...)
	if res == nil {
		log.Println("gagal")
	}
	if err1 != nil && execute == nil {
		log.Println(err1)
		return nil, "Missing required field in body request"
	}

	dataShow := enter_the_information.UpdateUsageRecordApplyForTravelExpenses{
		IdCommutingTrip:      initializeData.IdCommutingTrip,
		RouteProfileName:     initializeData.RouteProfileName,
		Date:                 initializeData.Date,
		Attendance:           initializeData.Attendance,
		CodeCommuting:        initializeData.CodeCommuting,
		IdGeneralInformation: initializeData.IdGeneralInformation,
		DataDetail:           initializeData.DataDetail,
	}

	it = append(it, dataShow)

	return it, "Success Response"

}

// indonesia
// Hapus Semua data commuting_trip and detail_commuting_trip berdasar kan id_commuting_trip = id_commuting_trip ( table detail)

// english
// delete all data commuting_trip and detail_commuting_trip by id_commuting_trip = id_commuting_trip in table detail

func (model Models_init_Usage_Record) Model_DeleteUsageRecordApplyForTravelExpenses(id string) (response int64, condition string) {

	rows, err := model.DB.Prepare(`DELETE commuting_trip, detail_commuting_trip FROM commuting_trip INNER JOIN detail_commuting_trip 
WHERE commuting_trip.id_commuting_trip = detail_commuting_trip.id_commuting_trip
and commuting_trip.id_commuting_trip =?`)

	if err != nil {
		log.Println(err)
	}

	execute, errExecute := rows.Exec(id)

	if errExecute != nil && execute == nil {
		return 0, "Wrong ID"
	}
	rowsAffected, errRowAffected := execute.RowsAffected()

	if errRowAffected != nil {
		log.Println(errRowAffected)
	}

	return rowsAffected, "Success Response"
}

// indonesia
// Usage Record -> update data menjadi draft [commuting_trip]

// english
// Usage Record -> update data be draft [commuting_trip]

func (model Models_init_Usage_Record) Model_UpdateUsageRecordDraft(id string) (response int64, condition string) {

	sqlUpdate := `update commuting_trip set draft = 'Y' where id_commuting_trip IN(`+id+`)`

	stmtUpdate, errStmtUpdate := model.DB.Query(sqlUpdate)

	if errStmtUpdate != nil {
		return 0, "Please Check Your ID"
	}
	defer stmtUpdate.Close()
	

	return 1, "Success Response"

}
