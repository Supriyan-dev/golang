package initialize

type Response struct {
	Status      int         `json:"status"`
	Message     string      `json:"message"`
	TotalPage   int         `json:"totalPage"`
	CurrentPage int         `json:"currentPage"`
	Data        interface{} `json:"data"`
}

type ResponseWithPagination struct {
	Status      int         `json:"status"`
	Message     string      `json:"message"`
	CountData   int         `json:"count_data"`
	TotalPage   int         `json:"totalPage"`
	CurrentPage int         `json:"currentPage"`
	Data        interface{} `json:"data"`
}

type ResponseMaster struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
type NameTest struct {
	RouteProfileName string `json:"route_profile_name"`
	Date             string `json:"date"`
	TypeOfTransport  string `json:"type_of_transport"`
	AttendanceCode   string `json:"attendance_code"`
	Purpose          string `json:"purpose"`
}
