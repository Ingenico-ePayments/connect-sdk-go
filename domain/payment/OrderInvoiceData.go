// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// OrderInvoiceData represents class OrderInvoiceData
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
