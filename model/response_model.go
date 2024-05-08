package model 

type ServiceResponse struct {
	Code    int
	Error   bool
	Message string
	Data    interface{}
}

type Response struct {
	Error   bool
	Message string
	Data    any
}