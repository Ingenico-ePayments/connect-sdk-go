// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

// OrderReferences represents class OrderReferences
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_OrderReferences
type OrderReferences struct {
	Descriptor        *string           `json:"descriptor,omitempty"`
	InvoiceData       *OrderInvoiceData `json:"invoiceData,omitempty"`
	MerchantOrderID   *int64            `json:"merchantOrderId,omitempty"`
	MerchantReference *string           `json:"merchantReference,omitempty"`
}

// NewOrderReferences constructs a new OrderReferences
func NewOrderReferences() *OrderReferences {
	return &OrderReferences{}
}
