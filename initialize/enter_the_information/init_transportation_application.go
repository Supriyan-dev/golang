package enter_the_information

import (
	models "../../models/map"
	models2 "Go_DX_Services/models/map"
)

//type ShowDetailTransportationApplication struct {
//	RouteProfileName string `json:"route_profile_name"`
//	Date string `json:"date"`
//	TypeOfTransport string `json:"type_of_transport"`
//	AttendanceCode string `json:"attendance_code"`
//	Purpose string `json:"purpose"`
//}

type ShowDetailTransportationApplicationGet struct {
	IdDetailCommutingTrip models.NullInt64 `json:"id_detail_commuting_trip"`
	IdCommutingTrip models.NullInt64 `json:"id_commuting_trip"`
	TypeOfTransport models.NullString `json:"type_of_transport"`
	Purpose models.NullString `json:"purpose"`
	DetailFrom models.NullString `json:"detail_from"`
	DetailTo models.NullString `json:"detail_to"`
	Distance models.NullFloat64 `json:"distance"`
	Cost models.NullInt64 `json:"cost"`
	PointTrip models.NullFloat64 `json:"point_trip"`
	TransitPoint models.NullString `json:"transit_point"`
	CommuteDistance models.NullFloat64 `json:"commute_distance"`
	GoOutDistance models.NullFloat64 `json:"go_out_distance"`
}

type ShowBasicInformation struct {
	IdCommutingBasicInformation int `json:"id_commuting_basic_information"`
	IdGeneralInformation int `json:"id_general_information"`
	InsuranceCompany models2.NullString `json:"insurance_company"`
	DriverLicenseExpiryDate models2.NullString `json:"driver_license_expiry_date"`
	PersonalInjury models2.NullString `json:"personal_injury"`
	PropertyDamage models2.NullString `json:"property_damage"`
	CarInsuranceDocumentExpiryDate models2.NullString `json:"car_insurance_document_expiry_date"`
}

type ShowUsageRecord struct {
	IdDetailCommutingTrip int `json:"id_detail_commuting_trip"`
	IdCommutingTrip int `json:"id_commuting_trip"`
	TypeOfTransport models2.NullString `json:"type_of_transport"`
	Purpose models2.NullString `json:"purpose"`
	DetailFrom models2.NullString `json:"detail_from"`
	DetailTo models2.NullString `json:"detail_to"`
	Distance models2.NullFloat64 `json:"distance"`
	Cost models2.NullInt64 `json:"cost"`
	PointTrip models2.NullFloat64 `json:"point_trip"`
	TransitPoint models2.NullString `json:"transit_point"`
	CommuteDistance models2.NullFloat64 `json:"commute_distance"`
	GoOutDistance models2.NullFloat64 `json:"go_out_distance"`
}

type ShowUseMyRoute struct {
	IdCommutingTrip int `json:"id_commuting_trip"`
	IdDetailCommutingTrip int `json:"id_detail_commuting_trip"`
	RouteProfileName models2.NullString `json:"route_profile_name"`
	TypeOfTransport models2.NullString `json:"type_of_transport"`
	AttendanceCode models2.NullString `json:"attendance_code"`
	Purpose models2.NullString `json:"purpose"`
	DetailTo models2.NullString `json:"detail_to"`
	DetailFrom models2.NullString `json:"detail_from"`
	CommuteDistance models2.NullFloat64 `json:"commute_distance"`
	GoOutDistance models2.NullFloat64 `json:"go_out_distance"`
	Cost models2.NullInt64 `json:"cost"`
}

type InsertBasicInformation struct {
	InsuranceCompany string `json:"insurance_company"`
	DriverLicenseExpiryDate string `json:"driver_license_expiry_date"`
	PersonalInjury string `json:"personal_injury"`
	PropertyDamage string `json:"property_damage"`
	CarInsuranceDocumentExpiryDate string `json:"car_insurance_document_expiry_date"`
	IdGeneralInformation string `json:"id_general_information"`
}

type InsertUsageRecordApplyForTravelExpenses struct {
	RouteProfileName string `json:"route_profile_name"`
	Date string `json:"date"`
	Attendance string `json:"attendance"`
	CodeCommuting string `json:"code_commuting"`
	IdGeneralInformation string `json:"id_general_information"`
	DataDetail []InsertDetailUsageRecordApplyForTravelExpenses `json:"data_detail"`
}
type InsertDetailUsageRecordApplyForTravelExpenses struct {
	IdCommutingTrip int `json:"id_commuting_trip"`
	TypeOfTransport string `json:"type_of_transport"`
	Purpose string `json:"purpose"`
	DetailFrom string `json:"detail_from"`
	DetailTo string `json:"detail_to"`
	Distance string `json:"distance"`
	Cost int `json:"cost"`
	PointTrip float64 `json:"point_trip"`
	TransitPoint string `json:"transit_point"`
	CommuteDistance float64 `json:"commute_distance"`
	GoOutDistance float64 `json:"go_out_distance"`
}

// indonesia
// digunakan untuk inisialisasi pada bagian Usage Record ConfirmationOfSubmissionContents

// english
// used to initialize Usage Record on 'Confirmation Of Submission Contents' / '提出内容の確認'

type UpdateUsageRecordApplyForTravelExpenses struct {
	IdCommutingTrip string `json:"id_commuting_trip"`
	RouteProfileName string `json:"route_profile_name"`
	Date string `json:"date"`
	Attendance string `json:"attendance"`
	CodeCommuting string `json:"code_commuting"`
	IdGeneralInformation string `json:"id_general_information"`
	DataDetail []UpdateDetailUsageRecordApplyForTravelExpenses `json:"data_detail"`
}
type UpdateDetailUsageRecordApplyForTravelExpenses struct {
	IdCommutingTripDetail int `json:"id_commuting_trip_detail"`
	IdCommutingTrip int `json:"id_commuting_trip"`
	TypeOfTransport string `json:"type_of_transport"`
	Purpose string `json:"purpose"`
	DetailFrom string `json:"detail_from"`
	DetailTo string `json:"detail_to"`
	Distance string `json:"distance"`
	Cost int `json:"cost"`
	PointTrip float64 `json:"point_trip"`
	TransitPoint string `json:"transit_point"`
	CommuteDistance float64 `json:"commute_distance"`
	GoOutDistance float64 `json:"go_out_distance"`
}
