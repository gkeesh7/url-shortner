package dto

// Error defines the custom error
type Error struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
}

// NewError returns a new Error
func NewError(code ErrorCode, msg string) *Error {
	return &Error{code, msg}
}

// Error returns the error message associated with the error.
func (err *Error) Error() string {
	return string(err.Code) + ":" + err.Message
}
