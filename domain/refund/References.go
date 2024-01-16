// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package refund

// References represents class RefundReferences
type References struct {
	MerchantReference *string `json:"merchantReference,omitempty"`
}

// NewReferences constructs a new References
func NewReferences() *References {
	return &References{}
}
