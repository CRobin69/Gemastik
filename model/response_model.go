package model 

type ServiceResponse struct {
	Code    int
	Error   bool
	Message string
	Data    any
}

type Response struct {
	Error   bool
	Message string
	Data    any
}