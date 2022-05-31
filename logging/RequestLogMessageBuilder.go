package logging

import (
	"bytes"
	"errors"
	"fmt"
	"net/url"

	"github.com/Ingenico-ePayments/connect-sdk-go/logging/obfuscation"
)

var errRequestIDEmpty = errors.New("requestID can't be empty")

const messageTemplateWithoutBody = `Outgoing request (requestId='%s'):
	method:       '%s'
	uri:          '%s'
	headers:      '%s'`
const messageTemplateWithBody = messageTemplateWithoutBody + `
	content-type: '%s'
	body:         '%s'`

// RequestLogMessageBuilder represents utility class to build request log messages.
type RequestLogMessageBuilder struct {
	requestID string
	method    string
	url       url.URL

	body        string
	contentType string

	headers map[string][]string

	headersBuffer bytes.Buffer

	bodyObfuscator   obfuscation.BodyObfuscator
	headerObfuscator obfuscation.HeaderObfuscator
}

// RequestLogMessage represents a log message about a Request
type RequestLogMessage struct {
	requestID string
	method    string
	url       url.URL

	body        string
	contentType string

	headers map[string][]string

	headersFormatted string

	bodyObfuscator   obfuscation.BodyObfuscator
	headerObfuscator obfuscation.HeaderObfuscator
}

// RequestID returns the request ID
func (rl *RequestLogMessage) RequestID() string {
	return rl.requestID
}

// Method returns the request method
func (rl *RequestLogMessage) Method() string {
	return rl.method
}

// URL returns the request URL
func (rl *RequestLogMessage) URL() url.URL {
	return rl.url
}

// Body returns the request body
func (rl *RequestLogMessage) Body() string {
	return rl.body
}

// ContentType returns the content type
func (rl *RequestLogMessage) ContentType() string {
	return rl.contentType
}

// Headers returns the headers
func (rl *RequestLogMessage) Headers() map[string][]string {
	return rl.headers
}

// String implements the Stringer interface
func (rl *RequestLogMessage) String() string {
	if len(rl.body) < 1 {
		return fmt.Sprintf(messageTemplateWithoutBody, rl.requestID, rl.method, rl.url.Path, rl.headersFormatted)
	}

	return fmt.Sprintf(messageTemplateWithBody, rl.requestID, rl.method, rl.url.Path, rl.headersFormatted, rl.contentType, rl.body)
}

// AddHeader adds a header to the log message using the name and value
func (rlm *RequestLogMessageBuilder) AddHeader(name, value string) error {
	if rlm.headersBuffer.Len() > 0 {
		rlm.headersBuffer.WriteString(", ")
	}

	rlm.headersBuffer.WriteString(name)
	rlm.headersBuffer.WriteString("=\"")

	if len(value) > 0 {
		obfuscatedValue := rlm.headerObfuscator.ObfuscateHeader(name, value)

		rlm.headersBuffer.WriteString(obfuscatedValue)

		rlm.headers[name] = append(rlm.headers[name], obfuscatedValue)
	}
	rlm.headersBuffer.WriteString("\"")

	return nil
}

// SetBody sets the request body
func (rlm *RequestLogMessageBuilder) SetBody(body, contentType string) error {
	obfuscatedBody, err := rlm.bodyObfuscator.ObfuscateBody(body)
	if err != nil {
		return err
	}

	rlm.contentType = contentType
	rlm.body = obfuscatedBody

	return nil
}

// SetBinaryBody sets the binary request body
func (rlm *RequestLogMessageBuilder) SetBinaryBody(contentType string) error {
	if !isBinaryContent(contentType) {
		return errors.New("Not a binary content type: " + contentType)
	}
	rlm.contentType = contentType
	rlm.body = "<binary content>"

	return nil
}

// BuildMessage builds the RequestLogMessage
func (rlm *RequestLogMessageBuilder) BuildMessage() (*RequestLogMessage, error) {
	return &RequestLogMessage{rlm.requestID, rlm.method, rlm.url, rlm.body, rlm.contentType, rlm.headers, rlm.headersBuffer.String(), rlm.bodyObfuscator, rlm.headerObfuscator}, nil
}

// NewRequestLogMessageBuilder creates a RequestLogMessageBuilder with the given requestID, method and url
//
// Deprecated: use NewRequestLogMessageBuilderWithObfuscators instead
func NewRequestLogMessageBuilder(requestID, method string, url url.URL) (*RequestLogMessageBuilder, error) {
	return NewRequestLogMessageBuilderWithObfuscators(requestID, method, url, obfuscation.DefaultBodyObfuscator(), obfuscation.DefaultHeaderObfuscator())
}

// NewRequestLogMessageBuilderWithObfuscators creates a RequestLogMessageBuilder with the given requestID, method, url and obfuscators
func NewRequestLogMessageBuilderWithObfuscators(requestID, method string, url url.URL, bodyObfuscator obfuscation.BodyObfuscator, headerObfuscator obfuscation.HeaderObfuscator) (*RequestLogMessageBuilder, error) {
	if len(requestID) < 1 {
		return nil, errRequestIDEmpty
	}

	return &RequestLogMessageBuilder{requestID, method, url, "", "", map[string][]string{}, bytes.Buffer{}, bodyObfuscator, headerObfuscator}, nil
}
