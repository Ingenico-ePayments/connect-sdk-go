// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package hostedcheckout

// SpecificInput represents class HostedCheckoutSpecificInput
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_HostedCheckoutSpecificInput
type SpecificInput struct {
	IsRecurring           *bool                                `json:"isRecurring,omitempty"`
	Locale                *string                              `json:"locale,omitempty"`
	PaymentProductFilters *PaymentProductFiltersHostedCheckout `json:"paymentProductFilters,omitempty"`
	ReturnURL             *string                              `json:"returnUrl,omitempty"`
	ShowResultPage        *bool                                `json:"showResultPage,omitempty"`
	Tokens                *string                              `json:"tokens,omitempty"`
	Variant               *string                              `json:"variant,omitempty"`
}

// NewSpecificInput constructs a new SpecificInput
func NewSpecificInput() *SpecificInput {
	return &SpecificInput{}
}
