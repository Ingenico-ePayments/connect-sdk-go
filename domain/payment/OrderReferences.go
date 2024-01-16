// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package payment

// OrderReferences represents class OrderReferences
type OrderReferences struct {
	Descriptor         *string           `json:"descriptor,omitempty"`
	InvoiceData        *OrderInvoiceData `json:"invoiceData,omitempty"`
	MerchantOrderID    *int64            `json:"merchantOrderId,omitempty"`
	MerchantReference  *string           `json:"merchantReference,omitempty"`
	ProviderID         *string           `json:"providerId,omitempty"`
	ProviderMerchantID *string           `json:"providerMerchantId,omitempty"`
}

// NewOrderReferences constructs a new OrderReferences
func NewOrderReferences() *OrderReferences {
	return &OrderReferences{}
}
