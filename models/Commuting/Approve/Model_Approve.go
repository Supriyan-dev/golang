package Approve

import (
	approve "../../../initialize/Commuting"
	"../../../models"
	utils_Global "../../../utils"
	utils_enter_the_information "../../../utils/enter_the_information"
	"context"
	"log"
	"strconv"
)

type Init_DB_CommutingApprove models.DB_init

// View data Commuting Agregat (SUM) By All Employee Code
func (model Init_DB_CommutingApprove) GetDataApproveCommutingSumByAllEmployeeCode(page string,
	filter string, showData string, searching string, condition string, store_code string, department_code string) (sh []approve.Init_CommutingApprove, err error, CountData int) {

	var Da approve.Init_CommutingApprove

	var pageInt int
	var showDataInt int
	var limitPage string

	queryManagerApprove := ""

	if store_code == "" || department_code == "" {
		queryManagerApprove = ""
	} else {
		queryManagerApprove = ` and store_information.code_store = ` + store_code + ` AND department_information.department_code = ` + department_code
	}

	ConditionString := ""
	if page != "" {
		parsePage, _ := strconv.Atoi(page)
		pageInt = parsePage
	}

	if condition == "Y" {
		ConditionString = ` AND commuting_trip.status_approval != 'N'`
	} else {
		ConditionString = ` AND commuting_trip.status_approval = 'T'`
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
		limitPage = ` LIMIT ` + DataPageInt + `,` + DataShowDataInt
	}
	filterMonth := ``
	if filter == "" {
		filterMonth = ``
	} else {
		filterMonth = ` and MONTH(commuting_trip.date) =` + filter
	}

	searchingAction := ``
	if searching == "" {
		searchingAction = ``
	} else {
		searchingAction = ` and (store_information.code_store LIKE '% ` + searching + `%' OR department_information.department_name LIKE '%` + searching + `%' OR store_section_information.store_section_name LIKE '%` + searching + `%' OR basic_information.first_name LIKE '%` + searching + `%' OR basic_information.last_name LIKE '%` + searching + `%' OR basic_information.adress LIKE '%` + searching + `%' OR basic_information.adress_kana LIKE '%` + searching + `%' OR basic_information.adress_detail LIKE '%` + searching + `%' OR basic_information.adress_detail_kana LIKE '%` + searching + `%' OR basic_information.add_phone_number LIKE '%` + searching + `%')`
	}

	CountDataApprove := model.DB.QueryRow(`SELECT
    COUNT(*)
FROM
    (
    SELECT
        MIN(
            commuting_trip.id_commuting_trip
        ) id_commuting_trip,MIN(basic_information.employee_code),
        MIN(store_information.code_store) AS code_store,
        MIN(
            department_information.department_name
        ),
        CONCAT(
            MIN(
                store_section_information.store_section_name
            ),
            ' ',
            MIN(
                store_section_information.store_section_code
            )
        ) AS store_section,
        MIN(unit_information.unit_code),
        CONCAT(
            MIN(basic_information.first_name),
            ' ',
            MIN(basic_information.last_name),' ',
        CONCAT(
            RIGHT(
                MIN(store_information.code_store),
                4
            ),
            LPAD(
                RIGHT(
                    MIN(
                        department_information.department_code
                    ),
                    2
                ),
                2,
                '0'
            ),
            LPAD(
                RIGHT(
                    MIN(
                        store_section_information.store_section_code
                    ),
                    2
                ),
                2,
                '0'
            ),
            LPAD(
                RIGHT(
                    MIN(unit_information.unit_code),
                    2
                ),
                2,
                '0'
            )
        )
        ),
        CONCAT(
        MIN(basic_information.adress),', ',
        MIN(basic_information.adress_kana),', ',
        MIN(
            basic_information.adress_detail
        ), ', ', MIN(
            basic_information.adress_detail_kana
        ),' ',  MIN(
            basic_information.add_phone_number
        ) 
        ),
        MIN(commuting_trip.date) date_trip,
        MIN(
            basic_information.employee_code
        ) employee_code,
        MIN(
            code_commuting.status_commuting
        ),
 		MIN(
            code_commuting.code_random
        ),
 		MIN(
            basic_information.id_basic_information
        )
    FROM
        commuting_trip
    LEFT OUTER JOIN general_information ON general_information.id_general_information = commuting_trip.id_general_information
    LEFT OUTER JOIN code_commuting ON commuting_trip.code_commuting = code_commuting.code_random
    LEFT OUTER JOIN store_section_information ON store_section_information.id_store_section = general_information.id_store_section
    LEFT OUTER JOIN basic_information ON basic_information.id_basic_information = general_information.id_basic_information
    LEFT OUTER JOIN store_information ON store_information.id_code_store = general_information.id_store_code
    LEFT OUTER JOIN department_information ON department_information.id_department = general_information.id_department
    LEFT OUTER JOIN unit_information ON unit_information.id_unit = general_information.id_unit
    WHERE
  commuting_trip.submit = 'Y' AND commuting_trip.date_submit IS NOT NULL ` + ConditionString + queryManagerApprove + ` AND commuting_trip.save_trip = 'N'
` + filterMonth + searchingAction +`   
GROUP BY
        basic_information.employee_code
) t
ORDER BY
    code_store ASC,
    date_trip
DESC
    ,
    id_commuting_trip
DESC`).Scan(&CountData)
	if CountDataApprove != nil {
		CountData = 0
	}

	showDataApprove, errShowDataApprove := model.DB.Query(
		`
SELECT
    *
FROM
    (
    SELECT
        MIN(
            commuting_trip.id_commuting_trip
        ) id_commuting_trip,MIN(basic_information.employee_code),
        MIN(store_information.code_store) AS code_store,
        MIN(
            department_information.department_name
        ),
        CONCAT(
            MIN(
                store_section_information.store_section_name
            ),
            ' ',
            MIN(
                store_section_information.store_section_code
            )
        ) AS store_section,
        MIN(unit_information.unit_code),
        CONCAT(
            MIN(basic_information.first_name),
            ' ',
            MIN(basic_information.last_name),' ',
        CONCAT(
            RIGHT(
                MIN(store_information.code_store),
                4
            ),
            LPAD(
                RIGHT(
                    MIN(
                        department_information.department_code
                    ),
                    2
                ),
                2,
                '0'
            ),
            LPAD(
                RIGHT(
                    MIN(
                        store_section_information.store_section_code
                    ),
                    2
                ),
                2,
                '0'
            ),
            LPAD(
                RIGHT(
                    MIN(unit_information.unit_code),
                    2
                ),
                2,
                '0'
            )
        )
        ),
        CONCAT(
        MIN(basic_information.adress),', ',
        MIN(basic_information.adress_kana),', ',
        MIN(
            basic_information.adress_detail
        ), ', ', MIN(
            basic_information.adress_detail_kana
        ),' ',  MIN(
            basic_information.add_phone_number
        ) 
        ),
        MIN(commuting_trip.date) date_trip,
        MIN(
            basic_information.employee_code
        ) employee_code,
        MIN(
            code_commuting.status_commuting
        ),
 		MIN(
            code_commuting.code_random
        ),
 		MIN(
            basic_information.id_basic_information
        )
    FROM
        commuting_trip
    LEFT OUTER JOIN general_information ON general_information.id_general_information = commuting_trip.id_general_information
    LEFT OUTER JOIN code_commuting ON commuting_trip.code_commuting = code_commuting.code_random
    LEFT OUTER JOIN store_section_information ON store_section_information.id_store_section = general_information.id_store_section
    LEFT OUTER JOIN basic_information ON basic_information.id_basic_information = general_information.id_basic_information
    LEFT OUTER JOIN store_information ON store_information.id_code_store = general_information.id_store_code
    LEFT OUTER JOIN department_information ON department_information.id_department = general_information.id_department
    LEFT OUTER JOIN unit_information ON unit_information.id_unit = general_information.id_unit
    WHERE
        commuting_trip.submit = 'Y' AND commuting_trip.date_submit IS NOT NULL ` + ConditionString + queryManagerApprove + ` AND commuting_trip.save_trip = 'N'
` + filterMonth + searchingAction +
			`    
GROUP BY
        basic_information.employee_code
) t
ORDER BY
    code_store ASC,
    date_trip
DESC
    ,
    id_commuting_trip
DESC` + limitPage)

	if errShowDataApprove != nil {
		log.Println(errShowDataApprove)
	}

	var Datatrip models.NullString
	var EmployeeNumber models.NullString
	var StatusCommuting models.NullString
	for showDataApprove.Next() {
		ScanData := showDataApprove.Scan(&Da.IdCommuting, &Da.EmployeeNumber, &Da.CodeStore, &Da.DepartmentCode, &Da.StoreSection, &Da.UnitCode, &Da.DivisiCodeAndName, &Da.AddressAndNumber, &Datatrip, &EmployeeNumber, &StatusCommuting, &Da.CodeCommuting, &Da.IdBasicInformation)

		if ScanData != nil {
			log.Println(Datatrip)
			log.Println(ScanData)
		}
		sumCost := utils_Global.GetDataByIdInt(`select COALESCE(SUM(cost),0) from detail_commuting_trip where id_commuting_trip = ? `, Da.IdCommuting)
		sumDistance := utils_Global.GetDataByIdFloat(`select COALESCE(SUM(distance),0.00) from detail_commuting_trip where id_commuting_trip = ? `, Da.IdCommuting)

		InitDataScan := approve.Init_CommutingApprove{
			IdCommuting:        Da.IdCommuting,
			EmployeeNumber:     Da.EmployeeNumber,
			CodeStore:          Da.CodeStore,
			DepartmentCode:     Da.DepartmentCode,
			StoreSection:       Da.StoreSection,
			UnitCode:           Da.UnitCode,
			DivisiCodeAndName:  Da.DivisiCodeAndName,
			AddressAndNumber:   Da.AddressAndNumber,
			SumDistance:        sumDistance,
			SumCost:            sumCost,
			CodeCommuting:      Da.CodeCommuting,
			IdBasicInformation: Da.IdBasicInformation,
		}

		sh = append(sh, InitDataScan)

	}

	return sh, nil, CountData
}

// View data Commuting Agregat (SUM) By All Employee Code
func (model Init_DB_CommutingApprove) GetDataApproveByCommutingEmployeeCode(page string, showData string, searching string, employee_number string) (sh []approve.Init_CommutingApprove, err error, CountData int) {

	var Da approve.Init_CommutingApprove

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
		limitPage = ` LIMIT ` + DataPageInt + `,` + DataShowDataInt
	}

	searchingAction := ``
	if searching == "" {
		searchingAction = ``
	} else {
		searchingAction = ` and (store_information.code_store LIKE '% ` + searching + `%' OR department_information.department_name LIKE '%` + searching + `%' OR store_section_information.store_section_name LIKE '%` + searching + `%' OR basic_information.first_name LIKE '%` + searching + `%' OR basic_information.last_name LIKE '%` + searching + `%' OR basic_information.adress LIKE '%` + searching + `%' OR basic_information.adress_kana LIKE '%` + searching + `%' OR basic_information.adress_detail LIKE '%` + searching + `%' OR basic_information.adress_detail_kana LIKE '%` + searching + `%' OR basic_information.add_phone_number LIKE '%` + searching + `%')`
	}

	RowCountData := model.DB.QueryRow(` select count(*) from (SELECT      MIN(
            commuting_trip.id_commuting_trip
        ) id_commuting_trip,
        MIN(store_information.code_store) AS code_store,
        MIN(
            department_information.department_name
        ),
        CONCAT(
            MIN(
                store_section_information.store_section_name
            ),
            ' ',
            MIN(
                store_section_information.store_section_code
            )
        ) AS store_section,
        MIN(unit_information.unit_code),
        CONCAT(
            MIN(basic_information.first_name),
            ' ',
            MIN(basic_information.last_name),' ',
        CONCAT(
            RIGHT(
                MIN(store_information.code_store),
                4
            ),
            LPAD(
                RIGHT(
                    MIN(
                        department_information.department_code
                    ),
                    2
                ),
                2,
                '0'
            ),
            LPAD(
                RIGHT(
                    MIN(
                        store_section_information.store_section_code
                    ),
                    2
                ),
                2,
                '0'
            ),
            LPAD(
                RIGHT(
                    MIN(unit_information.unit_code),
                    2
                ),
                2,
                '0'
            )
        )
        ),
        CONCAT(
        MIN(basic_information.adress),', ',
        MIN(basic_information.adress_kana),', ',
        MIN(
            basic_information.adress_detail
        ), ', ', MIN(
            basic_information.adress_detail_kana
        ),' ',  MIN(
            basic_information.add_phone_number
        ) 
        ),
        MIN(commuting_trip.date) date_trip,
        MIN(
            basic_information.employee_code
        ) employee_code,
        MIN(
            code_commuting.status_commuting
        ),
        SUM(detail_commuting_trip.cost) as sum_cost,
        SUM(detail_commuting_trip.distance) as sum_distance,
        MIN(basic_information.id_basic_information),
        MIN(code_commuting.code_random)
        FROM code_commuting LEFT OUTER JOIN commuting_trip ON code_commuting.code_random = commuting_trip.code_commuting LEFT OUTER JOIN general_information ON commuting_trip.id_general_information = general_information.id_general_information LEFT OUTER JOIN store_information ON store_information.id_code_store = general_information.id_store_code LEFT OUTER JOIN basic_information ON general_information.id_basic_information = basic_information.id_basic_information LEFT OUTER JOIN department_information ON general_information.id_department =department_information.id_department LEFT OUTER JOIN store_section_information ON general_information.id_store_section = store_section_information.id_store_section LEFT OUTER JOIN unit_information ON unit_information.id_unit = general_information.id_unit
        LEFT OUTER JOIN detail_commuting_trip ON detail_commuting_trip.id_commuting_trip = commuting_trip.id_commuting_trip
        WHERE commuting_trip.submit = 'Y' AND basic_information.employee_code = ? `+searchingAction+`	GROUP BY code_commuting.id_code ORDER BY code_commuting.id_code DESC) t`, employee_number).Scan(&CountData)

	if RowCountData != nil {
		log.Println(RowCountData)
	}

	showDataApprove, errShowDataApprove := model.DB.Query(
		`SELECT      MIN(
            commuting_trip.id_commuting_trip
        ) id_commuting_trip,
        MIN(store_information.code_store) AS code_store,
        MIN(
            department_information.department_name
        ),
        CONCAT(
            MIN(
                store_section_information.store_section_name
            ),
            ' ',
            MIN(
                store_section_information.store_section_code
            )
        ) AS store_section,
        MIN(unit_information.unit_code),
        CONCAT(
            MIN(basic_information.first_name),
            ' ',
            MIN(basic_information.last_name),' ',
        CONCAT(
            RIGHT(
                MIN(store_information.code_store),
                4
            ),
            LPAD(
                RIGHT(
                    MIN(
                        department_information.department_code
                    ),
                    2
                ),
                2,
                '0'
            ),
            LPAD(
                RIGHT(
                    MIN(
                        store_section_information.store_section_code
                    ),
                    2
                ),
                2,
                '0'
            ),
            LPAD(
                RIGHT(
                    MIN(unit_information.unit_code),
                    2
                ),
                2,
                '0'
            )
        )
        ),
        CONCAT(
        MIN(basic_information.adress),', ',
        MIN(basic_information.adress_kana),', ',
        MIN(
            basic_information.adress_detail
        ), ', ', MIN(
            basic_information.adress_detail_kana
        ),' ',  MIN(
            basic_information.add_phone_number
        ) 
        ),
        MIN(commuting_trip.date) date_trip,
        MIN(
            basic_information.employee_code
        ) employee_code,
        MIN(
            code_commuting.status_commuting
        ),
        SUM(detail_commuting_trip.cost) as sum_cost,
        SUM(detail_commuting_trip.distance) as sum_distance,
        MIN(basic_information.id_basic_information),
        MIN(code_commuting.code_random)
        FROM code_commuting LEFT OUTER JOIN commuting_trip ON code_commuting.code_random = commuting_trip.code_commuting LEFT OUTER JOIN general_information ON commuting_trip.id_general_information = general_information.id_general_information LEFT OUTER JOIN store_information ON store_information.id_code_store = general_information.id_store_code LEFT OUTER JOIN basic_information ON general_information.id_basic_information = basic_information.id_basic_information LEFT OUTER JOIN department_information ON general_information.id_department =department_information.id_department LEFT OUTER JOIN store_section_information ON general_information.id_store_section = store_section_information.id_store_section LEFT OUTER JOIN unit_information ON unit_information.id_unit = general_information.id_unit
        LEFT OUTER JOIN detail_commuting_trip ON detail_commuting_trip.id_commuting_trip = commuting_trip.id_commuting_trip
        WHERE commuting_trip.submit = 'Y' AND basic_information.employee_code = ? `+searchingAction+`	GROUP BY code_commuting.id_code ORDER BY code_commuting.id_code DESC`+limitPage, employee_number)

	if errShowDataApprove != nil {
		log.Println(errShowDataApprove)
	}

	var Datatrip models.NullString
	var StatusCommuting models.NullString
	var sumCost int64
	var sumDistance float64
	for showDataApprove.Next() {
		ScanData := showDataApprove.Scan(&Da.IdCommuting, &Da.CodeStore, &Da.DepartmentCode, &Da.StoreSection, &Da.UnitCode, &Da.DivisiCodeAndName, &Da.AddressAndNumber, &Datatrip, &Da.EmployeeNumber, &StatusCommuting, &sumCost, &sumDistance, &Da.IdBasicInformation, &Da.CodeCommuting)

		if ScanData != nil {
			log.Println(ScanData)
		}

		InitDataScan := approve.Init_CommutingApprove{
			IdCommuting:        Da.IdCommuting,
			EmployeeNumber:     Da.EmployeeNumber,
			CodeStore:          Da.CodeStore,
			DepartmentCode:     Da.DepartmentCode,
			StoreSection:       Da.StoreSection,
			UnitCode:           Da.UnitCode,
			DivisiCodeAndName:  Da.DivisiCodeAndName,
			AddressAndNumber:   Da.AddressAndNumber,
			SumDistance:        sumDistance,
			SumCost:            sumCost,
			CodeCommuting:      Da.CodeCommuting,
			IdBasicInformation: Da.IdBasicInformation,
		}

		sh = append(sh, InitDataScan)

	}

	return sh, nil, CountData
}

func (model Init_DB_CommutingApprove) DetailCommutingByEmployeeCode(employee_number string, id_basic_information string, CodeCommuting string) (sh []approve.FormatDataDetailCommutingByEmployeeCode, condition string) {

	var Da approve.Init_DataCommutingByEmployeeCodeApprove
	var DaD approve.Init_DetailCommutingByEmployeeCodeApprove
	var ArrDaD []approve.Init_DetailCommutingByEmployeeCodeApprove

	SumAllCost := 0
	SumAllDistance := 0.00

	GetSumData := model.DB.QueryRow(`select COALESCE(SUM(cost),0) as SUM_COST, COALESCE(SUM(distance),0.00) as SUM_DISTANCE from (select COALESCE(SUM(detcomtrip.cost),0)	as cost, COALESCE(SUM(detcomtrip.distance),0.00) as distance						from commuting_trip comtrip, code_commuting cc,
										detail_commuting_trip detcomtrip, general_information geninfo, basic_information bainfo, store_information storeinfo
										where comtrip.id_commuting_trip = detcomtrip.id_commuting_trip and geninfo.id_general_information = comtrip.id_general_information AND
										geninfo.id_basic_information = bainfo.id_basic_information and geninfo.id_store_code = storeinfo.id_code_store   and cc.code_random = comtrip.code_commuting
										and bainfo.employee_code = ? and bainfo.id_basic_information = ? and comtrip.code_commuting = ? and comtrip.save_trip ='N' and comtrip.submit = 'Y' and comtrip.status_approval = 'Y'
										group by detcomtrip.id_commuting_trip order by comtrip.date asc) t`,employee_number,id_basic_information,CodeCommuting).Scan(&SumAllCost,&SumAllDistance)

	if GetSumData != nil {
		log.Println(GetSumData)
	}

	GetData, errData := model.DB.Query(`select basic_information.first_name, basic_information.last_name, store_information.code_store, store_information.store_name,
basic_information.employee_code, commuting_basic_information.driver_license_expiry_date,commuting_basic_information.car_insurance_document_expiry_date,
(select COUNT(*) from basic_information a, commuting_basic_information b, store_information c, general_information d
where b.id_general_information = d.id_general_information and d.id_basic_information = a.id_basic_information and d.id_store_code = c.id_code_store and a.employee_code=basic_information.employee_code and commuting_basic_information.driver_license_expiry_date = DATE_FORMAT(CONVERT_TZ(NOW(), @@session.time_zone, '+09:00'),'%Y-%m-%d') ) as driver_license_status,
(select COUNT(*) from basic_information basic_informationA, commuting_basic_information commuting_basic_informationB, store_information store_informationC , general_information general_informationD
where commuting_basic_informationB.id_general_information = general_informationD.id_general_information and general_informationD.id_basic_information = basic_informationA.id_basic_information and general_informationD.id_store_code = store_informationC.id_code_store and basic_informationA.employee_code=basic_information.employee_code and commuting_basic_informationB.car_insurance_document_expiry_date = DATE_FORMAT(CONVERT_TZ(NOW(), @@session.time_zone, '+09:00'),'%Y-%m-%d') ) as Car_Insuranse_status
from basic_information, commuting_basic_information, store_information , general_information
where commuting_basic_information.id_general_information = general_information.id_general_information and general_information.id_basic_information = basic_information.id_basic_information and general_information.id_store_code = store_information.id_code_store and basic_information.employee_code=?`, employee_number)

	if errData != nil {
		log.Println(errData)
	}

	NextData := GetData.Next()

	if NextData == false {
		log.Println("FALSE?")
		log.Println(NextData)
	}
	StatusDriverLicenseExpiryDate := 0
	StatusCarInsuranceExpiryDate := 0
	StringDriverLicenseDate := ""
	StringCarInsuranceDate := ""
	errScanData := GetData.Scan(&Da.FirstName, &Da.LastName, &Da.CodeStore, &Da.StoreName, &Da.EmployeeCode, &Da.DriverLicenseExpiryDate, &Da.CarInsuranceExpiryDate, &StatusDriverLicenseExpiryDate, &StatusCarInsuranceExpiryDate)
	if errScanData != nil {
		log.Println("ERR SCAN ?")
		log.Println(errScanData)
	}
	QueryCheckDriverLicenseExpiryDate := model.DB.QueryRow(`select count(*) from commuting_basic_information where driver_license_expiry_date <= DATE_FORMAT(CONVERT_TZ(NOW(), @@session.time_zone, '+09:00'),'%Y-%m-%d') and driver_license_expiry_date = ?`,Da.DriverLicenseExpiryDate).Scan(&StatusDriverLicenseExpiryDate)

	if QueryCheckDriverLicenseExpiryDate != nil {
		log.Println(QueryCheckDriverLicenseExpiryDate)
	}
	QueryCheckCarInsuranceDate := model.DB.QueryRow(`select count(*) from commuting_basic_information where car_insurance_document_expiry_date <= DATE_FORMAT(CONVERT_TZ(NOW(), @@session.time_zone, '+09:00'),'%Y-%m-%d') and car_insurance_document_expiry_date = ?`,Da.CarInsuranceExpiryDate).Scan(&StatusCarInsuranceExpiryDate)
	if QueryCheckCarInsuranceDate != nil {
		log.Println(QueryCheckCarInsuranceDate)
	}
	if StatusDriverLicenseExpiryDate == 0 {
		StringDriverLicenseDate = "yes"
	} else {
		StringDriverLicenseDate = "no"
	}
	if StatusCarInsuranceExpiryDate == 0 {
		StringCarInsuranceDate = "yes"
	} else {
		StringCarInsuranceDate = "no"
	}
	var ContainerDataBasicInformation interface{}
	BasicInformation := approve.Init_DataCommutingByEmployeeCodeApprove{
		FirstName:                     Da.FirstName,
		LastName:                      Da.LastName,
		CodeStore:                     Da.CodeStore,
		StoreName:                     Da.StoreName,
		EmployeeCode:                  employee_number,
		DriverLicenseExpiryDate:       Da.DriverLicenseExpiryDate,
		CarInsuranceExpiryDate:        Da.CarInsuranceExpiryDate,
		StatusDriverLicenseExpiryDate: StringDriverLicenseDate,
		StatusCarInsuranceExpiryDate:  StringCarInsuranceDate,
		SumCost:                       SumAllCost,
		SumDistance:                   SumAllDistance,
	}

	ContainerDataBasicInformation = BasicInformation


	GetDataDetail, errDataDetail := model.DB.Query(`select  MIN(comtrip.id_commuting_trip), MIN(detcomtrip.id_detail_commuting_trip), comtrip.date, MIN(comtrip.route_profile_name), MIN(comtrip.attendance_code), 
										MIN(detcomtrip.purpose), COALESCE(SUM(detcomtrip.distance),0), COALESCE(SUM(detcomtrip.commute_distance),0) , COALESCE(SUM(detcomtrip.cost),0), MIN(cc.status_commuting), CAST(comtrip.date_time_approve as DATE) as date_time_approve
										from commuting_trip comtrip, code_commuting cc,
										detail_commuting_trip detcomtrip, general_information geninfo, basic_information bainfo, store_information storeinfo
										where comtrip.id_commuting_trip = detcomtrip.id_commuting_trip and geninfo.id_general_information = comtrip.id_general_information AND
										geninfo.id_basic_information = bainfo.id_basic_information and geninfo.id_store_code = storeinfo.id_code_store   and cc.code_random = comtrip.code_commuting
										and bainfo.employee_code = ? and bainfo.id_basic_information = ? and comtrip.code_commuting = ? and comtrip.save_trip ='N' and comtrip.submit = 'Y' 
										group by detcomtrip.id_commuting_trip order by comtrip.date asc`, employee_number, id_basic_information, CodeCommuting)

	if errDataDetail != nil {
		log.Println(errDataDetail)
	}

	for GetDataDetail.Next() {

		errScanDataDetail := GetDataDetail.Scan(&DaD.IdCommutingTrip, &DaD.IdDetailCommutingTrip, &DaD.Date, &DaD.RouteProfileName, &DaD.AttendanceCode, &DaD.Purpose, &DaD.Distance, &DaD.CommuteDistance, &DaD.Cost, &DaD.StatusCommuting, &DaD.DateApprove)
		if errScanData != nil {
			log.Println(errScanDataDetail)
		}
		DatatypeOfTransportation, DataRoute, _ := utils_enter_the_information.GetAdditionalUsageRecord("", employee_number, DaD.IdCommutingTrip, "DetailCommutingByEmployeeCode")

		DataDetailInit := approve.Init_DetailCommutingByEmployeeCodeApprove{
			IdCommutingTrip:       DaD.IdCommutingTrip,
			IdDetailCommutingTrip: DaD.IdDetailCommutingTrip,
			RouteProfileName:      DaD.RouteProfileName,
			Date:                  DaD.Date,
			TypeOfTransport:       DatatypeOfTransportation,
			AttendanceCode:        DaD.AttendanceCode,
			Purpose:               DaD.Purpose,
			Distance:              DaD.Distance,
			CommuteDistance:       DaD.CommuteDistance,
			Cost:                  DaD.Cost,
			Route:                 DataRoute,
			StatusCommuting:       DaD.StatusCommuting,
			DateApprove:           DaD.DateApprove,
		}
		ArrDaD = append(ArrDaD, DataDetailInit)
		//append data looping

	}

	FinallyData := approve.FormatDataDetailCommutingByEmployeeCode{
		Data:       ContainerDataBasicInformation,
		DataDetail: ArrDaD,
	}
	sh = append(sh, FinallyData)

	return sh, `Success Response`
}

func (model Init_DB_CommutingApprove) CommutingApproveOrReject(dataApprove []approve.Init_InputDataApprove) (dataapprove []approve.Init_InputDataApprove, condition string) {

	// merah T
	// hijau Y
	// datetime approve -> waktu tokyo

	//vals := []interface{}{}
	//sqlUpdate := `insert into commuting_trip(id_commuting_trip,status_approval,date_time_approve) values `
	//sqlUpdate := `update commuting_trip set status_approval = ?,date_time_approve =? where id_commuting_trip = ? `
	//sqlUpdate := ""
	//var dataapprover approve.Init_InputDataApprove
	ctx := context.Background()
	tx, errTx := model.DB.BeginTx(ctx, nil)
	if errTx != nil {
		log.Fatal(errTx.Error())
	}
	for _, dataapprover := range dataApprove {

		//sqlUpdate += "(?,?,CONVERT_TZ(NOW(),'+00:00','+09:00')),"
		_, res := model.DB.ExecContext(ctx, `update commuting_trip set status_approval = ?,date_time_approve = CONVERT_TZ(NOW(),@@session.time_zone,'+09:00') where id_commuting_trip = ?`, dataapprover.StatusDataApprove, dataapprover.IdCommuting)
		if res != nil {
			log.Println(res.Error())
			tx.Rollback()
			return
		}
		dataapprove = append(dataapprove, dataapprover)

		//log.Println(res.RowsAffected())
		//log.Println(dataapprover.IdCommuting)
		//log.Println(dataapprover.StatusDataApprove)
		//sqlUpdate += `update commuting_trip set status_approval = ?,date_time_approve =? where id_commuting_trip = ? `
		//vals = append(vals,  dataapprover.StatusDataApprove,dataapprover.IdCommuting)
	}
	CommitErr := tx.Commit()
	if CommitErr != nil {
		log.Println(CommitErr.Error())
	}

	//log.Println(sqlUpdate)
	//log.Println(dataapprover)
	//sqlUpdate += `;`
	//sqlUpdate = sqlUpdate[0 : len(sqlUpdate)-1]
	// id data ganti id
	//sqlUpdate += ` ON DUPLICATE KEY UPDATE id_commuting_trip=VALUES(id_commuting_trip),status_approval=VALUES(status_approval),date_time_approve=VALUES(date_time_approve)`
	//stmtApprove, _ := model.DB.Prepare(sqlUpdate)

	//res, _ := stmtApprove.Exec(vals...)
	//if res == nil {
	//	log.Println(res)
	//}

	return dataapprove, `Success Response`
}
