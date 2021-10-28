package common

import "net/http"

type errorControllerResponseCode string

const (
	ErrBadRequest errorControllerResponseCode = "bad_request"
	ErrForbidden  errorControllerResponseCode = "forbidden"
	ErrConflict   errorControllerResponseCode = "conflict"
)

//ControllerResponse default payload response
type ControllerResponse struct {
	Code    errorControllerResponseCode `json:"code"`
	Message string                      `json:"message"`
	Data    interface{}                 `json:"data"`
}

//NewBadRequestResponse bad request format response
func NewBadRequestResponse() (int, ControllerResponse) {
	return http.StatusBadRequest, ControllerResponse{
		ErrBadRequest,
		"Bad request",
		map[string]interface{}{},
	}
}

//NewForbiddenResponse default for Forbidden error response
func NewForbiddenResponse() (int, ControllerResponse) {
	return http.StatusForbidden, ControllerResponse{
		ErrForbidden,
		"Forbidden",
		map[string]interface{}{},
	}
}

//NewConflictResponse default for Conflict error response
func NewConflictResponse() (int, ControllerResponse) {
	return http.StatusConflict, ControllerResponse{
		ErrConflict,
		"conflict, data may be already exists",
		map[string]interface{}{},
	}
}
