package initialize

type CodeCommuting struct {
	Id_code          int     `json:"id_code"`
	Code_random      string  `json:"code_random"`
	Std_deviation    float64 `json:"std_deviation"`
	Created_time     string  `json:"created_time"`
	Created_date     string  `json:"created_date"`
	Status_commuting string  `json:"status_commuting"`
}
