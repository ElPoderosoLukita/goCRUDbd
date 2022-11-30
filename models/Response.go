package models

type Response struct {
	Status  int
	Message string
}

func NewResponse(status int, message string) Response {
	return Response{status, message}
}
