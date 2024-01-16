// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package product

// MobilePaymentProductSession302SpecificInput represents class MobilePaymentProductSession302SpecificInput
type MobilePaymentProductSession302SpecificInput struct {
	DisplayName   *string `json:"displayName,omitempty"`
	DomainName    *string `json:"domainName,omitempty"`
	ValidationURL *string `json:"validationUrl,omitempty"`
}

// NewMobilePaymentProductSession302SpecificInput constructs a new MobilePaymentProductSession302SpecificInput
func NewMobilePaymentProductSession302SpecificInput() *MobilePaymentProductSession302SpecificInput {
	return &MobilePaymentProductSession302SpecificInput{}
}
