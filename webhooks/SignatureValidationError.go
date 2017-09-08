package webhooks

// SignatureValidationError represents an error while validating webhooks signatures.
type SignatureValidationError struct {
	message string
}

// Error implements the Error interface
func (sve *SignatureValidationError) Error() string {
	return sve.message
}

// NewSignatureValidationError creates a SignatureValidationError with the given message
func NewSignatureValidationError(message string) (*SignatureValidationError, error) {
	return &SignatureValidationError{message}, nil
}
