package Commuting

import "../../models"

type Init_CommutingApprove struct {
	IdCommuting        int               `json:"id_commuting"`
	EmployeeNumber     string            `json:"employee_number"`
	CodeStore          models.NullString `json:"code_store"`
	DepartmentCode     models.NullString `json:"department_code"`
	StoreSection       models.NullString `json:"store_section"`
	UnitCode           models.NullString `json:"unit_code"`
	DivisiCodeAndName  models.NullString `json:"divisi_code_and_name"`
	AddressAndNumber   models.NullString `json:"address_and_number"`
	SumDistance        float64           `json:"sum_distance"`
	SumCost            int64             `json:"sum_cost"`
	CodeCommuting      string            `json:"code_commuting"`
	IdBasicInformation string            `json:"id_basic_information"`
}

type Init_DetailCommutingByEmployeeCode struct {
	NameEmployee          string `json:"name_employee"`
	StoreCode             string `json:"store_code"`
	ExpireDate            string `json:"expire_date"`
	InsuranseExpire       string `json:"insuranse_expire"`
	StatusExpireDate      string `json:"status_expire_date"`
	StatusInsuranseExpire string `json:"status_insuranse_expire"`
	DataDetail            string `json:"data_detail"` // data detail belum tau apa
	CountCost             string `json:"count_cost"`
	Distanse              string `json:"distanse"`
}

type Init_DetailCommutingByEmployeeCodeApprove struct {
	IdCommutingTrip       int               `json:"id_commuting_trip"`
	IdDetailCommutingTrip int               `json:"id_detail_commuting_trip"`
	RouteProfileName      models.NullString `json:"route_profile_name"`
	Date                  models.NullString `json:"date"`
	TypeOfTransport       string            `json:"type_of_transport"`
	AttendanceCode        models.NullString `json:"attendance_code"`
	Purpose               models.NullString `json:"purpose"`
	Distance              float64           `json:"distance"`
	CommuteDistance       float64           `json:"commute_distance"`
	Cost                  int64             `json:"cost"`
	Route                 string            `json:"route"`
	StatusCommuting       string            `json:"status_commuting"`
	DateApprove           models.NullString `json:"date_approve"`
}

type Init_DataCommutingByEmployeeCodeApprove struct {
	FirstName                     models.NullString `json:"first_name"`
	LastName                      models.NullString `json:"last_name"`
	CodeStore                     models.NullString `json:"code_store"`
	StoreName                     models.NullString `json:"store_name"`
	EmployeeCode                  string            `json:"employee_code"`
	DriverLicenseExpiryDate       models.NullString `json:"driver_license_expiry_date"`
	CarInsuranceExpiryDate        models.NullString `json:"car_insurance_expiry_date"`
	StatusDriverLicenseExpiryDate string            `json:"status_driver_license_expiry_date"`
	StatusCarInsuranceExpiryDate  string            `json:"status_car_insurance_expiry_date"`
}

type FormatDataDetailCommutingByEmployeeCode struct {
	Data       interface{} `json:"data"`
	DataDetail interface{} `json:"data_detail"`
}

type Init_InputDataApprove struct {
	IdCommuting       string `json:"id_commuting"`
	StatusDataApprove string `json:"status_data_approve"`
}
