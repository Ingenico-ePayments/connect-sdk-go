// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

// OrderInvoiceData represents class OrderInvoiceData
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_OrderInvoiceData
type OrderInvoiceData struct {
	AdditionalData *string   `json:"additionalData,omitempty"`
	InvoiceDate    *string   `json:"invoiceDate,omitempty"`
	InvoiceNumber  *string   `json:"invoiceNumber,omitempty"`
	TextQualifiers *[]string `json:"textQualifiers,omitempty"`
}

// NewOrderInvoiceData constructs a new OrderInvoiceData
func NewOrderInvoiceData() *OrderInvoiceData {
	return &OrderInvoiceData{}
}
