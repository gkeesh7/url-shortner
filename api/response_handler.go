package api

import (
	"encoding/json"
	"net/http"
	"url-shortner/api/dto"
)

// GenerateResponseByErrorCode creates the error Json payload and appends the appropriate HTTP status code from the predefined map
func GenerateResponseByErrorCode(response dto.Error) (int, []byte) {
	return outputJSON(dto.MapErrorCodeToStatusCode[response.Code], response)
}

// GenerateSuccessResponse creates the Json payload and appends the 200 Ok status
func GenerateSuccessResponse(response interface{}) (int, []byte) {
	return outputJSON(http.StatusOK, response)
}

// outputJSON converts a payload into a byte array
func outputJSON(respCode int, payload interface{}) (int, []byte) {
	output, err := json.Marshal(payload)
	if err != nil {
		output, _ = json.Marshal(payload)
		return respCode, output
	}
	return respCode, output
}

// ErrorHandler returns an Error Json as response
func ErrorHandler(w http.ResponseWriter, errorDto dto.Error) {
	//TODO: Send a Custom Error Html Page in case of 404 401 etc
	statusCode, jsonResponse := GenerateResponseByErrorCode(errorDto)
	w.WriteHeader(statusCode)
	w.Write(jsonResponse)
}

// SuccessHandler returns Success Json as Response
func SuccessHandler(w http.ResponseWriter, successResponse interface{}) {
	statusCode, jsonResponse := GenerateSuccessResponse(successResponse)
	w.WriteHeader(statusCode)
	w.Write(jsonResponse)
}
