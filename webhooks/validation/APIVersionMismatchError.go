package validation

import "fmt"

var (
	mismatchErrorFormat = `"event API version "%v" is not compatible with SDK API version "%v"`
)

// APIVersionMismatchError represents an error because a webhooks event has an API version
// that this version of the SDK does not support.
type APIVersionMismatchError struct {
	eventAPIVersion string
	sdkAPIVersion   string

	message string
}

// Error implements the error interface
func (ame *APIVersionMismatchError) Error() string {
	return ame.message
}

// EventAPIVersion represents the APIVersion found in the event
func (ame *APIVersionMismatchError) EventAPIVersion() string {
	return ame.eventAPIVersion
}

// SDKAPIVersion represents the APIVersion found in the SDK
func (ame *APIVersionMismatchError) SDKAPIVersion() string {
	return ame.sdkAPIVersion
}

// NewAPIVersionMismatchError creates a APIVersionMismatchError with the given eventAPIVersion and sdkAPIVersion
func NewAPIVersionMismatchError(eventAPIVersion, sdkAPIVersion string) *APIVersionMismatchError {
	message := fmt.Sprintf(mismatchErrorFormat, eventAPIVersion, sdkAPIVersion)

	return &APIVersionMismatchError{eventAPIVersion, sdkAPIVersion, message}
}
