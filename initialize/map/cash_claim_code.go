package initialize

import (
	"database/sql"
)

type CashClaimCode struct {
	Id_code           int            `json :"id_code"`
	Submission_number string         `json :"submission_number "`
	Created_date      string         `json :"created_date"`
	Created_time      string         `json :"created_time"`
	File_csv          sql.NullString `json :"file_csv"`
	Submit_to_approve string         `json :"submit_to_approve"`
	Date_submit       string         `json :"date_submit"`
	Lock_by_1         sql.NullString `json :"lock_by_1"`
	Lock_by_2         sql.NullString `json :"lock_by_2"`
	Lock_by_3         sql.NullString `json :"lock_by_3"`
}
