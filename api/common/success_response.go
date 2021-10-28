package common

import "net/http"

type SuccessResponseCode string
type SuccessWithoutDataResponseCode string

//List of success response status
const (
	Success            SuccessResponseCode            = "200"
	SuccessWithoutData SuccessWithoutDataResponseCode = "204"
)

//SuccessResponse default payload response
type SuccessResponse struct {
	Code    SuccessResponseCode `json:"code"`
	Message string              `json:"message"`
	Data    interface{}         `json:"data"`
}

//SuccessWithoutDataResponse default payload response without data
type SuccessWithoutDataResponse struct {
	Code    SuccessWithoutDataResponseCode `json:"code"`
	Message string                         `json:"message"`
	Data    interface{}                    `json:"data"`
}

//NewSuccessResponse create new success payload
func NewSuccessResponse(data interface{}) (int, SuccessResponse) {
	return http.StatusOK, SuccessResponse{
		Success,
		"Success",
		data,
	}
}

// NewSuccessResponseWithoutData create new success payload
func NewSuccessResponseWithoutData() (int, SuccessWithoutDataResponse) {
	return http.StatusNoContent, SuccessWithoutDataResponse{
		SuccessWithoutData,
		"Success",
		map[string]interface{}{},
	}
}
