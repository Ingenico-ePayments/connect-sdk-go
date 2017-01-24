package errors

// NotFoundError indicates an exception that occurs when the requested resource is not found.
// In normal usage of the SDK, this exception should not occur, however it is possible.
// For example when path parameters are set with invalid values.
type NotFoundError struct {
	errorMessage  string
	internalError error
}

// InternalError returns the internal error encountered
func (nfe *NotFoundError) InternalError() error {
	return nfe.internalError
}

// Error implements the Error interface
func (nfe *NotFoundError) Error() string {
	return nfe.String()
}

// String implements the Stringer interface
// Format: 'errorMessage internalError'
func (nfe *NotFoundError) String() string {
	return nfe.errorMessage + " " + nfe.internalError.Error()
}

// NewNotFoundErrorVerbose creates a NotFoundError with the given errorMessage and internalError
func NewNotFoundErrorVerbose(errorMessage string, internalError error) (*NotFoundError, error) {
	return &NotFoundError{errorMessage, internalError}, nil
}
