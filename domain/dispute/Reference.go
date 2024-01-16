// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package dispute

// Reference represents class DisputeReference
type Reference struct {
	MerchantOrderID   *string `json:"merchantOrderId,omitempty"`
	MerchantReference *string `json:"merchantReference,omitempty"`
	PaymentReference  *string `json:"paymentReference,omitempty"`
	ProviderID        *string `json:"providerId,omitempty"`
	ProviderReference *string `json:"providerReference,omitempty"`
}

// NewReference constructs a new Reference
func NewReference() *Reference {
	return &Reference{}
}
