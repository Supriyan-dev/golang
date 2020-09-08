package initialize

import "../models"

type Bank struct {
	Id_bank     int    `json:"id_bank"`
	Bank_code   string `json:"bank_code"`
	Bank_name   string `json:"bank_name"`
	Branch_code string `json:"branch_code"`
	Branch_name string `json:"branch_name"`
	Special     string `json:"special"`
}

type NullStringBank struct {
	Bank_name models.NullString `json:"bank_name"`
}
