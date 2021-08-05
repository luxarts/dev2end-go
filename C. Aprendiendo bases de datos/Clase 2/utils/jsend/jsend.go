package jsend

import "fmt"

const (
	StatusSuccess = "success"
	StatusError = "error"
	StatusFail = "fail"
)

// JSend specification
// Status: success, fail, error
// Message: Only exists if status == error
type ApiResponse struct {
	Status string `json:"status"`
	Message string `json:"message,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

func NewError(message string, data interface{}) *ApiResponse {
	return &ApiResponse{
		Status:  StatusError,
		Message: message,
		Data:    data,
	}
}
func NewFailure(data interface{}) *ApiResponse {
	return &ApiResponse{
		Status: StatusFail,
		Data: data,
	}
}
func NewResponse(data interface{}) *ApiResponse {
	return &ApiResponse{
		Status: StatusSuccess,
		Data: data,
	}
}

func (e *ApiResponse) Error() string {
	if e.Status == StatusError{
		return e.Message
	} else {
		return fmt.Sprintf("%v", e.Data)
	}
}