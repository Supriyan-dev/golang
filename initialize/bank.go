package initialize

import "database/sql"

type Bank struct {
	Id_bank     int            `json:"id_bank"`
	Bank_code   string         `json:"bank_code"`
	Bank_name   sql.NullString `json:"bank_name"`
	Branch_code string         `json:"branch_code"`
	Branch_name string         `json:"branch_name"`
	Special     bool           `json:"special"`
}
