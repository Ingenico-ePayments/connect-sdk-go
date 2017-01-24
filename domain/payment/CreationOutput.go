// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

// CreationOutput represents class PaymentCreationOutput
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_PaymentCreationOutput
type CreationOutput struct {
	AdditionalReference *string `json:"additionalReference,omitempty"`
	ExternalReference   *string `json:"externalReference,omitempty"`
	IsNewToken          *bool   `json:"isNewToken,omitempty"`
	Token               *string `json:"token,omitempty"`
}

// NewCreationOutput constructs a new CreationOutput
func NewCreationOutput() *CreationOutput {
	return &CreationOutput{}
}
