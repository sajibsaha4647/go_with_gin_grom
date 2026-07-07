package domain

type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Data    any    `json:"data"`
}

type ResponseLogin struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Data    any    `json:"data"`
	Token   string `json:token`
}
