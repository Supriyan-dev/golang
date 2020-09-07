package initialize

import "database/sql"

type Join struct {
	Id_store_code                      int            `json:"id_store_code"`
	Employee_code                      int            `json:"employee_code"`
	First_name                         string         `json:"first_name"`
	Last_name                          string         `json:"last_name"`
	Driver_license_expiry_date         sql.NullString `json:"driver_license_expiry_date"`
	Car_insurance_document_expiry_date sql.NullString `json:"car_insurance_document_expiry_date"`
}

type UpdatePermissionToDrive struct {
	Id_commuting_basic_information int    `json:"id_commuting_basic_information"`
	Permitted_to_drive             string `json:"permitted_to_drive"`
	Status_approve                 string `json:"status_approve"`
}
