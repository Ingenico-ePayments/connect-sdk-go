package communication

import "io"

// ResponseHandlerFunc is a handler function for an incoming response
type ResponseHandlerFunc func(statusCode int, headers []Header, reader io.Reader) (interface{}, error)

// ResponseHandler is a handler for an incoming response
type ResponseHandler interface {
	Handle(statusCode int, headers []Header, reader io.Reader) (interface{}, error)
}

// Handle calls f(statusCode, headers, reader)
func (f ResponseHandlerFunc) Handle(statusCode int, headers []Header, reader io.Reader) (interface{}, error) {
	return f(statusCode, headers, reader)
}
