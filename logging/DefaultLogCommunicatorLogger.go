package logging

import "log"

// DefaultLogCommunicatorLogger adapts a log.Logger to the CommunicatorLogger interface
type DefaultLogCommunicatorLogger struct {
	goLog *log.Logger
}

// Log the specified message
func (dl *DefaultLogCommunicatorLogger) Log(message string) {
	dl.goLog.Println(message)
}

// LogError logs the error with a message
func (dl *DefaultLogCommunicatorLogger) LogError(message string, err error) {
	dl.goLog.Println(message, err)
}

// LogResponseLogMessage logs a ResponseLogMessage
func (dl *DefaultLogCommunicatorLogger) LogResponseLogMessage(response *ResponseLogMessage) {
	dl.goLog.Println(response)
}

// LogRequestLogMessage logs a RequestLogMessage
func (dl *DefaultLogCommunicatorLogger) LogRequestLogMessage(request *RequestLogMessage) {
	dl.goLog.Println(request)
}

// NewDefaultLogCommunicatorLogger creates a DefaultLogCommunicatorLogger with the given log.Logger
func NewDefaultLogCommunicatorLogger(goLog *log.Logger) (*DefaultLogCommunicatorLogger, error) {
	return &DefaultLogCommunicatorLogger{goLog}, nil
}
