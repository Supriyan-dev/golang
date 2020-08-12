package initialize

import (
	"database/sql"
)

type CashClaim struct {
	Id_cash_claim                      int            `json :"id_cash_claim"`
	Code_store                         int            `json :"code_store"`
	Employee_code                      int            `json :"employee_code"`
	Division                           int            `json :"division"`
	Employee_name                      string         `json :"employee_name"`
	Description                        string         `json :"description"`
	Reason                             sql.NullString `json :"reason"`
	Amount                             string         `json :"amount"`
	Gbr_bill                           string         `json :"gbr_bill"`
	Exp_category                       int            `json :"exp_category"`
	Bsn_partner                        sql.NullString `json :"bsn_partner"`
	Pymt_method                        sql.NullInt64  `json :"pymt_method"`
	Ppl_joined                         sql.NullString `json :"ppl_joined"`
	Join_kasumi                        sql.NullString `json :"join_kasumi"`
	Created_at                         string         `json :"created_at"`
	Created_time                       string         `json :"created_time"`
	Submission_number                  string         `json :"submission_number"`
	Approve_level_1                    sql.NullString `json :"approve_level_1"`
	Approve_date_level_1               sql.NullString `json :"approve_date_level_1"`
	Approve_time_level_1               sql.NullString `json :"approve_time_level_1"`
	Approve_level_2                    sql.NullString `json :"approve_level_2"`
	Approve_date_level_2               sql.NullString `json :"approve_date_level_2"`
	Approve_time_level_2               sql.NullString `json :"approve_time_level_2"`
	Approve_level_3                    sql.NullString `json :"approve_level_3"`
	Approve_date_level_3               sql.NullString `json :"approve_date_level_3"`
	Approve_time_level_3               sql.NullString `json :"approve_time_level_3"`
	Reason_level_1                     sql.NullString `json :"reason_level_1"`
	Reason_level_2                     sql.NullString `json :"reason_level_2"`
	Reason_level_3                     sql.NullString `json :"reason_level_3"`
	Approve_code_1                     sql.NullString `json :"approve_code_1"`
	Approve_name_1                     sql.NullString `json :"approve_name_1"`
	Approve_code_2                     sql.NullString `json :"approve_code_2"`
	Approve_name_2                     sql.NullString `json :"approve_name_2"`
	Approve_code_3                     sql.NullString `json :"approve_code_3"`
	Approve_name_3                     sql.NullString `json :"approve_name_3"`
	Name_of_representative_from_client sql.NullString `json :"name_of_representative_from_client"`
	Number_persons_from_client         sql.NullString `json :"number_persons_from_client"`
	Purpose_joined                     sql.NullString `json :"purpose_joined"`
	Client_name                        sql.NullString `json :"client_name"`
}
