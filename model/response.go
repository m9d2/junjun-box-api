package model

type Response struct {
	Status     string      `json:"status"`
	HttpCode   int         `json:"httpCode"`
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}
