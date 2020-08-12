package initialize

import (
	"database/sql"
)

type CommutingTrip struct {
	Id_commuting_trip      int            `json :"id_commuting_trip "`
	Id_general_information sql.NullInt64  `json :"id_general_information  "`
	Route_profile_name     sql.NullString `json :"route_profile_name"`
	Date                   sql.NullString `json :"date"`
	Attendance_code        sql.NullString `json :"attendance_code"`
	Save_draft_status      sql.NullString `json :"save_draft_status"`
	Status_approval        sql.NullString `json :"status_approval"`
	Reason                 sql.NullString `json :"reason"`
	Draft                  sql.NullString `json :"draft"`
	Save_trip              sql.NullString `json :"save_trip"`
	Submit                 sql.NullString `json :"submit"`
	Date_submit            sql.NullString `json :"date_submit"`
	Time_submit            sql.NullString `json :"time_submit"`
	Date_time_approve      sql.NullString `json :"date_time_approve"`
	Code_commuting         sql.NullString `json :"code_commuting "`
	Created_date           sql.NullString `json :"created_date"`
	Created_time           sql.NullString `json :"created_time"`
	Superior_verbal        sql.NullString `json :"superior_verbal"`
}
