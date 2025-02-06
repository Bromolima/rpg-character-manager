package customErrors

import (
	"net/http"
)

type CustomError struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int      `json:"code"`
	Causes  []Causes `json:"causes,omitempty"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (r *CustomError) Error() string {
	return r.Message
}

func NewCustomErr(message, err string, code int, causes []Causes) *CustomError {
	return &CustomError{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

func NewBadRequestErr(message string) *CustomError {
	return &CustomError{
		Message: message,
		Err:     "bad request",
		Code:    http.StatusBadRequest,
	}
}

func NewBadRequestValidatorErr(message string, causes []Causes) *CustomError {
	return &CustomError{
		Message: message,
		Err:     "bad request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewInternalServerErr(message string) *CustomError {
	return &CustomError{
		Message: message,
		Err:     "internal server error",
		Code:    http.StatusInternalServerError,
	}
}
