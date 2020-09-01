package transportation_application

import (
	"../../../initialize/enter_the_information"
	"../../../models"
	utils_enter_the_information "../../../utils/enter_the_information"
	"log"
	"math/rand"
	"strconv"
)

type Models_init_Usage_Record models.DB_init

// indonesia
// Menampilkan semua detail_commuting_trip berdasarkan code_store dan employee_store di group by berdasarkan id_commuting_trip
// data di looping

// english
// Show all detail_commuting_trip based on code_store and employee_store in group by based on id_commuting_trip
// data is looped
func (model Models_init_Usage_Record) Model_GetByIdUsageRecord(store_number string, employee_number string) (sh []enter_the_information.FormatShowUsageRecord, err error) {

	CountHistory := utils_enter_the_information.CheckDataByStoreAndEmployee(`SELECT COUNT(*) FROM (SELECT COUNT(detcomtrip.id_commuting_trip) FROM commuting_trip ct INNER JOIN detail_commuting_trip detcomtrip ON ct.id_commuting_trip = detcomtrip.id_commuting_trip inner join general_information gi on gi.id_general_information = ct.id_general_information INNER JOIN basic_information bi ON bi.id_basic_information = gi.id_basic_information inner join store_information si on si.id_code_store = gi.id_store_code where ct.save_trip ='N' and ct.submit ='Y' and si.code_store =? and bi.employee_code =? group by detcomtrip.id_commuting_trip) t`, store_number, employee_number)

	GetBasicInformation, errGetBasicInformation := model.DB.Query(`select bi.id_basic_information,bi.first_name, bi.last_name, bi.adress, bi.adress_kana,
										bi.adress_detail,bi.adress_detail_kana, bi.add_phone_number
										from basic_information bi,store_information si , general_information gi where
										gi.id_basic_information = bi.id_basic_information and 
 									   	gi.id_store_code = si.id_code_store and si.code_store =? and 
 									   	bi.employee_code=?`, store_number, employee_number)

	rows, err := model.DB.Query(`select  b.id_commuting_trip,COALESCE(SUM(b.distance),0)
 										as distance,COALESCE(SUM(commute_distance),0) as commute_distance, COALESCE(SUM(b.cost),0) as cost , ct.draft,b.purpose
 										 from basic_information bi, commuting_trip ct, detail_commuting_trip b, store_information si , general_information gi, 
										master_transportation trans
										where ct.id_commuting_trip = b.id_commuting_trip and gi.id_basic_information = bi.id_basic_information
										and b.type_of_transport =  trans.code_transportation
										and gi.id_store_code = si.id_code_store and ct.id_general_information = gi.id_general_information and si.code_store =? and bi.employee_code=?
										and ct.submit ='N' and ct.save_trip ='N'
										group by b.id_commuting_trip`, store_number, employee_number)
	var init_biC interface{}
	var init_bi enter_the_information.ShowBasicInformation1
	//var Arr_bi []enter_the_information.ShowBasicInformation1
	var init_ur enter_the_information.ShowUsageRecord2
	var Arr_ur []enter_the_information.ShowUsageRecord2
	if err != nil && errGetBasicInformation != nil {
		log.Println(err.Error())
		log.Println(errGetBasicInformation.Error())
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

	for rows.Next() {
		err := rows.Scan(&init_ur.IdCommutingTrip, &init_ur.Distance, &init_ur.CommuteDistance, &init_ur.Cost, &StatusTemporari, &init_ur.Purpose)
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

			dataCommutingTrip := enter_the_information.ShowUsageRecord2{
				IdCommutingTrip: init_ur.IdCommutingTrip,
				TypeOfTransport: DatatypeOfTransportation,
				Purpose:         DataPurpose,
				Route:           DataRoute,
				Distance:        init_ur.Distance,
				CommuteDistance: init_ur.CommuteDistance,
				Cost:            init_ur.Cost,
				StatusTemporary: StatusTemporari,
			}
			Arr_ur = append(Arr_ur, dataCommutingTrip)
		}
	}
	if init_biC != nil && Arr_ur != nil {
		FinallyData := enter_the_information.FormatShowUsageRecord{
			CountHistory:         CountHistory,
			KodeBasicInformation: KodeBasicInformation,
			DataBasicInformation: init_biC,
			DataUsageRecord:      Arr_ur,
		}
		sh = append(sh, FinallyData)
		return sh, nil
	}
	return nil, nil
}

// indonesia
// Menampilkan data Usage Record untuk di edit berdasarkan id commuting trip ,store number dan employee number

// english
// Show Data Usage Record to edit by id commuting trip, store number dan employee number
func (model Models_init_Usage_Record) Model_GetByIdUsageRecordForEdit(store_number string, employee_number string, id_commuting_trip string) (sh []enter_the_information.FormatShowUsageRecordForEdit, err error) {
	var inter_CommutingTrip interface{}
	var shCommutingTrip enter_the_information.ShowCommutingTrip
	//var Arr_shCommutingTrip []enter_the_information.ShowCommutingTrip
	var Arr_shCommutingTripDetail []enter_the_information.ShowUsageRecord
	QueryShowCommutingTrip, errShowCommutingTrip := model.DB.Query(`select ct.id_commuting_trip, ct.route_profile_name,
 										ct.date, ct.attendance_code from commuting_trip ct where ct.id_commuting_trip = ?`, id_commuting_trip)

	if errShowCommutingTrip != nil {
		log.Println(errShowCommutingTrip)
	}

	QueryShowCommutingTrip.Next()
	errShowCommutingTripScan := QueryShowCommutingTrip.Scan(&shCommutingTrip.IdCommutingTrip, &shCommutingTrip.RouteProfileName, &shCommutingTrip.Date, &shCommutingTrip.AttendanceCode)

	if errShowCommutingTripScan != nil {
		inter_CommutingTrip = nil
	} else {
		inter_CommutingTrip = shCommutingTrip
	}

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
			return nil, nil
		}

		Arr_shCommutingTripDetail = append(Arr_shCommutingTripDetail, init_container)

	}

	FinnalyData := enter_the_information.FormatShowUsageRecordForEdit{
		DataTrip:       inter_CommutingTrip,
		DetailDataTrip: Arr_shCommutingTripDetail,
	}
	sh = append(sh, FinnalyData)

	return sh, nil
}

// indonesia
// Menampilkan Semua route favorit berdasarkan store number dan employee number

// english
// get all data route favorite by store number and employee number

func (model Models_init_Usage_Record) Model_GetByIdUsageRecordUseMyRoute(store_number string, employee_number string) (sh []enter_the_information.ShowUseMyRoute, err error) {

	rows, err := model.DB.Query(`select MIN(detcomtrip.id_commuting_trip), MIN(detcomtrip.id_detail_commuting_trip), MIN(comtrip.route_profile_name),  MIN(comtrip.attendance_code),
MIN(detcomtrip.purpose), COALESCE(SUM(detcomtrip.distance),0) ,COALESCE(SUM(detcomtrip.commute_distance),0), COALESCE(SUM(detcomtrip.cost),0)  from commuting_trip comtrip, detail_commuting_trip detcomtrip, general_information geninfo, basic_information bainfo, store_information storeinfo
where comtrip.id_commuting_trip = detcomtrip.id_commuting_trip and geninfo.id_general_information = comtrip.id_general_information AND
geninfo.id_basic_information = bainfo.id_basic_information and geninfo.id_store_code = storeinfo.id_code_store  and storeinfo.code_store =? and bainfo.employee_code =? and comtrip.save_trip ='Y'
group by comtrip.id_commuting_trip ORDER BY MIN(comtrip.date) asc`, store_number, employee_number)

	var init_container enter_the_information.ShowUseMyRoute
	if err != nil {
		log.Println(err.Error())
	}

	for rows.Next() {
		err := rows.Scan(&init_container.IdCommutingTrip, &init_container.IdDetailCommutingTrip, &init_container.RouteProfileName, &init_container.AttendanceCode, &init_container.Purpose, &init_container.Distance, &init_container.CommuteDistance, &init_container.Cost)

		if err != nil {
			panic(err.Error())
		}

		DatatypeOfTransportation, DataRoute, _ :=	utils_enter_the_information.GetAdditionalUsageRecord(store_number,employee_number,init_container.IdCommutingTrip,"usageRecordUseRoute")
		FinallyData := enter_the_information.ShowUseMyRoute{
			IdCommutingTrip:       init_container.IdCommutingTrip,
			IdDetailCommutingTrip: init_container.IdDetailCommutingTrip,
			RouteProfileName:      init_container.RouteProfileName,
			TypeOfTransport:       DatatypeOfTransportation,
			AttendanceCode:        init_container.AttendanceCode,
			Purpose:               init_container.Purpose,
			Route:                 DataRoute,
			Distance:              init_container.Distance,
			CommuteDistance:       init_container.CommuteDistance,
			Cost:                  init_container.Cost,
		}
		sh = append(sh, FinallyData)
		//sh = append(sh, init_container)

	}

	return sh, nil
}

func (model Models_init_Usage_Record) Model_GetByIdUsageRecordHistory(store_number string,
	employee_number string, page string, filter string, showData string, searching string) (sh []enter_the_information.FormatHistory, err error) {

	var Arr_History []enter_the_information.ShowHistory
	var pageInt int
	var showDataInt int
	var limitPage string
	if page != "" {
		parsePage, _ := strconv.Atoi(page)
		pageInt = parsePage
	}

	if showData != "" {
		parseShowData, _ := strconv.Atoi(showData)
		showDataInt = parseShowData
	}

	if page == "" && showData == "" {
		limitPage = ""
	} else {
		limitPageInt := (pageInt - 1) * showDataInt
		DataPageInt := strconv.Itoa(limitPageInt)
		DataShowDataInt := strconv.Itoa(showDataInt)
		limitPage = `LIMIT ` + DataPageInt + `,` + DataShowDataInt
	}
	filterMonth := ``
	if filter == "" {
		filterMonth = ``
	} else {
		filterMonth = ` and MONTH(comtrip.date) =` + filter
	}
	searchingAction := ``
	if searching == "" {
		searchingAction = ``
	} else {
		searchingAction = ` and (comtrip.date LIKE '% ` + searching + `%' OR comtrip.route_profile_name LIKE '%` + searching + `%' OR detcomtrip.purpose LIKE '%` + searching + `%'OR comtrip.attendance_code LIKE '%` + searching + `%')`
	}

	rows, err := model.DB.Query(`select  MIN(comtrip.id_commuting_trip), MIN(detcomtrip.id_detail_commuting_trip), comtrip.date, MIN(comtrip.route_profile_name), MIN(comtrip.attendance_code), 
										MIN(detcomtrip.purpose), COALESCE(SUM(detcomtrip.distance),0), COALESCE(SUM(detcomtrip.commute_distance),0) , COALESCE(SUM(detcomtrip.cost),0), MIN(cc.status_commuting), CAST(comtrip.date_time_approve as DATE) as date_time_approve
										from commuting_trip comtrip, code_commuting cc,
										detail_commuting_trip detcomtrip, general_information geninfo, basic_information bainfo, store_information storeinfo
										where comtrip.id_commuting_trip = detcomtrip.id_commuting_trip and geninfo.id_general_information = comtrip.id_general_information AND
										geninfo.id_basic_information = bainfo.id_basic_information and geninfo.id_store_code = storeinfo.id_code_store  and storeinfo.code_store =? and cc.code_random = comtrip.code_commuting
										and bainfo.employee_code =? and comtrip.save_trip ='N' and comtrip.submit = 'Y' `+filterMonth+searchingAction+`
										group by detcomtrip.id_commuting_trip order by comtrip.date asc `+limitPage, store_number, employee_number)

	var init_container enter_the_information.ShowHistory
	if err != nil {
		log.Println(err.Error())
	}

	for rows.Next() {
		err := rows.Scan(&init_container.IdCommutingTrip, &init_container.IdDetailCommutingTrip, &init_container.Date, &init_container.RouteProfileName, &init_container.AttendanceCode, &init_container.Purpose, &init_container.Distance, &init_container.CommuteDistance, &init_container.Cost, &init_container.StatusCommuting, &init_container.DateApprove)
		if err != nil {
			panic(err.Error())
		}
		DatatypeOfTransportation, DataRoute,_ := utils_enter_the_information.GetAdditionalUsageRecord(store_number,employee_number,init_container.IdCommutingTrip,"usageRecordHistory")
		// Get Data Transportation, detail from, detail to and purpose (horizontal)
		//GetDataTypeOfTransportationAndRoute, errGetDataTypeOfTransportationAndRoute := model.DB.Query(`
		//select MIN(detcomtrip.id_commuting_trip) as id, MIN(comtrip.route_profile_name), MIN(detcomtrip.type_of_transport),
		//MIN(comtrip.attendance_code) from commuting_trip comtrip, detail_commuting_trip detcomtrip, general_information geninfo,
		//basic_information bainfo, store_information storeinfo where comtrip.id_commuting_trip = detcomtrip.id_commuting_trip and
		//geninfo.id_general_information = comtrip.id_general_information AND geninfo.id_basic_information = bainfo.id_basic_information and
		//geninfo.id_store_code = storeinfo.id_code_store and storeinfo.code_store =? and bainfo.employee_code =? and comtrip.submit ='Y'
		//and comtrip.save_trip ='N' and detcomtrip.id_commuting_trip =? group by detcomtrip.id_commuting_trip order by comtrip.date asc
		//`, store_number, employee_number, init_container.IdCommutingTrip)
		//
		//if errGetDataTypeOfTransportationAndRoute != nil {
		//	log.Println(0)
		//	typeOfTransportation = ""
		//	DetailTo = ""
		//	DetailFrom = ""
		//	Purpose = ""
		//} else {
		//	for GetDataTypeOfTransportationAndRoute.Next() {
		//		errGetDataT := GetDataTypeOfTransportationAndRoute.Scan(&typeOfTransportation, &DetailFrom, &DetailTo, &Purpose)
		//
		//		if errGetDataT != nil {
		//			log.Println(errGetDataT.Error())
		//		}
		//		DatatypeOfTransportation += typeOfTransportation + ` - `
		//		DataRoute += DetailFrom + ` - - ` + DetailTo + `-`
		//		DataPurpose += Purpose + ` - `
		//
		//	}
		//	if typeOfTransportation != "" {
		//		DatatypeOfTransportation = DatatypeOfTransportation[0 : len(DatatypeOfTransportation)-3]
		//	}
		//	if DataRoute != "" {
		//		DataRoute = DataRoute[0 : len(DataRoute)-1]
		//	}
		//	if DataPurpose != "" {
		//		DataPurpose = DataPurpose[0 : len(DataPurpose)-3]
		//	}
			//log.Println(DatatypeOfTransportation)
			//log.Println(DataRoute)
			//log.Println(DataPurpose)
		//}
		// end Get Data Transportation, detail from, detail to and purpose (horizontal)
		FinnalyData := enter_the_information.ShowHistory{
			IdDetailCommutingTrip: init_container.IdDetailCommutingTrip,
			IdCommutingTrip:       init_container.IdCommutingTrip,
			RouteProfileName:      init_container.RouteProfileName,
			Date:                  init_container.Date,
			TypeOfTransport:       DatatypeOfTransportation,
			AttendanceCode:        init_container.AttendanceCode,
			Purpose:               init_container.Purpose,
			Distance:              init_container.Distance,
			CommuteDistance:       init_container.CommuteDistance,
			Cost:                  init_container.Cost,
			Route:                 DataRoute,
			StatusCommuting:       init_container.StatusCommuting,
			DateApprove:           init_container.DateApprove,
		}
		Arr_History = append(Arr_History, FinnalyData)

	}

	DataSubmit := utils_enter_the_information.CheckDataByStoreAndEmployee(`select COUNT(*) from (select COUNT(*)
										from commuting_trip comtrip, code_commuting cc,
										detail_commuting_trip detcomtrip, general_information geninfo, basic_information bainfo, store_information storeinfo
										where comtrip.id_commuting_trip = detcomtrip.id_commuting_trip and geninfo.id_general_information = comtrip.id_general_information AND
										geninfo.id_basic_information = bainfo.id_basic_information and geninfo.id_store_code = storeinfo.id_code_store  and storeinfo.code_store =? and cc.code_random = comtrip.code_commuting
										and bainfo.employee_code =? and comtrip.save_trip ='N' and comtrip.submit = 'Y' and cc.status_commuting ='submit'
										group by detcomtrip.id_commuting_trip order by comtrip.date asc) t`, store_number, employee_number)
	DataDraft := utils_enter_the_information.CheckDataByStoreAndEmployee(`select COUNT(*) from (select COUNT(*)
										from commuting_trip comtrip, code_commuting cc,
										detail_commuting_trip detcomtrip, general_information geninfo, basic_information bainfo, store_information storeinfo
										where comtrip.id_commuting_trip = detcomtrip.id_commuting_trip and geninfo.id_general_information = comtrip.id_general_information AND
										geninfo.id_basic_information = bainfo.id_basic_information and geninfo.id_store_code = storeinfo.id_code_store  and storeinfo.code_store =? and cc.code_random = comtrip.code_commuting
										and bainfo.employee_code =? and comtrip.save_trip ='N' and comtrip.submit = 'Y' and cc.status_commuting ='draft'
										group by detcomtrip.id_commuting_trip order by comtrip.date asc) t`, store_number, employee_number)
	DataPartial := utils_enter_the_information.CheckDataByStoreAndEmployee(`select COUNT(*) from (select COUNT(*)
										from commuting_trip comtrip, code_commuting cc,
										detail_commuting_trip detcomtrip, general_information geninfo, basic_information bainfo, store_information storeinfo
										where comtrip.id_commuting_trip = detcomtrip.id_commuting_trip and geninfo.id_general_information = comtrip.id_general_information AND
										geninfo.id_basic_information = bainfo.id_basic_information and geninfo.id_store_code = storeinfo.id_code_store  and storeinfo.code_store =? and cc.code_random = comtrip.code_commuting
										and bainfo.employee_code =? and comtrip.save_trip ='N' and comtrip.submit = 'Y' and cc.status_commuting ='partial'
										group by detcomtrip.id_commuting_trip order by comtrip.date asc) t`, store_number, employee_number)
	DataNotApproved := utils_enter_the_information.CheckDataByStoreAndEmployee(`select COUNT(*) from (select COUNT(*)
										from commuting_trip comtrip, code_commuting cc,
										detail_commuting_trip detcomtrip, general_information geninfo, basic_information bainfo, store_information storeinfo
										where comtrip.id_commuting_trip = detcomtrip.id_commuting_trip and geninfo.id_general_information = comtrip.id_general_information AND
										geninfo.id_basic_information = bainfo.id_basic_information and geninfo.id_store_code = storeinfo.id_code_store  and storeinfo.code_store =? and cc.code_random = comtrip.code_commuting
										and bainfo.employee_code =? and comtrip.save_trip ='N' and comtrip.submit = 'Y' and cc.status_commuting ='not_approved'
										group by detcomtrip.id_commuting_trip order by comtrip.date asc) t`, store_number, employee_number)
	DataCountStatusHistory := enter_the_information.ShowAdditionalHistory{
		CountDataSubmit:      DataSubmit,
		CountDataDraft:       DataDraft,
		CountDataPartial:     DataPartial,
		CountDataNotApproved: DataNotApproved,
	}

	FinallyData := enter_the_information.FormatHistory{
		DataCount:   DataCountStatusHistory,
		Datahistory: Arr_History,
	}
	sh = append(sh, FinallyData)
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
	Status_Draft := "N"
	RandomInte := rand.Intn(999999)
	var RandomInt int

	checkIntRandom := utils_enter_the_information.CheckDataByIdInt(`select COUNT(*) from code_commuting where code_random = ?`, RandomInte)

	if checkIntRandom == 0 {
		RandomInt = RandomInte
	} else {
		for {
			RandomInteg := rand.Intn(999999)
			checkIntRandom := utils_enter_the_information.CheckDataByIdInt(`select COUNT(*) from code_commuting where code_random = ?`, RandomInteg)
			if checkIntRandom == 0 {
				RandomInt = RandomInteg
				break
			}
		}
	}

	if con != "Y" && con != "N" {
		return nil, "Missing required, Please use /Y or /N"
	}
	//else {
	//	Status_Draft = "Y"
	//}
	checkCountData := utils_enter_the_information.CheckDataByStoreAndEmployee(`select COUNT(*) from commuting_trip comtrip, detail_commuting_trip detcomtrip, general_information geninfo, basic_information bainfo, store_information storeinfo where comtrip.id_commuting_trip = detcomtrip.id_commuting_trip and geninfo.id_general_information = comtrip.id_general_information AND geninfo.id_basic_information = bainfo.id_basic_information and geninfo.id_store_code = storeinfo.id_code_store and storeinfo.code_store =? and bainfo.employee_code =? and comtrip.save_trip ='Y'`, store_id, employee_id)

	if con == "Y" && checkCountData >= 3 {
		return nil, "You cannot register up to more than 3 routes"
	}
	vals := []interface{}{}
	//var Arr_DetailDataInsert [] enter_the_information.InsertDetailUsageRecordApplyForTravelExpenses

	insertCommutingTrip, errInsertCommutingTrip := model.DB.Prepare(`insert into commuting_trip(id_general_information,route_profile_name,date,attendance_code,code_commuting,created_date,created_time,save_trip,draft)
 		VALUES(?,?,?,?,?,DATE_FORMAT(CONVERT_TZ(NOW(), @@session.time_zone, '+09:00'),'%Y-%m-%d'),TIME_FORMAT(CONVERT_TZ(NOW(), @@session.time_zone, '+09:00'),'%H:%i:%s'),?,?)`)
	insertCodeCommuting, errInsertCodeCommuting := model.DB.Prepare(`insert into code_commuting(code_random,std_deviation,created_time,created_date,status_commuting)
	VALUES(?,'0',TIME_FORMAT(CONVERT_TZ(NOW(), @@session.time_zone, '+09:00'),'%H:%i:%s'),DATE_FORMAT(CONVERT_TZ(NOW(), @@session.time_zone, '+09:00'),'%Y-%m-%d'),'not_approved')`)
	sqlDetail := `insert into detail_commuting_trip(id_commuting_trip,
									type_of_Transport ,purpose ,detail_from ,detail_to,distance,
									cost,point_trip,transit_point,commute_distance,go_out_distance)
									VALUES`
	//var initializeDataD enter_the_information.InsertDetailUsageRecordApplyForTravelExpenses
	for _, initializeDataD := range initializeData.DataDetail {
		sqlDetail += "(?,?,?,?,?,?,?,?,?,?,?),"
		//Distance := initializeDataD.CommuteDistance + initializeDataD.GoOutDistance
		//initializeDataD.Distance = Distance
		//log.Println(initializeDataD.Distance)
		vals = append(vals, initializeDataD.IdCommutingTrip, initializeDataD.TypeOfTransport, initializeDataD.Purpose, initializeDataD.DetailFrom, initializeDataD.DetailTo, initializeDataD.Distance, initializeDataD.Cost, initializeDataD.PointTrip, initializeDataD.TransitPoint, initializeDataD.CommuteDistance, initializeDataD.GoOutDistance)
	}
	sqlDetail = sqlDetail[0 : len(sqlDetail)-1]
	stmtDetail, _ := model.DB.Prepare(sqlDetail)

	if errInsertCommutingTrip != nil {
		log.Println(errInsertCommutingTrip.Error())
	}

	if errInsertCodeCommuting != nil {
		log.Println(errInsertCodeCommuting.Error())
	}

	defer model.DB.Close()

	valid, message := utils_enter_the_information.ValidatorInsertUsageRecordApplyForTravelExpenses(initializeData)
	if valid == false {
		return nil, message
	}

	executeCodeCommuting, _ := insertCodeCommuting.Exec(RandomInt)
	execute, err1 := insertCommutingTrip.Exec(initializeData.IdGeneralInformation, initializeData.RouteProfileName, initializeData.Date, initializeData.Attendance, RandomInt, con, Status_Draft)
	executeDetailCommutingTrip, _ := stmtDetail.Exec(vals...)
	if executeDetailCommutingTrip == nil {
		log.Println("gagal insert Detail Commuting Trip")
	}
	if executeCodeCommuting == nil {
		log.Println("gagal insert Code Commuting")
	}
	if err1 != nil && execute == nil {
		log.Println(err1)
		return nil, "Missing required field in body request"
	}
	//DetailDataInsert := enter_the_information.InsertDetailUsageRecordApplyForTravelExpenses{
	//	IdCommutingTrip: initializeDataD.IdCommutingTrip,
	//	TypeOfTransport: initializeDataD.TypeOfTransport,
	//	Purpose:         initializeDataD.Purpose,
	//	DetailFrom:      initializeDataD.DetailFrom,
	//	DetailTo:        initializeDataD.DetailTo,
	//	Distance:        initializeDataD.Distance,
	//	Cost:            initializeDataD.Cost,
	//	PointTrip:       initializeDataD.PointTrip,
	//	TransitPoint:    initializeDataD.TransitPoint,
	//	CommuteDistance: initializeDataD.CommuteDistance,
	//	GoOutDistance:   initializeDataD.GoOutDistance,
	//}
	//Arr_DetailDataInsert = append(Arr_DetailDataInsert,DetailDataInsert)

	datainsert := enter_the_information.InsertUsageRecordApplyForTravelExpenses{
		RouteProfileName:     initializeData.RouteProfileName,
		Date:                 initializeData.Date,
		Attendance:           initializeData.Attendance,
		CodeCommuting:        RandomInt,
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
	route_profile_name = ?,date =?,attendance_code = ? where id_commuting_trip = ?`)

	sqlDeleteDataDetail := `DELETE FROM detail_commuting_trip where id_commuting_trip =` + initializeData.IdCommutingTrip
	log.Println(sqlDeleteDataDetail)
	stmtDeleteDataDetail, errstmtDeteleDataDetail := model.DB.Query(sqlDeleteDataDetail)
	if errstmtDeteleDataDetail != nil {
		log.Print("delete?")
		log.Println(errstmtDeteleDataDetail)
	}
	defer stmtDeleteDataDetail.Close()

	sqlDetail := `insert into detail_commuting_trip(id_detail_commuting_trip,id_commuting_trip,
									type_of_Transport ,purpose ,detail_from ,detail_to,distance,
									cost,point_trip,transit_point,commute_distance,go_out_distance)
									VALUES`
	DataIdCommutingTripDetail := ""
	DataCheckIdCommutingTripDetail := ""
	for _, initializeDataD := range initializeData.DataDetail {
		sqlDetail += "(?,?,?,?,?,?,?,?,?,?,?,?),"
		DataCheckIdCommutingTripDetail = strconv.Itoa(initializeDataD.IdCommutingTripDetail)
		DataIdCommutingTripDetail += strconv.Itoa(initializeDataD.IdCommutingTripDetail) + `,`
		vals = append(vals, initializeDataD.IdCommutingTripDetail, initializeDataD.IdCommutingTrip, initializeDataD.TypeOfTransport, initializeDataD.Purpose, initializeDataD.DetailFrom, initializeDataD.DetailTo, initializeDataD.Distance, initializeDataD.Cost, initializeDataD.PointTrip, initializeDataD.TransitPoint, initializeDataD.CommuteDistance, initializeDataD.GoOutDistance)
	}
	if DataCheckIdCommutingTripDetail != "" {
		DataIdCommutingTripDetail = DataIdCommutingTripDetail[0 : len(DataIdCommutingTripDetail)-1]
	} else {
		DataIdCommutingTripDetail = ""
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

	execute, err1 := rows.Exec(initializeData.IdGeneralInformation, initializeData.RouteProfileName, initializeData.Date, initializeData.Attendance, initializeData.IdCommutingTrip)
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

	sqlDelete := `DELETE commuting_trip, detail_commuting_trip FROM commuting_trip
					INNER JOIN detail_commuting_trip 
					WHERE commuting_trip.id_commuting_trip = detail_commuting_trip.id_commuting_trip
					and commuting_trip.id_commuting_trip IN(` + id + `)`

	QueryDelete, errQueryDelete := model.DB.Query(sqlDelete)

	if errQueryDelete != nil {
		return 0, "Please Check Your ID"
	}
	defer QueryDelete.Close()

	return 1, "Success Response"
}

// indonesia
// Usage Record -> update data menjadi draft [commuting_trip]

// english
// Usage Record -> update data be draft [commuting_trip]

func (model Models_init_Usage_Record) Model_UpdateUsageRecordDraft(id string) (response int64, condition string) {

	sqlUpdate := `update commuting_trip set draft = 'Y' where id_commuting_trip IN(` + id + `)`

	stmtUpdate, errStmtUpdate := model.DB.Query(sqlUpdate)

	if errStmtUpdate != nil {
		return 0, "Please Check Your ID"
	}
	defer stmtUpdate.Close()

	return 1, "Success Response"
}

func (model Models_init_Usage_Record) Model_UseUsageRecord(id string) (response int, condition string) {

	sqlUseCommutingTrip := `insert into commuting_trip (id_general_information,
	route_profile_name, date, attendance_code, created_date,created_time) 
	select a.id_general_information, a.route_profile_name, a.date, a.attendance_code, 
	DATE_FORMAT(CONVERT_TZ(NOW(), @@session.time_zone, '+09:00'),'%Y-%m-%d'),
	TIME_FORMAT(CONVERT_TZ(NOW(), @@session.time_zone, '+09:00'),'%H:%i:%s')
	from commuting_trip a where a.id_commuting_trip IN(` + id + `)`

	sqlUseDetailCommutingTrip := `INSERT INTO detail_commuting_trip( id_commuting_trip, type_of_transport, 
	purpose, detail_from,detail_to, distance, cost, point_trip, transit_point, commute_distance,
	go_out_distance)
	select detcomtrip.id_commuting_trip,detcomtrip.type_of_transport,detcomtrip.purpose,
	detcomtrip.detail_from, detcomtrip.detail_to, detcomtrip.distance, detcomtrip.cost, 
	detcomtrip.point_trip, detcomtrip.transit_point, detcomtrip.commute_distance,
	detcomtrip.go_out_distance from detail_commuting_trip detcomtrip where id_commuting_trip IN(` + id + `)`

	stmtUseCommutingTrip, errstmtCommutingTrip := model.DB.Query(sqlUseCommutingTrip)
	stmtUseDetailCommutingTrip, errstmtDetailCommutingTrip := model.DB.Query(sqlUseDetailCommutingTrip)

	if errstmtCommutingTrip != nil && errstmtDetailCommutingTrip != nil {
		return 0, "Please Check Your ID"
	}
	defer stmtUseCommutingTrip.Close()
	defer stmtUseDetailCommutingTrip.Close()

	return 1, "Success Response"
}
