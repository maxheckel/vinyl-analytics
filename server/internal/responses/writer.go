package responses

import (
	"encoding/json"
	"net/http"
)

func WriteResponse(statusCode int, data interface{}, writer http.ResponseWriter){
	if data == nil {
		writer.WriteHeader(statusCode)
		return
	}
	responseJson, err := json.Marshal(data)
	if err != nil {
		NewInternalError("Could not marshal struct to json", writer)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
	_, err = writer.Write(responseJson)
	if err != nil {
		NewInternalError("Could not write response", writer)
		return
	}

}
