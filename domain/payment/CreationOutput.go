// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// CreationOutput represents class PaymentCreationOutput
type CreationOutput struct {
	AdditionalReference   *string `json:"additionalReference,omitempty"`
	ExternalReference     *string `json:"externalReference,omitempty"`
	IsNewToken            *bool   `json:"isNewToken,omitempty"`
	Token                 *string `json:"token,omitempty"`
	TokenizationSucceeded *bool   `json:"tokenizationSucceeded,omitempty"`
}

// NewCreationOutput constructs a new CreationOutput
func NewCreationOutput() *CreationOutput {
	return &CreationOutput{}
}
