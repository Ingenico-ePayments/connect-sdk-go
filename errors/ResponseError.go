package errors

import (
	"strconv"

	"github.com/Ingenico-ePayments/connect-sdk-go/communicator/communication"
)

// ResponseError is returned when a response was received from the GlobalCollect platform which indicates an error.
type ResponseError struct {
	response   communication.Response
	statusCode int
	body       string
	headers    []communication.Header
}

// Response gets the response that was returned by the GlobalCollect platform.
func (e *ResponseError) Response() communication.Response {
	return e.response
}

// StatusCode gets the HTTP status code that was returned by the GlobalCollect platform.
func (e *ResponseError) StatusCode() int {
	return e.statusCode
}

// Body gets the raw response body that was returned by the GlobalConnect platform.
func (e *ResponseError) Body() string {
	return e.body
}

// Headers gets the headers that were returned by the GlobalConnect platform.
func (e *ResponseError) Headers() []communication.Header {
	return e.headers
}

// GetHeader returns the header with the given name, or nil if there was no such header.
func (e *ResponseError) GetHeader(headerName string) *communication.Header {
	return e.response.GetHeader(headerName)
}

// String implements the Stringer interface
// Format: 'errorMessage; statusCode=; responseBody='
func (e *ResponseError) String() string {
	retString := "the GlobalCollect platform returned an error response"

	statusCode := e.response.StatusCode()
	if statusCode > 0 {
		retString += ";  statusCode=" + strconv.Itoa(e.response.StatusCode())
	}

	responseBody := e.response.Body()
	if len(responseBody) > 0 {
		retString += "; responseBody='" + e.response.Body() + "'"
	}

	return retString
}

// Error implements the Error interface
func (e *ResponseError) Error() string {
	return e.String()
}

// NewResponseError creates a new ResponseError with the specified response
func NewResponseError(r communication.Response) *ResponseError {
	return &ResponseError{response: r, statusCode: r.StatusCode(), body: r.Body(), headers: r.Headers()}
}
