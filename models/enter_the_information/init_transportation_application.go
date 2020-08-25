package enter_the_information

type ShowDetailTransportationApplication struct {
	RouteProfileName string `json:"route_profile_name"`
	Date string `json:"date"`
	TypeOfTransport string `json:"type_of_transport"`
	AttendanceCode string `json:"attendance_code"`
	Purpose string `json:"purpose"`
}
type InsertTransportationApplication struct {
	RouteProfileName string `json:"route_profile_name"`
	Date string `json:"date"`
	Attendance string `json:"attendance"`
	CodeCommuting string `json:"code_commuting"`
}
type InsertDetailTransportationApplication struct {
	TypeOfTransport string `json:"type_of_transport"`
	Purpose string `json:"purpose"`
	DetailFrom string `json:"detail_from"`
	DetailTo string `json:"detail_to"`
	Distance string `json:"distance"`
	Cost string `json:"cost"`
	PointTrip string `json:"point_trip"`
	TransitPoint string `json:"transit_point"`
	CommuteDistance string `json:"commute_distance"`
	GoOutDistance string `json:"go_out_distance"`
}
