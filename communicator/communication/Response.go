package communication

import (
	"bytes"
	"reflect"
	"strconv"
	"strings"
)

// Response represents a response from the Ingenico ePayments platform.
type Response struct {
	statusCode int //see net/http/status.go accessible through http.{status name}
	body       string
	headers    []Header
}

// NewResponse creates a Response with the given status code, body and headers
func NewResponse(statusCode int, body string, headers []Header) (*Response, error) {
	return &Response{statusCode: statusCode, body: body, headers: headers}, nil
}

// String() is the implementation of the Stringer interface
// Format: '[statusCode=, body=, headers=]'
func (r *Response) String() string {
	var buffer bytes.Buffer

	buffer.WriteString(reflect.TypeOf(r).Name())
	buffer.WriteString("[statusCode=")
	buffer.WriteString(strconv.Itoa(r.statusCode))

	if len(r.body) > 0 {
		buffer.WriteString(",body='")
		buffer.WriteString(r.body)
		buffer.WriteString("'")
	}

	if len(r.headers) > 0 {
		buffer.WriteString(",headers=")

		for _, v := range r.headers {
			buffer.WriteString(v.String())
			buffer.WriteString(" ")
		}
	}

	buffer.WriteString("]")

	return buffer.String()
}

// GetHeader gets the specified header.
func (r *Response) GetHeader(headerName string) *Header {
	for _, header := range r.headers {
		if strings.EqualFold(header.Name(), headerName) {
			return &header
		}
	}
	return nil
}

// StatusCode gets The HTTP status code that was returned by the Ingenico ePayments platform
func (r *Response) StatusCode() int {
	return r.statusCode
}

// Body gets the raw response body that was returned by the Ingenico ePayments platform
func (r *Response) Body() string {
	return r.body
}

// Headers gets the headers that were returned by the Ingenico ePayments platform
func (r *Response) Headers() []Header {
	// Return a clone instead of the original slice - immutability insurance
	return append([]Header{}, r.headers...)
}
