package responses

import (
	"encoding/json"
	"net/http"
)

func WriteResponse(statusCode int, data interface{}, writer http.ResponseWriter){
	responseJson, err := json.Marshal(data)
	if err != nil {
		response := NewInternalError("Could not marshal struct to json")
		WriteResponse(response.StatusCode, response, writer)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
	_, err = writer.Write(responseJson)
	if err != nil {
		response := NewInternalError("Could not write response")
		WriteResponse(response.StatusCode, response, writer)
		return
	}

}
