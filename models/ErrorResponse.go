package models

type ErrorResponse struct {
	Status int
	Error  string
}

func NewErrorStruct(status int, errorMessage string) ErrorResponse {
	return ErrorResponse{status, errorMessage}
}
