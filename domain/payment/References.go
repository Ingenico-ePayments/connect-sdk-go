// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

// References represents class PaymentReferences
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_PaymentReferences
type References struct {
	MerchantOrderID      *int64  `json:"merchantOrderId,omitempty"`
	MerchantReference    *string `json:"merchantReference,omitempty"`
	PaymentReference     *string `json:"paymentReference,omitempty"`
	ProviderID           *string `json:"providerId,omitempty"`
	ProviderReference    *string `json:"providerReference,omitempty"`
	ReferenceOrigPayment *string `json:"referenceOrigPayment,omitempty"`
}

// NewReferences constructs a new References
func NewReferences() *References {
	return &References{}
}
