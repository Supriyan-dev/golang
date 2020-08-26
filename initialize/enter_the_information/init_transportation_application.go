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

type InsertBasicInformation struct {
	InsuranceCompany string `json:"insurance_company"`
	DriverLicenseExpiryDate string `json:"driver_license_expiry_date"`
	PersonalInjury string `json:"personal_injury"`
	PropertyDamage string `json:"property_damage"`
	CarInsuranceDocumentExpiryDate string `json:"car_insurance_document_expiry_date"`
	IdGeneralInformation string `json:"id_general_information"`
}

type InsertTransportationApplication struct {
	RouteProfileName string `json:"route_profile_name"`
	Date string `json:"date"`
	Attendance string `json:"attendance"`
	CodeCommuting string `json:"code_commuting"`
	IdGeneralInformation string `json:"id_general_information"`
	DataDetail []InsertDetailTransportationApplication `json:"data_detail"`
}
type InsertDetailTransportationApplication struct {
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
