package initialize

import (
	"../models"
)

type ExpCategory struct {
	Id_exp        int               `json:"id_exp"`
	Exp_category  string            `json:"exp_category"`
	Created_date  models.NullString `json:"created_date"`
	Created_time  models.NullString `json:"created_time"`
	Code_category string            `json:"code_category"`
	Content       string            `json:"content"`
	Rule_code     models.NullString `json:"rule_code"`
}
