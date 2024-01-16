package webhooks

import (
	"github.com/Ingenico-ePayments/connect-sdk-go/webhooks/validation"
)

// APIVersionMismatchError represents an error because a webhooks event has an API version
// that this version of the SDK does not support.
type APIVersionMismatchError struct {
	error *validation.APIVersionMismatchError
}

// Error implements the error interface
func (ame *APIVersionMismatchError) Error() string {
	return ame.error.Error()
}

// EventAPIVersion represents the APIVersion found in the event
func (ame *APIVersionMismatchError) EventAPIVersion() string {
	return ame.error.EventAPIVersion()
}

// SDKAPIVersion represents the APIVersion found in the SDK
func (ame *APIVersionMismatchError) SDKAPIVersion() string {
	return ame.error.SDKAPIVersion()
}

// NewAPIVersionMismatchError creates a APIVersionMismatchError with the given eventAPIVersion and sdkAPIVersion
func NewAPIVersionMismatchError(eventAPIVersion, sdkAPIVersion string) (*APIVersionMismatchError, error) {
	err := validation.NewAPIVersionMismatchError(eventAPIVersion, sdkAPIVersion)

	return &APIVersionMismatchError{err}, nil
}
