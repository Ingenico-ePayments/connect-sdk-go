// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payout

// References represents class PayoutReferences
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_PayoutReferences
type References struct {
	InvoiceNumber     *string `json:"invoiceNumber,omitempty"`
	MerchantOrderID   *int64  `json:"merchantOrderId,omitempty"`
	MerchantReference *string `json:"merchantReference,omitempty"`
}

// NewReferences constructs a new References
func NewReferences() *References {
	return &References{}
}
