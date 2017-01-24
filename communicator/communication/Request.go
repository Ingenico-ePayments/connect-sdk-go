package communication

import (
	"io"
	"io/ioutil"
	"net/url"
	"strings"
)

// Request is a container containing everything needed to send a request
type Request struct {
	Method      string
	Destination url.URL
	Headers     []Header
	body        io.Reader
}

// NewRequest creates a new request
func NewRequest(httpMethod string, destination url.URL, headers []Header) *Request {
	return &Request{httpMethod, destination, headers, nil}
}

// Body returns the raw body of the request
func (r *Request) Body() io.Reader {
	return r.body
}

// BodyString returns the body of the request as a string
func (r *Request) BodyString() (string, error) {
	contents, err := ioutil.ReadAll(r.body)
	return string(contents), err
}

// SetBody sets the raw body
func (r *Request) SetBody(body io.Reader) {
	r.body = body
}

// SetBodyString sets the body as a string
func (r *Request) SetBodyString(body string) {
	r.body = strings.NewReader(body)
}
