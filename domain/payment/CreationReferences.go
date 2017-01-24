// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

// CreationReferences represents class PaymentCreationReferences
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_PaymentCreationReferences
type CreationReferences struct {
	AdditionalReference *string `json:"additionalReference,omitempty"`
	ExternalReference   *string `json:"externalReference,omitempty"`
}

// NewCreationReferences constructs a new CreationReferences
func NewCreationReferences() *CreationReferences {
	return &CreationReferences{}
}
