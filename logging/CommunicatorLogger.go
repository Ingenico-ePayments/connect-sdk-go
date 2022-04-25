package logging

// CommunicatorLogger is used to log messages from communicators. Thread-safe.
type CommunicatorLogger interface {
	// Log logs the specified message.
	Log(message string)

	// LogError logs an error with an accompanying message.
	LogError(message string, err error)

	// LogResponseLogMessage logs a response log message
	LogResponseLogMessage(response *ResponseLogMessage)

	// LogRequestLogMessage logs a request log message
	LogRequestLogMessage(request *RequestLogMessage)
}
