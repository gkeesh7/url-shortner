package api

import (
	"encoding/json"
	"net/http"
	"url-shortner/api/dto"
)

func GenerateResponseByErrorCode(response dto.Error) (int, []byte) {
	return outputJSON(dto.MapErrorCodeToStatusCode[response.Code], response)
}

func GenerateSuccessResponse(response interface{}) (int, []byte) {
	return outputJSON(http.StatusOK, response)
}

func outputJSON(respCode int, payload interface{}) (int, []byte) {
	output, err := json.Marshal(payload)
	if err != nil {
		output, _ = json.Marshal(payload)
		return respCode, output
	}
	return respCode, output
}

func ErrorHandler(w http.ResponseWriter, errorDto dto.Error) {
	//TODO: Send a Custom Error Html Page
	statusCode, jsonResponse := GenerateResponseByErrorCode(errorDto)
	w.WriteHeader(statusCode)
	w.Write(jsonResponse)
}

func SuccessHandler(w http.ResponseWriter, successResponse interface{}) {
	statusCode, jsonResponse := GenerateSuccessResponse(successResponse)
	w.WriteHeader(statusCode)
	w.Write(jsonResponse)
}
