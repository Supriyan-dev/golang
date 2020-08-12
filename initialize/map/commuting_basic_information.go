package initialize

import (
	"database/sql"
)

type CommutingBasicInformation struct {
	Id_commuting_basic_information     int            `json :"id_commuting_basic_information"`
	Id_general_information             int            `json :"id_general_information "`
	Driver_license_document            sql.NullString `json :"driver_license_document"`
	Driver_license_document_url        sql.NullString `json :"driver_license_document_url"`
	Driver_license_expiry_date         sql.NullString `json :"driver_license_expiry_date"`
	Car_insurance_document             sql.NullString `json :"car_insurance_document"`
	Car_insurance_document_url         sql.NullString `json :"car_insurance_document_url"`
	Car_insurance_document_expiry_date sql.NullString `json :"car_insurance_document_expiry_date"`
	Daily_commuting_method             sql.NullString `json :"daily_commuting_method"`
	Default_transportation             sql.NullString `json :"default_transportation"`
	Permitted_to_drive                 sql.NullString `json :"permitted_to_drive"`
	Insurance_company                  sql.NullString `json :"insurance_company"`
	Personal_injury                    sql.NullString `json :"personal_injury"`
	Property_damage                    sql.NullString `json :"property_damage"`
	Status_approve                     sql.NullString `json :"status_approve"`
	Date_approve                       sql.NullString `json :"date_approve"`
	Time_approve                       sql.NullString `json :"time_approve"`
	Date_submit                        sql.NullString `json :"date_submit"`
}
