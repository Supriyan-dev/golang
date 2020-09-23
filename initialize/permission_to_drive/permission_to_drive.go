package initialize

import "../../models"

type Join struct {
	Id_code_store                      int               `json:"id_code_store"`
	Code_store                         string            `json:"code_store"`
	Employee_code                      string            `json:"employee_code"`
	First_name                         string            `json:"first_name"`
	Last_name                          string            `json:"last_name"`
	Driver_license_expiry_date         models.NullString `json:"driver_license_expiry_date"`
	Car_insurance_document_expiry_date models.NullString `json:"car_insurance_document_expiry_date"`
	Insurance_company models.NullString `json:"insurance_company"`
	Personal_injury	 models.NullString `json:"personal_injury"`
	Property_damage models.NullString `json:"property_damage"`
	Status_approve string`json:"status_approve"`
}

type UpdatePermissionToDrive struct {
	Id_commuting_basic_information int    `json:"id_commuting_basic_information"`
	Permitted_to_drive             string `json:"permitted_to_drive"`
	Status_approve                 string `json:"status_approve"`
}
