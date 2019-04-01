package responses

import "net/http"

type ErrorResponse struct {
	StatusCode int
	Message    string
}

func NewInternalError(message string, writer http.ResponseWriter) {
	response :=  &ErrorResponse{
		StatusCode: http.StatusInternalServerError,
		Message:    message,
	}
	WriteResponse(response.StatusCode, response, writer)
}

func NewNotFoundError(message string, writer http.ResponseWriter) {
	response :=  &ErrorResponse{
		StatusCode: http.StatusNotFound,
		Message:    message,
	}
	WriteResponse(response.StatusCode, response, writer)
}

func NewBadRequestErrror(message string, writer http.ResponseWriter) {
	response :=  &ErrorResponse{
		StatusCode: http.StatusBadRequest,
		Message: message,
	}
	WriteResponse(response.StatusCode, response, writer)
}

func NewIntegrationError(message string, writer http.ResponseWriter){
	response :=  &ErrorResponse{
		StatusCode: http.StatusFailedDependency,
		Message:message,
	}
	WriteResponse(response.StatusCode, response, writer)
}