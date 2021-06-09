// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// RedirectPaymentMethodSpecificInput represents class RedirectPaymentMethodSpecificInput
type RedirectPaymentMethodSpecificInput struct {
	ExpirationPeriod                  *int32                                   `json:"expirationPeriod,omitempty"`
	IsRecurring                       *bool                                    `json:"isRecurring,omitempty"`
	PaymentProduct4101SpecificInput   *RedirectPaymentProduct4101SpecificInput `json:"paymentProduct4101SpecificInput,omitempty"`
	PaymentProduct809SpecificInput    *RedirectPaymentProduct809SpecificInput  `json:"paymentProduct809SpecificInput,omitempty"`
	PaymentProduct816SpecificInput    *RedirectPaymentProduct816SpecificInput  `json:"paymentProduct816SpecificInput,omitempty"`
	PaymentProduct840SpecificInput    *RedirectPaymentProduct840SpecificInput  `json:"paymentProduct840SpecificInput,omitempty"`
	PaymentProduct861SpecificInput    *RedirectPaymentProduct861SpecificInput  `json:"paymentProduct861SpecificInput,omitempty"`
	PaymentProduct863SpecificInput    *RedirectPaymentProduct863SpecificInput  `json:"paymentProduct863SpecificInput,omitempty"`
	PaymentProduct869SpecificInput    *RedirectPaymentProduct869SpecificInput  `json:"paymentProduct869SpecificInput,omitempty"`
	PaymentProduct882SpecificInput    *RedirectPaymentProduct882SpecificInput  `json:"paymentProduct882SpecificInput,omitempty"`
	PaymentProductID                  *int32                                   `json:"paymentProductId,omitempty"`
	RecurringPaymentSequenceIndicator *string                                  `json:"recurringPaymentSequenceIndicator,omitempty"`
	RedirectionData                   *RedirectionData                         `json:"redirectionData,omitempty"`
	RequiresApproval                  *bool                                    `json:"requiresApproval,omitempty"`
	// Deprecated: Use redirectionData.returnUrl instead
	ReturnURL                         *string                                  `json:"returnUrl,omitempty"`
	Token                             *string                                  `json:"token,omitempty"`
	Tokenize                          *bool                                    `json:"tokenize,omitempty"`
}

// NewRedirectPaymentMethodSpecificInput constructs a new RedirectPaymentMethodSpecificInput
func NewRedirectPaymentMethodSpecificInput() *RedirectPaymentMethodSpecificInput {
	return &RedirectPaymentMethodSpecificInput{}
}
