package logging

import (
	"log"
	"os"
)

var stdOutLogger = log.New(os.Stdout, "", log.Ldate | log.Ltime)
var stdOutCommunicatorLogger, _ = NewDefaultLogCommunicatorLogger(stdOutLogger)

// StdOutCommunicatorLogger returns a CommunicatorLogger that outputs to os.Stdout
func StdOutCommunicatorLogger() *DefaultLogCommunicatorLogger {
	return stdOutCommunicatorLogger
}
