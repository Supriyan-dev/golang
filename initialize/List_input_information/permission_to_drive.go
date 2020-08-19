package initialize

// type GeneralInformation struct {
// 	id_general_information                        int    `json:"id_general_information"`
// 	id_basic_information                          int    `json:"id_basic_information "`
// 	id_store_code                                 int    `json:"id_store_code "`
// 	id_department                                 int    `json:"id_department "`
// 	id_store_section                              int    `json:"id_store_section "`
// 	id_unit                                       int    `json:"id_unit "`
// 	join_date                                     string `json:"join_date"`
// 	id_bank                                       int    `json:"id_bank "`
// 	account_type                                  string `json:"account_type"`
// 	account_number                                int    `json:"account_number"`
// 	account_name                                  string `json:"account_name"`
// 	first_smester_in_other_company                string `json:"first_smester_in_other_company"`
// 	distance_trip                                 string `json:"distance_trip"`
// 	resume_document                               string `json:"resume_document"`
// 	written_oath_document                         string `json:"written_oath_document"`
// 	employee_agreement_document                   string `json:"employee_agreement_document"`
// 	certificate_of_residence_card_document        string `json:"certificate_of_residence_card_document"`
// 	application_form_of_commuting_method_document string `json:"application_form_of_commuting_method_document"`
// 	compliance_agreement_document                 string `json:"compliance_agreement_document"`
// 	with_holding_slip_document                    string `json:"with_holding_slip_document"`
// 	dependent_deduction_form_document             string `json:"dependent_deduction_form_document"`
// 	pension_book_document                         string `json:"pension_book_document"`
// 	health_check_report_document                  string `json:"health_check_report_document"`
// 	office_code                                   string `json:"office_code"`
// }

// type StoreInformation struct {
// 	Id_code_store int    `json:"id_code_store"`
// 	Code_store    string `json:"code_store"`
// 	Store_name    string `json:"store_name"`
// }

// type BasicInformation struct {
// 	id_basic_information int    `json:"id_basic_information"`
// 	employee_code        int    `json:"employee_code "`
// 	first_name           string `json:"first_name "`
// 	last_name            string `json:"last_name "`
// 	gender               string `json:"gender "`
// 	birthdate            string `json:"birthdate "`
// 	add_postal_code      string `json:"add_postal_code "`
// 	id_prefecture        string `json:"id_prefecture "`
// 	adress               string `json:"adress "`
// 	adress_kana          string `json:"adress_kana"`
// 	adress_detail        string `json:"adress_detail"`
// 	adress_detail_kana   string `json:"adress_detail_kana"`
// 	add_phone_number     string `json:"add_phone_number"`
// 	marital_status       string `json:"marital_status"`
// 	dormitory_status     string `json:"dormitory_status"`
// }

// type CommutingBasicInformation struct {
// 	id_commuting_basic_information     int    `json:"id_commuting_basic_information "`
// 	id_general_information             int    `json:"id_general_information  "`
// 	driver_license_document            string `json:"driver_license_document "`
// 	driver_license_document_url        string `json:"driver_license_document_url "`
// 	driver_license_expiry_date         string `json:"driver_license_expiry_date "`
// 	car_insurance_document             string `json:"car_insurance_document "`
// 	car_insurance_document_url         string `json:"car_insurance_document_url "`
// 	car_insurance_document_expiry_date string `json:"car_insurance_document_expiry_date "`
// 	daily_commuting_method             string `json:"daily_commuting_method"`
// 	default_transportation             string `json:"default_transportation"`
// 	permitted_to_drive                 string `json:"permitted_to_drive"`
// 	insurance_company                  string `json:"insurance_company"`
// 	personal_injury                    string `json:"personal_injury"`
// 	property_damage                    string `json:"property_damage"`
// 	status_approve                     string `json:"status_approve"`
// 	date_approve                       string `json:"date_approve"`
// 	time_approve                       string `json:"time_approve"`
// 	date_submit                        string `json:"date_submit"`
// }

type Join struct {
	id_general_information             int    `json:"id_general_information"`
	id_basic_information               int    `json:"id_basic_information "`
	id_store_code                      int    `json:"id_store_code "`
	code_store                         string `json:"code_store"`
	employee_code                      int    `json:"employee_code "`
	first_name                         string `json:"first_name "`
	last_name                          string `json:"last_name "`
	driver_license_expiry_date         string `json:"driver_license_expiry_date "`
	car_insurance_document_expiry_date string `json:"car_insurance_document_expiry_date "`
}
