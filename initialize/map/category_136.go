package initialize

import "database/sql"

type Category_136 struct {
	Id_data       int            `json :"id_data "`
	Buy_name      string         `json :"buy_name"`
	Purpose       sql.NullString `json :"purpose"`
	Created_date  string         `json :"created_date"`
	Created_time  string         `json :"created_time  "`
	Id_cash_claim int            `json :"id_cash_claim   "`
}
