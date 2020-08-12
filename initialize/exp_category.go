package initialize

import "database/sql"

type ExpCategory struct {
	Id_exp        int            `json:"id_exp "`
	Exp_category  string         `json:"exp_category"`
	Created_date  sql.NullString `json:"created_date"`
	Created_time  sql.NullString `json:"created_time"`
	Code_category string         `json:"code_category "`
	Content       string         `json:"content"`
	Rule_code     string         `json:"rule_code "`
}
