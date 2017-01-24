package logging

import (
	"bytes"
	"errors"
	"log"
	"strings"
	"testing"
)

func TestDefaultLog(t *testing.T) {
	testString := "Hello world!"
	expectedSuffix := testString + "\n"

	mockBuffer := &bytes.Buffer{}
	mockLogger := log.New(mockBuffer, "", log.Ldate | log.Ltime)

	var logger CommunicatorLogger
	logger, err := NewDefaultLogCommunicatorLogger(mockLogger)
	if err != nil {
		t.Fatalf("TestDefaultLog : %v", err)
	}

	logger.Log(testString)

	if strings.HasSuffix(mockBuffer.String(), expectedSuffix) == false {
		t.Fatalf("TestDefaultLog : expected suffix '%s' got '%s'", expectedSuffix, mockBuffer.String())
	}
}

func TestDefaultLogError(t *testing.T) {
	testString := "testString"
	testError := errors.New("testError")
	expectedSuffix := testString + " " + testError.Error() + "\n"

	mockBuffer := &bytes.Buffer{}
	mockLogger := log.New(mockBuffer, "", log.Ldate | log.Ltime)

	var logger CommunicatorLogger
	logger, err := NewDefaultLogCommunicatorLogger(mockLogger)
	if err != nil {
		t.Fatalf("TestDefaultLogError : %v", err)
	}

	logger.LogError(testString, testError)

	if strings.HasSuffix(mockBuffer.String(), expectedSuffix) == false {
		t.Fatalf("TestDefaultLog : expected suffix : '%s' got '%s'", expectedSuffix, mockBuffer.String())
	}
}
