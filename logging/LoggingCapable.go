package logging

// Capable is an interface used for logging messages from communicators.
type Capable interface {
	// EnableLogging turns on logging using the given communicator logger.
	EnableLogging(communicatorLogger CommunicatorLogger)
	// DisableLogging turns off logging.
	DisableLogging()
}
