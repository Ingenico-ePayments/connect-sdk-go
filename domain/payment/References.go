// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// References represents class PaymentReferences
type References struct {
	MerchantOrderID      *int64  `json:"merchantOrderId,omitempty"`
	MerchantReference    *string `json:"merchantReference,omitempty"`
	PaymentReference     *string `json:"paymentReference,omitempty"`
	ProviderID           *string `json:"providerId,omitempty"`
	ProviderMerchantID   *string `json:"providerMerchantId,omitempty"`
	ProviderReference    *string `json:"providerReference,omitempty"`
	ReferenceOrigPayment *string `json:"referenceOrigPayment,omitempty"`
}

// NewReferences constructs a new References
func NewReferences() *References {
	return &References{}
}
