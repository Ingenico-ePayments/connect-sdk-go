package logging

import (
	"bytes"
	"fmt"
	"time"
)

const messageTemplate = `Incoming response (requestId='%s' in '%v'):
	status-code:  '%d'
	headers:      '%s'
	content-type: '%s'
	body:         '%s'`

// ResponseLogMessageBuilder represents a utility class to build request log messages.
type ResponseLogMessageBuilder struct {
	responseID string
	statusCode int
	duration   time.Duration

	body        string
	contentType string

	headers map[string][]string

	headersBuffer bytes.Buffer
}

// ResponseLogMessage represents a log message about a Response
type ResponseLogMessage struct {
	responseID string
	statusCode int
	duration   time.Duration

	body        string
	contentType string

	headers map[string][]string

	headersFormatted string
}

// ResponseID returns the response id
func (rl *ResponseLogMessage) ResponseID() string {
	return rl.responseID
}

// StatusCode returns the response status code
func (rl *ResponseLogMessage) StatusCode() int {
	return rl.statusCode
}

// Duration returns the transport duration
func (rl *ResponseLogMessage) Duration() time.Duration {
	return rl.duration
}

// Body returns the response body
func (rl *ResponseLogMessage) Body() string {
	return rl.body
}

// ContentType returns the response content-type
func (rl *ResponseLogMessage) ContentType() string {
	return rl.contentType
}

// Headers returns the response headers
func (rl *ResponseLogMessage) Headers() map[string][]string {
	return rl.headers
}

// String implements the Stringer interface
func (rl *ResponseLogMessage) String() string {
	return fmt.Sprintf(messageTemplate, rl.responseID, rl.duration, rl.statusCode, rl.headersFormatted, rl.contentType, rl.body)
}

// AddHeader adds a header to the log message using the name and value
func (rlm *ResponseLogMessageBuilder) AddHeader(name, value string) error {
	if rlm.headersBuffer.Len() > 0 {
		rlm.headersBuffer.WriteString(", ")
	}

	rlm.headersBuffer.WriteString(name)
	rlm.headersBuffer.WriteString("=\"")

	if len(value) > 0 {
		obfuscatedValue, oHErr := ObfuscateHeader(name, value)

		if oHErr != nil {
			return oHErr
		}

		rlm.headersBuffer.WriteString(obfuscatedValue)

		rlm.headers[name] = append(rlm.headers[name], obfuscatedValue)
	}
	rlm.headersBuffer.WriteString("\"")

	return nil
}

// SetBody sets the request body
func (rlm *ResponseLogMessageBuilder) SetBody(body, contentType string) error {
	obfuscatedBody, err := ObfuscateBody(body)
	if err != nil {
		return err
	}

	rlm.contentType = contentType
	rlm.body = obfuscatedBody

	return nil
}

// BuildMessage builds the ResponseLogMessage
func (rlm *ResponseLogMessageBuilder) BuildMessage() (*ResponseLogMessage, error) {
	return &ResponseLogMessage{rlm.responseID, rlm.statusCode, rlm.duration, rlm.body, rlm.contentType, rlm.headers, rlm.headersBuffer.String()}, nil
}

// NewResponseLogMessageBuilder creates a ResponseLogMessageBuilder with the given responseID, statusCode and duration
func NewResponseLogMessageBuilder(responseID string, statusCode int, duration time.Duration) (*ResponseLogMessageBuilder, error) {
	if len(responseID) < 1 {
		return nil, errRequestIDEmpty
	}

	return &ResponseLogMessageBuilder{responseID, statusCode, duration, "", "", map[string][]string{}, bytes.Buffer{}}, nil
}
