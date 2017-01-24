// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package errors

// ErrorResponse represents class ErrorResponse
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_ErrorResponse
type ErrorResponse struct {
	ErrorID *string     `json:"errorId,omitempty"`
	Errors  *[]APIError `json:"errors,omitempty"`
}

// NewErrorResponse constructs a new ErrorResponse
func NewErrorResponse() *ErrorResponse {
	return &ErrorResponse{}
}
