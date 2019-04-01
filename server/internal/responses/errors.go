package responses

import "net/http"

type ErrorResponse struct {
	StatusCode int
	Message    string
}

func NewInternalError(message string) *ErrorResponse {
	return &ErrorResponse{
		StatusCode: http.StatusInternalServerError,
		Message:    message,
	}
}

func NewNotFoundError(message string) *ErrorResponse {
	return &ErrorResponse{
		StatusCode: http.StatusNotFound,
		Message:    message,
	}
}
