package logging

// CommunicatorLogger is used to log messages from communicators. Thread-safe.
type CommunicatorLogger interface {
	// Logs the specified message.
	Log(message string)

	// Logs an error with an accompanying message.
	LogError(message string, err error)

	// Logs a response log message
	LogResponseLogMessage(response *ResponseLogMessage)

	// Logs a request log message
	LogRequestLogMessage(request *RequestLogMessage)
}
