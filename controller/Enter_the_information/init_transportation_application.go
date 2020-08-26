package enter_the_information

//type ShowDetailTransportationApplication struct {
//	RouteProfileName string `json:"route_profile_name"`
//	Date string `json:"date"`
//	TypeOfTransport string `json:"type_of_transport"`
//	AttendanceCode string `json:"attendance_code"`
//	Purpose string `json:"purpose"`
//}

type ShowDetailTransportationApplication struct {
	IdDetailCommutingTrip int `json:"id_detail_commuting_trip"`
	IdCommutingTrip int `json:"id_commuting_trip"`
	TypeOfTransport string `json:"type_of_transport"`
	Purpose string `json:"purpose"`
	DetailFrom string `json:"detail_from"`
	DetailTo string `json:"detail_to"`
	Distance float64 `json:"distance"`
	Cost int `json:"cost"`
	PointTrip float64 `json:"point_trip"`
	TransitPoint string `json:"transit_point"`
	CommuteDistance float64 `json:"commute_distance"`
	GoOutDistance float64 `json:"go_out_distance"`
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
