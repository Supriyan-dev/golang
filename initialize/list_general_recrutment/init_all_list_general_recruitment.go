package list_general_recrutment

import (
	"../../models"
)

type ShowAllListGeneralRecrutment struct {
	IdBasicInformation      int               `json:"id_basic_information"`
	IdEmploymentType        int               `json:"id_employment_type"`
	EmployeeCode            string            `json:"employee_code"`
	FirstName               models.NullString `json:"first_name"`
	LastName                models.NullString `json:"last_name"`
	JoinDate                models.NullString `json:"join_date"`
	CodeStore               models.NullString            `json:"code_store"`
	StoreSectionName        models.NullString `json:"store_section_name"`
	EmploymentType          string            `json:"employment_type"`
	StartWorking            string            `json:"start_working"`
	EndWorking              string            `json:"end_working"`
	EmploymentStatusApprove models.NullString
}

type ShowDetailListGeneralRecrutment struct {
	//basic information
	FirstName        models.NullString `json:"first_name"`
	LastName         models.NullString `json:"last_name"`
	EmployeeCode     string            `json:"employee_code"`
	AddPhoneNumber   models.NullString `json:"add_phone_number"`
	AddPostalCode    models.NullString `json:"add_postal_code"`
	PrefectureName   models.NullString `json:"prefecture_name"`
	Adress           models.NullString `json:"adress"`
	AdressKana       models.NullString `json:"adress_kana"`
	AdressDetail     models.NullString `json:"adress_detail"`
	AdressDetailKana models.NullString `json:"adress_detail_kana"`
	MaritalStatus    models.NullString `json:"marital_status"`
	DormitoryStatus  models.NullString `json:"dormitory_status"`
	Birthdate        models.NullString `json:"birthdate"`
	// data 2
	CodeStore                                models.NullString `json:"code_store"`
	DepartmentCode                           models.NullString `json:"department_code"`
	UnitCode                                 models.NullString `json:"unit_code"`
	JoinDate                                 models.NullString `json:"join_date"`
	BankCode                                 models.NullString `json:"bank_code"`
	FirstSmesterInOtherCompany               string            `json:"first_smester_in_other_company"`
	DistanceTrip                             string            `json:"distance_trip"`
	ResumeDocument                           string            `json:"resume_document"`
	WrittenOathDocument                      string            `json:"written_oath_document"`
	EmployeeAgreementDocument                string            `json:"employee_agreement_document"`
	CertificateOfResidenceCardDocument       string            `json:"certificate_of_residence_card_document"`
	ApplicationFormOfCommutingMethodDocument string            `json:"application_form_of_commuting_method_document"`
	ComplianceAgreementDocument              string            `json:"compliance_agreement_document"`
	WithHoldingSlipDocument                  string            `json:"with_holding_slip_document"`
	DependentDeductionFormDocument           string            `json:"dependent_deduction_form_document"`
	PensionBookDocument                      string            `json:"pension_book_document"`
	HealthCheckReportDocument                string            `json:"health_check_report_document"`
	// data 3
	Salary                int            `json:"salary"`
	EmploymentInsuranceNo string         `json:"employment_insurance_no"`
	PensionNo             string         `json:"pension_no"`
	StartWorking          string         `json:"start_working"`
	EndWorking            string         `json:"end_working"`
	TotalMonthlySalary    models.NullInt `json:"total_monthly_salary"`
}
