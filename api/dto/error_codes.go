package dto

import (
	"fmt"
	"net/http"
)

type ErrorCode int

const (
	IllegalRequest      ErrorCode = 995
	JsonError           ErrorCode = 996
	RecordNotFound      ErrorCode = 997
	DatabaseError       ErrorCode = 998
	InternalServerError ErrorCode = 999
)

var MapErrorCodeToStatusCode = map[ErrorCode]int{
	InternalServerError: http.StatusInternalServerError,
	DatabaseError:       http.StatusInternalServerError,
	RecordNotFound:      http.StatusNotFound,
	JsonError:           http.StatusInternalServerError,
	IllegalRequest:      http.StatusBadRequest,
}

var MapErrorCodeToErrorMessageTemplate = map[ErrorCode]string{
	InternalServerError: "Internal Server Error %v",
	DatabaseError:       "Facing issues with our DataStore %v",
	RecordNotFound:      "The requested data was not found %v",
	JsonError:           "Issues while encoding/decoding json %v",
	IllegalRequest:      "Some required fields are missing from the object %v",
}

func GenerateError(code ErrorCode, args ...string) *Error {
	return &Error{
		Code:    code,
		Message: fmt.Sprintf(MapErrorCodeToErrorMessageTemplate[code], args),
	}
}
