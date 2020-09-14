package list_general_recruitment

import (
	init_list_GR "../../initialize/list_general_recrutment"
	"../../models"
	utils_working_time "../../utils/list_general_recruitment"
	"log"
	"strconv"
)

type Models_init_listGeneralRecruitment models.DB_init

// 4 form approve, reject, not approve, all data
func (model Models_init_listGeneralRecruitment) GetListGeneralRecruitment(status string, page string, filter string, showData string, searching string) (sh []init_list_GR.ShowAllListGeneralRecrutment, err error, CountData int) {

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
	filterMonth := ``
	if filter == "" {
		filterMonth = ``
	} else {
		filterMonth = ` and MONTH(gi.join_date) =` + filter
	}
	searchingAction := ``
	if searching == "" {
		searchingAction = ``
	} else {
		searchingAction = ` and (bi.employee_code LIKE '% ` + searching + `%' OR bi.first_name LIKE '%` + searching + `%' OR bi.last_name LIKE '%` + searching + `%'OR si.code_store LIKE '%` + searching + `%' OR ssi.store_section_name LIKE '%` + searching + `%' )`
	}

	queryStatus := ""

	if status == "all" {
		queryStatus = ``
	} else if status == "approve" {
		queryStatus = ` AND gisa.status ='approve'`
	} else if status == "waiting" {
		queryStatus = ` AND gisa.status IS NULL`
	} else if status == "reject" {
		queryStatus = ` AND gisa.status ='reject'`
	} else {
		return nil, nil, 0
	}

	queryCountAllList := `select count(*) from (select et.id_employment_type , et.employment_status, bi.id_basic_information, bi.employee_code, bi.first_name, bi.last_name, gi.join_date, si.code_store, ssi.store_section_name, gisa.status from basic_information bi, general_information gi, employment_type et,general_recruitment_status_approval gisa, store_information si , store_section_information ssi
		where bi.id_basic_information = gi.id_basic_information and bi.id_basic_information = et.id_basic_information
        and gi.id_store_code = si.id_code_store and gi.id_store_section = ssi.id_store_section
		and gisa.id_basic_information = et.id_basic_information` + filterMonth + queryStatus + searchingAction + `) t`

	errGetCountData := model.DB.QueryRow(queryCountAllList).Scan(&CountData)

	if errGetCountData != nil {
		log.Println(errGetCountData)
	}

	queryAllList := `select et.id_employment_type , et.employment_status, bi.id_basic_information, bi.employee_code, bi.first_name, bi.last_name, gi.join_date, si.code_store, ssi.store_section_name, gisa.status from basic_information bi, general_information gi, employment_type et,general_recruitment_status_approval gisa, store_information si , store_section_information ssi
		where bi.id_basic_information = gi.id_basic_information and bi.id_basic_information = et.id_basic_information
        and gi.id_store_code = si.id_code_store and gi.id_store_section = ssi.id_store_section
		and gisa.id_basic_information = et.id_basic_information` + filterMonth + queryStatus + searchingAction + ` order by bi.employee_code asc ` + limitPage

	GetAllDataGR, errGetAllDataGR := model.DB.Query(queryAllList)

	if errGetAllDataGR != nil {
		log.Println(errGetAllDataGR)
		return nil, errGetAllDataGR, CountData
	}

	var dGR init_list_GR.ShowAllListGeneralRecrutment

	for GetAllDataGR.Next() {

		errScanGetAllDataGR := GetAllDataGR.Scan(&dGR.IdEmploymentType, &dGR.EmploymentType, &dGR.IdBasicInformation, &dGR.EmployeeCode, &dGR.FirstName, &dGR.LastName, &dGR.JoinDate, &dGR.CodeStore, &dGR.StoreSectionName, &dGR.EmploymentStatusApprove)

		if errScanGetAllDataGR != nil {
			log.Println(errScanGetAllDataGR)
			return nil, errScanGetAllDataGR, CountData
		}
		start_working, end_working := utils_working_time.GetDataWorkingTime(dGR.IdEmploymentType, dGR.EmploymentType, dGR.IdBasicInformation)

		FinnalyData := init_list_GR.ShowAllListGeneralRecrutment{
			IdBasicInformation:      dGR.IdBasicInformation,
			IdEmploymentType:        dGR.IdEmploymentType,
			EmployeeCode:            dGR.EmployeeCode,
			FirstName:               dGR.FirstName,
			LastName:                dGR.LastName,
			JoinDate:                dGR.JoinDate,
			CodeStore:               dGR.CodeStore,
			StoreSectionName:        dGR.StoreSectionName,
			EmploymentType:          dGR.EmploymentType,
			StartWorking:            start_working,
			EndWorking:              end_working,
			EmploymentStatusApprove: dGR.EmploymentStatusApprove,
		}

		sh = append(sh, FinnalyData)

	}

	return sh, nil, CountData

}

func (model Models_init_listGeneralRecruitment) GetDetailListGeneralRecruitment(IdBasicInformation int, EmployeeType string) (sh init_list_GR.ShowDetailListGeneralRecrutment, condition string) {

	query := `select bi.first_name, bi.last_name, bi.employee_code, bi.add_phone_number, bi.add_postal_code, pre.prefecture_name, bi.adress,
bi.adress_kana, bi.adress_detail, bi.adress_detail_kana, bi.marital_status, bi.dormitory_status, bi.birthdate,
CONCAT(si.code_store,' (',si.store_name,')') as store_code, 
CONCAT(di.department_code,' (',di.department_name,')') as department_code, 
CONCAT(ssi.store_section_code,' (',ssi.store_section_name,')'),
CONCAT(bank.bank_code,' - ',bank.bank_name,' - ',bank.branch_code,' - ', bank.branch_name),
gi.join_date,gi.first_smester_in_other_company, gi.distance_trip, gi.resume_document, gi.written_oath_document, gi.employee_agreement_document, gi.certificate_of_residence_card_document,
gi.application_form_of_commuting_method_document, gi.compliance_agreement_document, gi.with_holding_slip_document, gi.dependent_deduction_form_document, gi.pension_book_document,
gi.health_check_report_document, et.total_monthly_salary
from basic_information bi
LEFT OUTER JOIN general_information gi ON gi.id_basic_information = bi.id_basic_information
LEFT OUTER JOIN prefecture pre ON pre.id_prefecture = bi.id_prefecture
LEFT OUTER JOIN employment_type et ON et.id_basic_information = bi.id_basic_information 
LEFT OUTER JOIN general_recruitment_status_approval gisa ON gisa.id_basic_information = bi.id_basic_information 
LEFT OUTER JOIN store_information si ON si.id_code_store = gi.id_store_code 
LEFT OUTER JOIN store_section_information ssi ON ssi.id_store_section = gi.id_store_section
LEFT OUTER JOIN bank bank ON bank.id_bank = gi.id_bank
LEFT OUTER JOIN unit_information ui ON ui.id_unit = gi.id_unit
LEFT OUTER JOIN department_information di ON di.id_department = gi.id_department
where  bi.id_basic_information = ? and et.employment_status = ?`

	GetDataDetail, errGetDataDetail := model.DB.Query(query, IdBasicInformation, EmployeeType)

	if errGetDataDetail != nil {
		log.Println(errGetDataDetail)
	}

	nextData := GetDataDetail.Next()

	if nextData == false {
		return sh, "Missing Body Request"
	}
	var gdda init_list_GR.ShowDetailListGeneralRecrutment
	ScanData := GetDataDetail.Scan(&gdda.FirstName,&gdda.LastName,&gdda.EmployeeCode,&gdda.AddPhoneNumber,&gdda.AddPostalCode,&gdda.PrefectureName,&gdda.Adress,&gdda.AdressKana,&gdda.AdressDetail,&gdda.AdressDetailKana,&gdda.MaritalStatus,&gdda.DormitoryStatus,&gdda.Birthdate,&gdda.CodeStore,&gdda.DepartmentCode,&gdda.UnitCode,&gdda.BankCode,&gdda.JoinDate,&gdda.FirstSmesterInOtherCompany,&gdda.DistanceTrip,&gdda.ResumeDocument,&gdda.WrittenOathDocument,&gdda.EmployeeAgreementDocument,&gdda.CertificateOfResidenceCardDocument,&gdda.ApplicationFormOfCommutingMethodDocument,&gdda.ComplianceAgreementDocument,&gdda.WithHoldingSlipDocument,&gdda.DependentDeductionFormDocument,&gdda.PensionBookDocument,&gdda.HealthCheckReportDocument,&gdda.TotalMonthlySalary)

	if ScanData != nil {
		return sh, ScanData.Error()
	}

	salary,EmployentInsuranceNo,PensionNo,start_working,end_working := utils_working_time.GetDataDetailWorkingTime(EmployeeType,IdBasicInformation)

	EndData := init_list_GR.ShowDetailListGeneralRecrutment{
		FirstName:                                gdda.FirstName,
		LastName:                                 gdda.LastName,
		EmployeeCode:                             gdda.EmployeeCode,
		AddPhoneNumber:                           gdda.AddPhoneNumber,
		AddPostalCode:                            gdda.AddPostalCode,
		PrefectureName:                           gdda.PrefectureName,
		Adress:                                   gdda.Adress,
		AdressKana:                               gdda.AdressKana,
		AdressDetail:                             gdda.AdressDetail,
		AdressDetailKana:                         gdda.AdressDetailKana,
		MaritalStatus:                            gdda.MaritalStatus,
		DormitoryStatus:                          gdda.DormitoryStatus,
		Birthdate:                                gdda.Birthdate,
		CodeStore:                                gdda.CodeStore,
		DepartmentCode:                           gdda.DepartmentCode,
		UnitCode:                                 gdda.UnitCode,
		JoinDate:                                 gdda.JoinDate,
		BankCode:                                 gdda.BankCode,
		FirstSmesterInOtherCompany:               gdda.FirstSmesterInOtherCompany,
		DistanceTrip:                             gdda.DistanceTrip,
		ResumeDocument:                           gdda.ResumeDocument,
		WrittenOathDocument:                      gdda.WrittenOathDocument,
		EmployeeAgreementDocument:                gdda.EmployeeAgreementDocument,
		CertificateOfResidenceCardDocument:       gdda.CertificateOfResidenceCardDocument,
		ApplicationFormOfCommutingMethodDocument: gdda.ApplicationFormOfCommutingMethodDocument,
		ComplianceAgreementDocument:              gdda.ComplianceAgreementDocument,
		WithHoldingSlipDocument:                  gdda.WithHoldingSlipDocument,
		DependentDeductionFormDocument:           gdda.DependentDeductionFormDocument,
		PensionBookDocument:                      gdda.PensionBookDocument,
		HealthCheckReportDocument:                gdda.HealthCheckReportDocument,
		Salary:                                   salary,
		EmploymentInsuranceNo:                    EmployentInsuranceNo,
		PensionNo:                                PensionNo,
		StartWorking:                             start_working,
		EndWorking:                               end_working,
		TotalMonthlySalary:                       gdda.TotalMonthlySalary,
	}

	return EndData, "Success Response"

}
