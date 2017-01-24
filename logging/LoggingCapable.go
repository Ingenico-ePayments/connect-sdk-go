package logging

// Capable is an interface used for logging messages from communicators.
type Capable interface {
	// Turns on logging using the given communicator logger.
	EnableLogging(communicatorLogger CommunicatorLogger)
	// Turns off logging.
	DisableLogging()
}
