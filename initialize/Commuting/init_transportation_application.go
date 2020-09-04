package Commuting

import (
	models3 "../../models"
)

//type ShowDetailTransportationApplication struct {
//	RouteProfileName string `json:"route_profile_name"`
//	Date string `json:"date"`
//	TypeOfTransport string `json:"type_of_transport"`
//	AttendanceCode string `json:"attendance_code"`
//	Purpose string `json:"purpose"`
//}

type ShowDetailTransportationApplicationGet struct {
	IdDetailCommutingTrip models3.NullInt    `json:"id_detail_commuting_trip"`
	IdCommutingTrip       models3.NullInt    `json:"id_commuting_trip"`
	TypeOfTransport       models3.NullString `json:"type_of_transport"`
	Purpose               models3.NullString `json:"purpose"`
	DetailFrom            models3.NullString `json:"detail_from"`
	DetailTo              models3.NullString `json:"detail_to"`
	Distance              models3.NullFloat  `json:"distance"`
	Cost                  models3.NullInt    `json:"cost"`
	PointTrip             models3.NullFloat  `json:"point_trip"`
	TransitPoint          models3.NullString `json:"transit_point"`
	CommuteDistance       models3.NullFloat  `json:"commute_distance"`
	GoOutDistance         models3.NullFloat  `json:"go_out_distance"`
}

type FormatShowBasicInformation struct {
	DataBasicInformation interface{} `json:"data_basic_information"`
	DataApprove          interface{} `json:"data_approve"`
	DataApply            interface{} `json:"data_apply"`
}

type ShowBasicInformation1 struct {
	IdBasicInformation int                `json:"id_basic_information"`
	FirstName          models3.NullString `json:"first_name"`
	LastName           models3.NullString `json:"last_name"`
	Address            models3.NullString `json:"address"`
	AddressKana        models3.NullString `json:"address_kana"`
	AddressDetail      models3.NullString `json:"address_detail"`
	AddressDetailKana  models3.NullString `json:"address_detail_kana"`
	AddPhoneNumber     models3.NullString `json:"add_phone_number"`
}

type ShowBasicInformation2 struct {
	IdGeneralBasicInformation int                `json:"id_commuting_basic_information"`
	StatusApproved            models3.NullString `json:"status_approved"`
	DateApprove               models3.NullString `json:"date_approve"`
	TimeApprove               models3.NullString `json:"time_approve"`
	DateSubmit                models3.NullString `json:"date_submit"`
}

type ShowBasicInformation3 struct {
	IdCommutingBasicInformation    int                `json:"id_commuting_basic_information"`
	IdGeneralInformation           int                `json:"id_general_information"`
	InsuranceCompany               models3.NullString `json:"insurance_company"`
	DriverLicenseExpiryDate        models3.NullString `json:"driver_license_expiry_date"`
	PersonalInjury                 models3.NullString `json:"personal_injury"`
	PropertyDamage                 models3.NullString `json:"property_damage"`
	CarInsuranceDocumentExpiryDate models3.NullString `json:"car_insurance_document_expiry_date"`
}

type FormatShowUsageRecord struct {
	CountHistory         int             `json:"count_history"`
	KodeBasicInformation models3.NullInt `json:"kode_basic_information"`
	DataBasicInformation interface{}     `json:"data_basic_information"`
	DataUsageRecord      interface{}     `json:"data_usage_record"`
}

type FormatShowUsageRecordForEdit struct {
	DataTrip       interface{} `json:"data_trip"`
	DetailDataTrip interface{} `json:"detail_data_trip"`
}

type ShowUsageRecord struct {
	IdDetailCommutingTrip int                `json:"id_detail_commuting_trip"`
	IdCommutingTrip       int                `json:"id_commuting_trip"`
	TypeOfTransport       models3.NullString `json:"type_of_transport"`
	Purpose               models3.NullString `json:"purpose"`
	DetailFrom            models3.NullString `json:"detail_from"`
	DetailTo              models3.NullString `json:"detail_to"`
	Distance              models3.NullFloat  `json:"distance"`
	Cost                  models3.NullInt    `json:"cost"`
	PointTrip             models3.NullFloat  `json:"point_trip"`
	TransitPoint          models3.NullString `json:"transit_point"`
	CommuteDistance       models3.NullFloat  `json:"commute_distance"`
	GoOutDistance         models3.NullFloat  `json:"go_out_distance"`
}

type ShowUsageRecord2 struct {
	//IdDetailCommutingTrip int       `json:"id_detail_commuting_trip"`
	IdCommutingTrip  int                `json:"id_commuting_trip"`
	RouteProfileName models3.NullString `json:"route_profile_name"`
	Date             models3.NullString `json:"date"`
	TypeOfTransport  string             `json:"type_of_transport"`
	Purpose          string             `json:"purpose"`
	Route            string             `json:"route"`
	Distance         models3.NullFloat  `json:"distance"`
	CommuteDistance  models3.NullFloat  `json:"commute_distance"`
	Cost             models3.NullInt    `json:"cost"`
	StatusTemporary  string             `json:"status_temporary"`
	//PointTrip             models3.NullFloat `json:"point_trip"`
	//TransitPoint          models3.NullString  `json:"transit_point"`
	//GoOutDistance         models3.NullFloat `json:"go_out_distance"`
}

type ShowCommutingTrip struct {
	IdCommutingTrip  int    `json:"id_commuting_trip"`
	Date             string `json:"date"`
	RouteProfileName string `json:"route_profile_name"`
	AttendanceCode   string `json:"attendance_code"`
}

type ShowUseMyRoute struct {
	IdCommutingTrip       int                `json:"id_commuting_trip"`
	IdDetailCommutingTrip int                `json:"id_detail_commuting_trip"`
	RouteProfileName      models3.NullString `json:"route_profile_name"`
	TypeOfTransport       string             `json:"type_of_transport"`
	AttendanceCode        models3.NullString `json:"attendance_code"`
	Purpose               string             `json:"purpose"`
	Route                 string             `json:"route"`
	Distance              float64            `json:"distance"`
	CommuteDistance       models3.NullFloat  `json:"commute_distance"`
	Cost                  models3.NullInt    `json:"cost"`
}

type InsertBasicInformation struct {
	IdCommutingBasicInformation string `json:"id_commuting_basic_information"`
	InsuranceCompany               string `json:"insurance_company"`
	DriverLicenseExpiryDate        string `json:"driver_license_expiry_date"`
	PersonalInjury                 string `json:"personal_injury"`
	PropertyDamage                 string `json:"property_damage"`
	CarInsuranceDocumentExpiryDate string `json:"car_insurance_document_expiry_date"`
	IdGeneralInformation           string `json:"id_general_information"`
}

type InsertUsageRecordApplyForTravelExpenses struct {
	RouteProfileName     string                                          `json:"route_profile_name"`
	Date                 string                                          `json:"date"`
	Attendance           string                                          `json:"attendance"`
	CodeCommuting        int                                             `json:"code_commuting"`
	IdGeneralInformation string                                          `json:"id_general_information"`
	DataDetail           []InsertDetailUsageRecordApplyForTravelExpenses `json:"data_detail"`
}
type InsertDetailUsageRecordApplyForTravelExpenses struct {
	IdCommutingTrip int     `json:"id_commuting_trip"`
	TypeOfTransport string  `json:"type_of_transport"`
	Purpose         string  `json:"purpose"`
	DetailFrom      string  `json:"detail_from"`
	DetailTo        string  `json:"detail_to"`
	Distance        float64 `json:"distance"`
	Cost            int     `json:"cost"`
	PointTrip       float64 `json:"point_trip"`
	TransitPoint    string  `json:"transit_point"`
	CommuteDistance float64 `json:"commute_distance"`
	GoOutDistance   float64 `json:"go_out_distance"`
}

// indonesia
// digunakan untuk inisialisasi pada bagian Usage Record ConfirmationOfSubmissionContents

// english
// used to initialize Usage Record on 'Confirmation Of Submission Contents' / '提出内容の確認'

type UpdateUsageRecordApplyForTravelExpenses struct {
	IdCommutingTrip      string                                          `json:"id_commuting_trip"`
	RouteProfileName     string                                          `json:"route_profile_name"`
	Date                 string                                          `json:"date"`
	Attendance           string                                          `json:"attendance"`
	IdGeneralInformation string                                          `json:"id_general_information"`
	DataDetail           []UpdateDetailUsageRecordApplyForTravelExpenses `json:"data_detail"`
}
type UpdateDetailUsageRecordApplyForTravelExpenses struct {
	IdCommutingTripDetail int     `json:"id_commuting_trip_detail"`
	IdCommutingTrip       int     `json:"id_commuting_trip"`
	TypeOfTransport       string  `json:"type_of_transport"`
	Purpose               string  `json:"purpose"`
	DetailFrom            string  `json:"detail_from"`
	DetailTo              string  `json:"detail_to"`
	Distance              string  `json:"distance"`
	Cost                  int     `json:"cost"`
	PointTrip             float64 `json:"point_trip"`
	TransitPoint          string  `json:"transit_point"`
	CommuteDistance       float64 `json:"commute_distance"`
	GoOutDistance         float64 `json:"go_out_distance"`
}

type FormatHistory struct {
	DataCount   interface{} `json:"data_count"`
	Datahistory interface{} `json:"datahistory"`
}

type ShowAdditionalHistory struct {
	CountDataSubmit      int `json:"count_data_submit"`
	CountDataDraft       int `json:"count_data_draft"`
	CountDataPartial     int `json:"count_data_partial"`
	CountDataNotApproved int `json:"count_data_not_approved"`
}

type ShowHistory struct {
	IdDetailCommutingTrip int                `json:"id_detail_commuting_trip"`
	IdCommutingTrip       int                `json:"id_commuting_trip"`
	RouteProfileName      models3.NullString `json:"route_profile_name"`
	Date                  models3.NullString `json:"date"`
	TypeOfTransport       string             `json:"type_of_transport"`
	AttendanceCode        models3.NullString `json:"attendance_code"`
	Purpose               models3.NullString `json:"purpose"`
	Distance              float64            `json:"distance"`
	CommuteDistance       float64            `json:"commute_distance"`
	Cost                  int64              `json:"cost"`
	Route                 string             `json:"route"`
	StatusCommuting       string             `json:"status_commuting"`
	DateApprove           models3.NullString `json:"date_approve"`
}
