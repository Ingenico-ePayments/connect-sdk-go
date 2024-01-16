// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package payment

// CreationReferences represents class PaymentCreationReferences
type CreationReferences struct {
	AdditionalReference *string `json:"additionalReference,omitempty"`
	ExternalReference   *string `json:"externalReference,omitempty"`
}

// NewCreationReferences constructs a new CreationReferences
func NewCreationReferences() *CreationReferences {
	return &CreationReferences{}
}
