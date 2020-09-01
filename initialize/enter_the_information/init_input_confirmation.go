package enter_the_information

import 	models3 "../../models"


type IC_BasicInformation struct {
	FirstName                    string `json:"first_name"`
	LastName                     string `json:"last_name"`
	DriverLicenseExpirationDate  string `json:"driver_license_expiration_date"`
	ExpirationDateOfCarInsurance string `json:"expiration_date_of_car_insurance"`
}

type IC_DetailInformation struct {
	Destination     string `json:"destination"`
	UseDay          string `json:"use_day"`
	MeansOFMovement string `json:"means_of_movement"`
	AttendanceCode  string `json:"attendance_code"`
	Purpose         string `json:"purpose"`
	UsageRoute      string `json:"usage_route"`
}

type IC_Format struct {
	StatusDriversLicense string `json:"status_drivers_license"`
	StatusCarInsurance string `json:"status_car_insurance"`
	KodeBasicInformation models3.NullInt `json:"kode_basic_information"`
	DataBasic interface{} `json:"data_basic"`
	DataCommutingBasic interface{} `json:"data_commuting_basic"`
	DataDetail interface{} `json:"data_detail"`
}