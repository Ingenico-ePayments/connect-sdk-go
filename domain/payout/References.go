// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package payout

// References represents class PayoutReferences
type References struct {
	InvoiceNumber     *string `json:"invoiceNumber,omitempty"`
	MerchantOrderID   *int64  `json:"merchantOrderId,omitempty"`
	MerchantReference *string `json:"merchantReference,omitempty"`
}

// NewReferences constructs a new References
func NewReferences() *References {
	return &References{}
}
