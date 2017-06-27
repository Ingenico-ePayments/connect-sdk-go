// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// LineItemInvoiceData represents class LineItemInvoiceData
type LineItemInvoiceData struct {
	Description        *string `json:"description,omitempty"`
	MerchantLinenumber *string `json:"merchantLinenumber,omitempty"`
	MerchantPagenumber *string `json:"merchantPagenumber,omitempty"`
	NrOfItems          *string `json:"nrOfItems,omitempty"`
	PricePerItem       *int64  `json:"pricePerItem,omitempty"`
}

// NewLineItemInvoiceData constructs a new LineItemInvoiceData
func NewLineItemInvoiceData() *LineItemInvoiceData {
	return &LineItemInvoiceData{}
}
