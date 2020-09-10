package initialize

import (
	"../../models"
)

type BasicInformationGeneral struct {
	Id_basic_information int               `json:"id_basic_information"`
	Employee_code        int               `json:"employee_code"`
	First_name           string            `json:"first_name"`
	Last_name            string            `json:"last_name"`
	Gender               string            `json:"gender"`
	Add_postal_code      models.NullString `json:"add_postal_code"`
	Id_prefecture        int               `json:"id_prefecture"`
	Adress               models.NullString `json:"addres"`
	Adress_kana          models.NullString `json:"adress_kana"`
	Adress_detail        models.NullString `json:"adress_detail"`
	Adress_detail_kana   models.NullString `json:"adress_detail_kana"`
	Add_phone_number     models.NullString `json:"add_phone_number"`
	Marital_status       string            `json:"marital_status"`
	Dormitory_status     string            `json:"dormitory_status"`
}

type GeneralInformationGeneral struct {
	Id_general_information                        int    `json:"id_general_information"`
	Id_basic_information                          int    `json:"id_basic_information"`
	Id_store_code                                 int    `json:"id_store_code"`
	Id_department                                 int    `json:"id_department"`
	Id_store_section                              int    `json:"id_store_section"`
	Id_unit                                       int    `json:"id_unit"`
	Id_bank                                       int    `json:"id_bank"`
	Account_type                                  string `json:"account_type"`
	Account_number                                int    `json:"account_number"`
	Account_name                                  string `json:"account_name"`
	First_smester_in_other_company                string `json:"first_smester_in_other_company"`
	Distance_trip                                 string `json:"distance_trip"`
	Resume_document                               string `json:"resume_document"`
	Written_oath_document                         string `json:"written_oath_document"`
	Employee_agreement_document                   string `json:"employee_agreement_document"`
	Certificate_of_residence_card_document        string `json:"certificate_of_residence_card_document"`
	Application_form_of_commuting_method_document string `json:"application_form_of_commuting_method_document"`
	Compliance_agreement_document                 string `json:"compliance_agreement_document"`
	With_holding_slip_document                    string `json:"with_holding_slip_document"`
	Dependent_deduction_form_document             string `json:"dependent_deduction_form_document"`
	Pension_book_document                         string `json:"pension_book_document"`
	Health_check_report_document                  string `json:"health_check_report_document"`
	Office_code                                   string `json:"office_code"`
}

type GeneralRecrutmentStatusApproval struct {
	Id_general_recruitment_status_approval int    `json:"id_general_recruitment_status_approval"`
	Id_basic_information                   int    `json:"id_basic_information"`
	Status                                 string `json:"status"`
	Reason                                 string `json:"reason"`
	Date_time                              string `json:"date_time"`
	Date_time_approve                      string `json:"date_time_approve"`
	Data_check                             string `json:"data_check"`
	Date_time_data_check                   string `json:"date_time_data_check"`
	Message_for_edit                       string `json:"message_for_edit"`
	By_employee_code                       string `json:"by_employee_code"`
	Flag                                   string `json:"flag"`
	Real_data                              string `json:"real_data"`
}

type GeneralRecrutmentJoin struct {
	// start struct basic information
	Id_basic_information int               `json:"id_basic_information"`
	Employee_code        int               `json:"employee_code"`
	First_name           string            `json:"first_name"`
	Last_name            string            `json:"last_name"`
	Gender               string            `json:"gender"`
	Add_postal_code      models.NullString `json:"add_postal_code"`
	Id_prefecture        int               `json:"id_prefecture"`
	Adress               models.NullString `json:"addres"`
	Adress_kana          models.NullString `json:"adress_kana"`
	Adress_detail        models.NullString `json:"adress_detail"`
	Adress_detail_kana   models.NullString `json:"adress_detail_kana"`
	Add_phone_number     models.NullString `json:"add_phone_number"`
	Marital_status       string            `json:"marital_status"`
	Dormitory_status     string            `json:"dormitory_status"`
	// end struct basic information
	// start struct general information
	Id_general_information                        int    `json:"id_general_information"`
	Id_store_code                                 int    `json:"id_store_code"`
	Id_department                                 int    `json:"id_department"`
	Id_store_section                              int    `json:"id_store_section"`
	Id_unit                                       int    `json:"id_unit"`
	Id_bank                                       int    `json:"id_bank"`
	Account_type                                  string `json:"account_type"`
	Account_number                                int    `json:"account_number"`
	Account_name                                  string `json:"account_name"`
	First_smester_in_other_company                string `json:"first_smester_in_other_company"`
	Distance_trip                                 string `json:"distance_trip"`
	Resume_document                               string `json:"resume_document"`
	Written_oath_document                         string `json:"written_oath_document"`
	Employee_agreement_document                   string `json:"employee_agreement_document"`
	Certificate_of_residence_card_document        string `json:"certificate_of_residence_card_document"`
	Application_form_of_commuting_method_document string `json:"application_form_of_commuting_method_document"`
	Compliance_agreement_document                 string `json:"compliance_agreement_document"`
	With_holding_slip_document                    string `json:"with_holding_slip_document"`
	Dependent_deduction_form_document             string `json:"dependent_deduction_form_document"`
	Pension_book_document                         string `json:"pension_book_document"`
	Health_check_report_document                  string `json:"health_check_report_document"`
	Office_code                                   string `json:"office_code"`
	// end struct general information
	// start struct general recrutment statsu approval
	Id_general_recruitment_status_approval int    `json:"id_general_recruitment_status_approval"`
	Status                                 string `json:"status"`
	Reason                                 string `json:"reason"`
	Data_check                             string `json:"data_check"`
	Message_for_edit                       string `json:"message_for_edit"`
	By_employee_code                       string `json:"by_employee_code"`
	Flag                                   string `json:"flag"`
	Real_data                              string `json:"real_data"`
	// end struct general recrutment statsu approval

}
