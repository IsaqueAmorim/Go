package rest_error

import "net/http"

type RestError struct {
	Message string   `json:"message"`
	Error   string   `json:"err"`
	Code    int      `json:"code"`
	Causes  []Causes `json:"causes"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (r *RestError) NewError() string {
	return r.Message
}

func NewRestErr(message, error string, code int, causes []Causes) *RestError {
	return &RestError{
		Message: message,
		Error:   error,
		Code:    code,
		Causes:  causes,
	}
}

func NewBadRequestError(message string) *RestError {
	return &RestError{
		Message: message,
		Error:   "bad_request",
		Code:    http.StatusBadRequest,
	}
}

func NewBadRequestValidationError(message string, causes []Causes) *RestError {
	return &RestError{
		Message: message,
		Error:   "bad_request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewInternalServerError(message string) *RestError {
	return &RestError{
		Message: message,
		Error:   "internal_server_error",
		Code:    http.StatusInternalServerError,
	}
}

func NewNotFoundError(message string) *RestError {
	return &RestError{
		Message: message,
		Error:   "not_found",
		Code:    http.StatusNotFound,
	}
}

func NewForbiddenError(message string) *RestError {
	return &RestError{
		Message: message,
		Error:   "forbiden",
		Code:    http.StatusForbidden,
	}
}
