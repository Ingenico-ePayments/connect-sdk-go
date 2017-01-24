// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package errors

// APIError represents class APIError
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_APIError
type APIError struct {
	Code           *string `json:"code,omitempty"`
	HTTPStatusCode *int32  `json:"httpStatusCode,omitempty"`
	Message        *string `json:"message,omitempty"`
	PropertyName   *string `json:"propertyName,omitempty"`
	RequestID      *string `json:"requestId,omitempty"`
}

// NewAPIError constructs a new APIError
func NewAPIError() *APIError {
	return &APIError{}
}
