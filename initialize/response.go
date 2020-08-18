package initialize

type Response struct {
	Status      int         `json:"status"`
	Message     string      `json:"message"`
	TotalPage   int         `json:"totalPage"`
	CurrentPage int         `json:"currentPage"`
	Data        interface{} `json:"data"`
}

type Pagination struct {
}
