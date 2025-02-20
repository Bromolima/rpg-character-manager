package apiErrors

import (
	"net/http"
)

type ApiError struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int      `json:"code"`
	Causes  []Causes `json:"causes"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (r *ApiError) Error() string {
	return r.Message
}

func NewApiError(message, err string, code int, causes []Causes) *ApiError {
	return &ApiError{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

func NewBadRequestErr(message string) *ApiError {
	return &ApiError{
		Message: message,
		Err:     "bad request",
		Code:    http.StatusBadRequest,
	}
}

func NewInternalServerErr(message string) *ApiError {
	return &ApiError{
		Message: message,
		Err:     "internal server error",
		Code:    http.StatusInternalServerError,
	}
}

func NewBadRequestValidationErr(message string, causes []Causes) *ApiError {
	return &ApiError{
		Message: message,
		Err:     "bad request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}
